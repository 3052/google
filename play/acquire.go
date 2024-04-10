package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

// play.google.com/store/apps/details?id=com.google.android.youtube
func (g GoogleAuth) Acquire(g GoogleCheckin, id string) error {
   var m protobuf.Message
   m.Add(1, func(m *protobuf.Message) {
      m.Add(1, func(m *protobuf.Message) {
         m.AddBytes(1, []byte(id))
         m.AddVarint(2, 1)
         m.AddVarint(3, 3)
      })
      m.AddVarint(2, 1)
   })
   m.AddVarint(13, 1)
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/acquire",
      bytes.NewReader(m.Encode()),
   )
   if err != nil {
      return err
   }
   a.authorization(req)
   if err := g.x_dfe_device_id(req); err != nil {
      return err
   }
   // with a new device, this needs to be included in the first request to
   // /fdfe/acquire, or you get:
   // Please open my apps to establish a connection with the server.
   // on following requests you can omit it
   if err := g.x_ps_rh(req); err != nil {
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
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   if err := m.Consume(data); err != nil {
      return err
   }
   m = <-m.Get(1)
   m = <-m.Get(94)
   m = <-m.Get(1)
   m = <-m.Get(2)
   if m, ok := <-m.Get(147291249); ok {
      return acquire_error{m}
   }
   return nil
}

type acquire_error struct {
   m protobuf.Message
}

func (a acquire_error) Error() string {
   var b []byte
   for v := range a.m.Get(1) {
      v = <-v.Get(10)
      v = <-v.Get(1)
      if v, ok := <-v.GetBytes(1); ok {
         if b != nil {
            b = append(b, '\n')
         }
         b = append(b, v...)
      }
   }
   return string(b)
}
