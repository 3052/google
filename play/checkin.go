package play

import (
   "41.neocities.org/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

type Checkin [1]protobuf.Message

func (c *Checkin) Unmarshal(data Byte[Checkin]) error {
   return c[0].Unmarshal(data)
}

func (d *Device) Checkin() (Byte[Checkin], error) {
   var value protobuf.Message
   value.Add(4, func(m *protobuf.Message) {
      m.Add(1, func(m *protobuf.Message) {
         m.AddVarint(10, android_api)
      })
   })
   value.AddVarint(14, 3)
   value.Add(18, func(m *protobuf.Message) {
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
      m.AddBytes(11, []byte(d.Abi))
      for _, texture := range d.Texture {
         m.AddBytes(15, []byte(texture))
      }
      for _, feature := range d.Feature {
         // this line needs to be in the loop:
         m.Add(26, func(m *protobuf.Message) {
            m.AddBytes(1, []byte(feature))
         })
      }
   })
   resp, err := http.Post(
      "https://android.googleapis.com/checkin",
      "application/x-protobuffer",
      bytes.NewReader(value.Marshal()),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   return io.ReadAll(resp.Body)
}

func (c Checkin) field_7() uint64 {
   for value := range c[0].GetI64(7) {
      return uint64(value)
   }
   return 0
}
