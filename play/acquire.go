package play

import (
   "41.neocities.org/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

func (a acquire_error) Error() string {
   var out []byte
   messages := protobuf.Message(a).Get(1)
   for {
      message, ok := messages()
      if !ok {
         break
      }
      message, _ = message.Get(10)()
      message, _ = message.Get(1)()
      in, _ := message.GetBytes(1)()
      if out != nil {
         out = append(out, '\n')
      }
      out = append(out, in...)
   }
   return string(out)
}

type acquire_error protobuf.Message

func (g GoogleAuth) Acquire(checkin GoogleCheckin, id string) error {
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
   message, ok := message.Get(147291249)()
   if ok {
      return acquire_error(message)
   }
   return nil
}
