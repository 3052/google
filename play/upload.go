package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "net/http"
)

func (h Header) Upload(d Device) error {
   var m protobuf.Message
   m.Add(1, func(m *protobuf.Message) {
      m.Add_Varint(1, 3)
      m.Add_Varint(2, 2)
      m.Add_Varint(3, 1)
      m.Add_Varint(4, 2)
      m.Add_Varint(5, 1)
      m.Add_Varint(6, 0)
      m.Add_Varint(7, 420)
      m.Add_Varint(8, 196609)
      for _, library := range d.Library {
         m.Add_String(9, library)
      }
      m.Add_String(11, d.Platform)
      for _, texture := range d.Texture {
         m.Add_String(15, texture)
      }
      // you cannot swap the next two lines:
      for _, feature := range d.Feature {
         m.Add(26, func(m *protobuf.Message) {
            m.Add_String(1, feature)
         })
      }
   })
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/uploadDeviceConfig",
      bytes.NewReader(m.Append(nil)),
   )
   if err != nil {
      return err
   }
   req.Header.Set("X-DFE-Client-Id", "am-unknown")
   req.Header.Set(h.Agent())
   req.Header.Set(h.Device_ID())
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
