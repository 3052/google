package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

func (g *GoogleAuth) Acquire(checkin *GoogleCheckin, doc string) error {
   message := protobuf.Message{}
   message.Add(1, func(m protobuf.Message) {
      m.Add(1, func(m protobuf.Message) {
         m.AddBytes(1, []byte(doc))
         m.AddVarint(2, 1)
         m.AddVarint(3, 3)
      })
      m.AddVarint(2, 1)
   })
   message.AddVarint(13, 1)
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/acquire",
      bytes.NewReader(message.Marshal()),
   )
   if err != nil {
      return err
   }
   authorization(req, g)
   if err = x_dfe_device_id(req, checkin); err != nil {
      return err
   }
   // with a new device, this needs to be included in the first request to
   // /fdfe/acquire, or you get:
   // Please open my apps to establish a connection with the server.
   // on following requests you can omit it
   if err = x_ps_rh(req, checkin); err != nil {
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
   message = protobuf.Message{}
   if err = message.Unmarshal(body); err != nil {
      return err
   }
   unknown := <-message.GetUnknown(1)
   unknown = <-unknown.Get(94)
   unknown = <-unknown.Get(1)
   unknown = <-unknown.Get(2)
   if unknown, ok := <-unknown.Get(147291249); ok {
      return &acquire_error{unknown}
   }
   return nil
}

type acquire_error struct {
   unknown protobuf.UnknownMessage
}

func (a *acquire_error) Error() string {
   var text []byte
   for u := range a.unknown.Get(1) {
      u = <-u.Get(10)
      u = <-u.Get(1)
      if v, ok := <-u.GetBytes(1); ok {
         if text != nil {
            text = append(text, '\n')
         }
         text = append(text, v...)
      }
   }
   return string(text)
}
