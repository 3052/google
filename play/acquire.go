package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

func (g *GoogleAuth) Acquire(checkin *GoogleCheckin, doc string) error {
   var message protobuf.Message
   message.Add(1, func(m *protobuf.Message) {
      m.Add(1, func(m *protobuf.Message) {
         m.AddBytes(1, []byte(doc))
         m.AddVarint(2, 1)
         m.AddVarint(3, 3)
      })
      m.AddVarint(2, 1)
   })
   message.AddVarint(13, 1)
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/acquire",
      bytes.NewReader(message.Encode()),
   )
   if err != nil {
      return err
   }
   authorization(req, g)
   err = x_dfe_device_id(req, checkin)
   if err != nil {
      return err
   }
   // with a new device, this needs to be included in the first request to
   // /fdfe/acquire, or you get:
   // Please open my apps to establish a connection with the server.
   // on following requests you can omit it
   err = x_ps_rh(req, checkin)
   if err != nil {
      return err
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b bytes.Buffer
      resp.Write(&b)
      return errors.New(b.String())
   }
   body, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   err = message.Consume(body)
   if err != nil {
      return err
   }
   message = <-message.Get(1)
   message = <-message.Get(94)
   message = <-message.Get(1)
   message = <-message.Get(2)
   if message, ok := <-message.Get(147291249); ok {
      return &acquire_error{message}
   }
   return nil
}

type acquire_error struct {
   message protobuf.Message
}

func (a *acquire_error) Error() string {
   var b []byte
   for v := range a.message.Get(1) {
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
