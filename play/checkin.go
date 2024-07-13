package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

type GoogleCheckin struct {
   Data []byte
   Message protobuf.Message
}

func (g *GoogleCheckin) Checkin(device GoogleDevice) error {
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
      for _, v := range device.Library {
         m.AddBytes(9, []byte(v))
      }
      m.AddBytes(11, []byte(device.Abi))
      for _, v := range device.Texture {
         m.AddBytes(15, []byte(v))
      }
      // you cannot swap the next two lines:
      for _, v := range device.Feature {
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
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   g.Data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (g *GoogleCheckin) Unmarshal() error {
   return g.Message.Consume(g.Data)
}

// x-dfe-device-id
func (g GoogleCheckin) device_id() (uint64, error) {
   if v, ok := <-g.Message.GetFixed64(7); ok {
      return uint64(v), nil
   }
   return 0, errors.New("x-dfe-device-id")
}
