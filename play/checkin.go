package play

import (
   "41.neocities.org/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

func (Checkin) Marshal(device0 *Device) ([]byte, error) {
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
      for _, library := range device0.Library {
         m.AddBytes(9, []byte(library))
      }
      m.AddBytes(11, []byte(device0.Abi))
      for _, texture := range device0.Texture {
         m.AddBytes(15, []byte(texture))
      }
      for _, feature := range device0.Feature {
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

func (c *Checkin) Unmarshal(data []byte) error {
   (*c)[0] = protobuf.Message{}
   return (*c)[0].Unmarshal(data)
}

func (c Checkin) field_7() uint64 {
   value, _ := c[0].GetI64(7)()
   return uint64(value)
}

type Checkin [1]protobuf.Message
