package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

func (g GoogleCheckin) device_id() (uint64, error) {
   if v, ok := <-g.Message.GetFixed64(7); ok {
      return uint64(v), nil
   }
   return 0, errors.New("x-dfe-device-id")
}

func (g GoogleDevice) GoogleCheckin() ([]byte, error) {
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
      for _, v := range g.Library {
         m.AddBytes(9, []byte(v))
      }
      m.AddBytes(11, []byte(g.Abi))
      for _, v := range g.Texture {
         m.AddBytes(15, []byte(v))
      }
      // you cannot swap the next two lines:
      for _, v := range g.Feature {
         m.Add(26, func(m *protobuf.Message) {
            m.AddBytes(1, []byte(v))
         })
      }
   })
   resp, err := http.Post(
      "https://android.googleapis.com/checkin",
      "application/x-protobuffer",
      bytes.NewReader(m.Encode()),
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

type GoogleCheckin struct {
   Message protobuf.Message
}

func (g *GoogleCheckin) Unmarshal(data []byte) error {
   return g.Message.Consume(data)
}
