package play

import (
   "41.neocities.org/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

func (GoogleCheckin) Marshal(device *GoogleDevice) ([]byte, error) {
   message := protobuf.Message{}
   message.Add(4, func(m protobuf.Message) {
      m.Add(1, func(m protobuf.Message) {
         m.AddVarint(10, android_api)
      })
   })
   message.AddVarint(14, 3)
   message.Add(18, func(m protobuf.Message) {
      m.AddVarint(1, 3)
      m.AddVarint(2, 2)
      m.AddVarint(3, 2)
      m.AddVarint(4, 2)
      m.AddVarint(5, 1)
      m.AddVarint(6, 1)
      m.AddVarint(7, 420)
      m.AddVarint(8, gl_es_version)
      for _, library := range device.Library {
         m.AddBytes(9, []byte(library))
      }
      m.AddBytes(11, []byte(device.Abi))
      for _, texture := range device.Texture {
         m.AddBytes(15, []byte(texture))
      }
      for _, feature := range device.Feature {
         // this line needs to be in the loop:
         m.Add(26, func(m protobuf.Message) {
            m.AddBytes(1, []byte(feature))
         })
      }
   })
   resp, err := http.Post(
      "https://android.googleapis.com/checkin",
      "application/x-protobuffer",
      bytes.NewReader(message.Marshal()),
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

func (g *GoogleCheckin) Unmarshal(data []byte) error {
   g.Message = protobuf.Message{}
   return g.Message.Unmarshal(data)
}

type GoogleCheckin struct {
   Message protobuf.Message
}

func (g GoogleCheckin) field_7() uint64 {
   value, _ := g.Message.GetFixed64(7)()
   return uint64(value)
}
