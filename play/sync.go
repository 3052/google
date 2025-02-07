package play

import (
   "41.neocities.org/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

func (d *Device) Sync(check Checkin) error {
   message := protobuf.Message{}
   message.Add(1, func(m protobuf.Message) {
      m.Add(10, func(m protobuf.Message) {
         for _, feature := range d.Feature {
            m.Add(1, func(m protobuf.Message) {
               m.AddBytes(1, []byte(feature))
            })
         }
         for _, library := range d.Library {
            m.AddBytes(2, []byte(library))
         }
         for _, texture := range d.Texture {
            m.AddBytes(4, []byte(texture))
         }
      })
   })
   message.Add(1, func(m protobuf.Message) {
      m.Add(15, func(m protobuf.Message) {
         m.AddBytes(4, []byte(d.Abi))
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
   x_dfe_device_id(req, check)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   _, err = io.Copy(io.Discard, resp.Body)
   if err != nil {
      return err
   }
   return nil
}
