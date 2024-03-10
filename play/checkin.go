package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

func (c *Checkin) Unmarshal() error {
   return c.m.Consume(c.Data)
}

type Checkin struct {
   Data []byte
   m protobuf.Message
}

func (c Checkin) DeviceId() (uint64, error) {
   if v, ok := <-c.m.GetFixed64(7); ok {
      return uint64(v), nil
   }
   return 0, errors.New("Checkin.DeviceId")
}

func (c *Checkin) Do(d Device) error {
   var m protobuf.Message
   m.Add(4, func(m *protobuf.Message) {
      m.Add(1, func(m *protobuf.Message) {
         m.AddVarint(10, android_api)
      })
   })
   m.AddVarint(14, 3)
   m.Add(18, func(m *protobuf.Message) {
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
         m.Add(26, func(m *protobuf.Message) {
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
   c.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}
