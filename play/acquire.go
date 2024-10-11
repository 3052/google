package play

import (
   "41.neocities.org/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

func (g GoogleAuth) Acquire(checkin *GoogleCheckin, id string) error {
   message := protobuf.Message{
      1: {protobuf.Message{
         1: {protobuf.Message{
            1: {protobuf.Bytes(id)},
            2: {protobuf.Varint(1)},
            3: {protobuf.Varint(3)},
         }},
         2: {protobuf.Varint(1)},
      }},
      13: {protobuf.Varint(1)},
   }
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/acquire",
      bytes.NewReader(message.Marshal()),
   )
   if err != nil {
      return err
   }
   authorization(req, g)
   x_dfe_device_id(req, checkin)
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
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   message = protobuf.Message{}
   err = message.Unmarshal(data)
   if err != nil {
      return err
   }
   message, _ = message.Get(1)()
   message, _ = message.Get(94)()
   message, _ = message.Get(1)()
   message, _ = message.Get(2)()
   if message, ok := message.Get(147291249)(); ok {
      return &acquire_error{message}
   }
   return nil
}

type acquire_error struct {
   message protobuf.Message
}

func (a acquire_error) Error() string {
   var text []byte
   next := a.message.Get(1)
   for {
      v, ok := next()
      if !ok {
         break
      }
      v, _ = v.Get(10)()
      v, _ = v.Get(1)()
      if v, ok := v.GetBytes(1)(); ok {
         if text != nil {
            text = append(text, '\n')
         }
         text = append(text, v...)
      }
   }
   return string(text)
}
