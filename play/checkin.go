package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

type Checkin struct {
   Raw []byte
   m protobuf.Message
}

func (c *Checkin) Checkin(d Device) error {
   var m protobuf.Message
   m.AddFunc(4, func(m *protobuf.Message) {
      m.AddFunc(1, func(m *protobuf.Message) {
         m.AddVarint(10, android_API)
      })
   })
   m.AddVarint(14, 3)
   m.AddFunc(18, func(m *protobuf.Message) {
      m.AddVarint(1, 3)
      m.AddVarint(2, 2)
      m.AddVarint(3, 2)
      m.AddVarint(4, 2)
      m.AddVarint(5, 1)
      m.AddVarint(6, 1)
      m.AddVarint(7, 420)
      m.AddVarint(8, gl_es_version)
      for _, library := range d.Library {
         m.AddBytes(9, []byte(library))
      }
      m.AddBytes(11, []byte(d.Platform))
      for _, texture := range d.Texture {
         m.AddBytes(15, []byte(texture))
      }
      // you cannot swap the next two lines:
      for _, feature := range d.Feature {
         m.AddFunc(26, func(m *protobuf.Message) {
            m.AddBytes(1, []byte(feature))
         })
      }
   })
   res, err := http.Post(
      "https://android.googleapis.com/checkin",
      "application/x-protobuffer",
      bytes.NewReader(m.Encode()),
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

func (c *Checkin) Unmarshal() error {
   return c.m.Consume(c.Raw)
}

func (c Checkin) Device_ID() (uint64, error) {
   if v, ok := c.m.GetFixed64(7); ok {
      return uint64(v), nil
   }
   return 0, errors.New("device ID")
}
