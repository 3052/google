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

type Checkin struct {
   m protobuf.Message
}

func (c Checkin) device_ID() (uint64, bool) {
   return c.m.Fixed64(7)
}

func (c Checkin) X_DFE_Device_ID() (string, string) {
   id, _ := c.device_ID()
   return "X-DFE-Device-ID", strconv.FormatUint(id, 16)
}

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

func (c Checkin) X_PS_RH() (string, string) {
   id, _ := c.device_ID()
   token, _ := func() (string, error) {
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
   ps_rh, _ := func() (string, error) {
      var m protobuf.Message
      m.Add(1, func(m *protobuf.Message) {
         m.Add_String(1, token)
      })
      return compress(m)
   }()
   return "X-PS-RH", ps_rh
}

//////////////////////////////////////////////////

// device is Pixel 2
func New_Checkin(d Device) ([]byte, error) {
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
