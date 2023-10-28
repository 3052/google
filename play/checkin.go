package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "compress/gzip"
   "encoding/base64"
   "errors"
   "io"
   "net/http"
   "strconv"
   "time"
)

func compress(m protobuf.Message) (string, error) {
   var b bytes.Buffer
   w := gzip.NewWriter(&b)
   _, err := w.Write(m.Append(nil))
   if err != nil {
      return "", err
   }
   if err := w.Close(); err != nil {
      return "", err
   }
   return base64.URLEncoding.EncodeToString(b.Bytes()), nil
}

func x_dfe_device_id(r *http.Request, c *Checkin) error {
   id, err := c.device_ID()
   if err != nil {
      return err
   }
   r.Header.Set("X-DFE-Device-ID", strconv.FormatUint(id, 16))
   return nil
}

func x_ps_rh(r *http.Request, c *Checkin) error {
   id, err := c.device_ID()
   if err != nil {
      return err
   }
   token, err := func() (string, error) {
      var m protobuf.Message
      m.Add(3, func(m *protobuf.Message) {
         m.Add_String(1, strconv.FormatUint(id, 10))
         m.Add(2, func(m *protobuf.Message) {
            v := time.Now().UnixMicro()
            m.Add_String(1, strconv.FormatInt(v, 10))
         })
      })
      return compress(m)
   }()
   if err != nil {
      return err
   }
   ps_rh, err := func() (string, error) {
      var m protobuf.Message
      m.Add(1, func(m *protobuf.Message) {
         m.Add_String(1, token)
      })
      return compress(m)
   }()
   if err != nil {
      return err
   }
   r.Header.Set("X-PS-RH", ps_rh)
   return nil
}

type Checkin struct {
   m protobuf.Message
}

func (c Checkin) device_ID() (uint64, error) {
   v, ok := c.m.Fixed64(7)
   if !ok {
      return 0, errors.New("device ID")
   }
   return v, nil
}

// device is Pixel 2
func (d Device) Checkin() (Raw_Checkin, error) {
   var m protobuf.Message
   m.Add(4, func(m *protobuf.Message) {
      m.Add(1, func(m *protobuf.Message) {
         m.Add_Varint(10, google_services_framework)
      })
   })
   m.Add_Varint(14, 3)
   m.Add(18, func(m *protobuf.Message) {
      m.Add_Varint(1, 3)
      m.Add_Varint(2, 2)
      m.Add_Varint(3, 1)
      m.Add_Varint(4, 2)
      m.Add_Varint(5, 1)
      m.Add_Varint(6, 0)
      m.Add_Varint(7, 420)
      // developer.android.com/guide/topics/manifest/uses-feature-element#glEsVersion
      m.Add_Varint(8, 196609)
      for _, library := range d.Library {
         m.Add_String(9, library)
      }
      m.Add_String(11, d.Platform)
      for _, texture := range d.Texture {
         m.Add_String(15, texture)
      }
      // you cannot swap the next two lines:
      for _, feature := range d.Feature {
         m.Add(26, func(m *protobuf.Message) {
            m.Add_String(1, feature)
         })
      }
   })
   res, err := http.Post(
      "https://android.clients.google.com/checkin",
      "application/x-protobuffer",
      bytes.NewReader(m.Append(nil)),
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   return io.ReadAll(res.Body)
}

type Raw_Checkin []byte

func (r Raw_Checkin) Checkin() (*Checkin, error) {
   m, err := protobuf.Consume(r)
   if err != nil {
      return nil, err
   }
   return &Checkin{m}, nil
}
