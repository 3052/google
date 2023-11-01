package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "net/http"
)

func (c Checkin) Sync(d Device) error {
   var m protobuf.Message
   m.Add(1, func(m *protobuf.Message) {
      m.Add(10, func(m *protobuf.Message) {
         for _, feature := range d.Feature {
            m.Add(1, func(m *protobuf.Message) {
               m.Add_String(1, feature)
            })
         }
         for _, library := range d.Library {
            m.Add_String(2, library)
         }
         for _, texture := range d.Texture {
            m.Add_String(4, texture)
         }
      })
   })
   m.Add(1, func(m *protobuf.Message) {
      m.Add(15, func(m *protobuf.Message) {
         m.Add_String(4, d.Platform)
      })
   })
   m.Add(1, func(m *protobuf.Message) {
      m.Add(18, func(m *protobuf.Message) {
         m.Add_String(1, "am-unknown") // X-DFE-Client-Id
      })
   })
   m.Add(1, func(m *protobuf.Message) {
      m.Add(19, func(m *protobuf.Message) {
         m.Add_Varint(2, google_play_store)
      })
   })
   m.Add(1, func(m *protobuf.Message) {
      m.Add(21, func(m *protobuf.Message) {
         m.Add_Varint(6, gl_es_version)
      })
   })
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/sync",
      bytes.NewReader(m.Append(nil)),
   )
   if err != nil {
      return err
   }
   if err := x_dfe_device_id(req, c); err != nil {
      return err
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return nil
}
