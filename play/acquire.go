package play

import (
   "41.neocities.org/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
   "strings"
)

func (a Auth) Acquire(check Checkin, id string) error {
   value := protobuf.Message{
      {1, protobuf.Message{
         {1, protobuf.Message{
            {1, protobuf.Bytes(id)},
            {2, protobuf.Varint(1)},
            {3, protobuf.Varint(3)},
         }},
         {2, protobuf.Varint(1)},
      }},
      {13, protobuf.Varint(1)},
   }
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/acquire",
      bytes.NewReader(value.Marshal()),
   )
   if err != nil {
      return err
   }
   authorization(req, a)
   x_dfe_device_id(req, check)
   // with a new device, this needs to be included in the first request to
   // /fdfe/acquire, or you get:
   // Please open my apps to establish a connection with the server.
   // on following requests you can omit it
   err = x_ps_rh(req, check)
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
   value = nil
   err = value.Unmarshal(data)
   if err != nil {
      return err
   }
   for value := range value.Get(1) {
      for value := range value.Get(94) {
         for value := range value.Get(1) {
            for value := range value.Get(2) {
               for value := range value.Get(147291249) {
                  return acquire_error{value}
               }
            }
         }
      }
   }
   return nil
}

func (a acquire_error) Error() string {
   var data strings.Builder
   for m := range a[0].Get(1) {
      for m := range m.Get(10) {
         for m := range m.Get(1) {
            for b := range m.GetBytes(1) {
               if data.Len() >= 1 {
                  data.WriteByte('\n')
               }
               data.Write(b)
            }
         }
      }
   }
   return data.String()
}

type acquire_error [1]protobuf.Message
