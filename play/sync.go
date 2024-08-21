package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "net/http"
)

func (g *GoogleDevice) Sync(checkin *GoogleCheckin) error {
   message := protobuf.Message{}
   message.Add(1, func(m protobuf.Message) {
      m.Add(10, func(m protobuf.Message) {
         for _, value := range g.Feature {
            m.Add(1, func(m protobuf.Message) {
               m.AddBytes(1, []byte(value))
            })
         }
         for _, value := range g.Library {
            m.AddBytes(2, []byte(value))
         }
         for _, value := range g.Texture {
            m.AddBytes(4, []byte(value))
         }
      })
   })
   message.Add(1, func(m protobuf.Message) {
      m.Add(15, func(m protobuf.Message) {
         m.AddBytes(4, []byte(g.Abi))
      })
   })
   message.Add(1, func(m protobuf.Message) {
      m.Add(18, func(m protobuf.Message) {
         m.AddBytes(1, []byte("am-unknown")) // X-DFE-Client-Id
      })
   })
   message.Add(1, func(m protobuf.Message) {
      m.Add(19, func(m protobuf.Message) {
         m.AddVarint(2, google_play_store)
      })
   })
   message.Add(1, func(m protobuf.Message) {
      m.Add(21, func(m protobuf.Message) {
         m.AddVarint(6, gl_es_version)
      })
   })
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/sync",
      bytes.NewReader(message.Marshal()),
   )
   if err != nil {
      return err
   }
   if err = x_dfe_device_id(req, checkin); err != nil {
      return err
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   return nil
}
