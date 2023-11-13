package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

func (c *Checkin) Checkin(d Device) error {
   var m protobuf.Message
   m.Add(4, func(m *protobuf.Message) {
      m.Add(1, func(m *protobuf.Message) {
         m.Add_Varint(10, android_API)
      })
   })
   m.Add_Varint(14, 3)
   m.Add(18, func(m *protobuf.Message) {
      m.Add_Varint(1, 3)
      m.Add_Varint(2, 2)
      m.Add_Varint(3, 2)
      m.Add_Varint(4, 2)
      m.Add_Varint(5, 1)
      m.Add_Varint(6, 1)
      m.Add_Varint(7, 420)
      m.Add_Varint(8, gl_es_version)
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
      "https://android.googleapis.com/checkin",
      "application/x-protobuffer",
      bytes.NewReader(m.Append(nil)),
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   c.Raw, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}

type Checkin struct {
   Raw []byte
   m protobuf.Message
}

func (c Checkin) Device_ID() (uint64, error) {
   if v, ok := c.m.Fixed64(7); ok {
      return v, nil
   }
   return 0, errors.New("device ID")
}

func (c *Checkin) Unmarshal() error {
   var err error
   c.m, err = protobuf.Consume(c.Raw)
   if err != nil {
      return err
   }
   return nil
}
