package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

type GoogleCheckin struct {
   Message protobuf.Message
   Raw     []byte
}

func (g *GoogleCheckin) Unmarshal() error {
   g.Message = protobuf.Message{}
   return g.Message.Unmarshal(g.Raw)
}

func (g *GoogleDevice) Checkin() (*GoogleCheckin, error) {
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
      for _, library := range g.Library {
         m.AddBytes(9, []byte(library))
      }
      m.AddBytes(11, []byte(g.Abi))
      for _, texture := range g.Texture {
         m.AddBytes(15, []byte(texture))
      }
      // you cannot swap the next two lines:
      for _, feature := range g.Feature {
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
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   return &GoogleCheckin{Raw: data}, nil
}

func (g *GoogleCheckin) device_id() (uint64, error) {
   if v, ok := g.Message.GetFixed64(7)(); ok {
      return uint64(v), nil
   }
   return 0, errors.New("x-dfe-device-id")
}

