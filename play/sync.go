package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "net/http"
)

type Sync_Request struct {
   Device Device
   Checkin *Checkin
}

func (s Sync_Request) Do() error {
   var m protobuf.Message
   m.Add(1, func(m *protobuf.Message) {
      m.Add(10, func(m *protobuf.Message) {
         for _, feature := range s.Device.Feature {
            m.Add(1, func(m *protobuf.Message) {
               m.Add_String(1, feature)
            })
         }
         for _, library := range s.Device.Library {
            m.Add_String(2, library)
         }
         for _, texture := range s.Device.Texture {
            m.Add_String(4, texture)
         }
      })
   })
   m.Add(1, func(m *protobuf.Message) {
      m.Add(15, func(m *protobuf.Message) {
         m.Add_String(4, s.Device.Platform)
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
   // developer.android.com/guide/topics/manifest/uses-feature-element#glEsVersion
   m.Add(1, func(m *protobuf.Message) {
      m.Add(21, func(m *protobuf.Message) {
         m.Add_Varint(6, 196609)
      })
   })
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/sync",
      bytes.NewReader(m.Append(nil)),
   )
   if err != nil {
      return err
   }
   X_DFE_Device_ID(req, s.Checkin)
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
