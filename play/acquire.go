package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

func (a Acquire) Do(app string) error {
   var m protobuf.Message
   m.Add(1, func(m *protobuf.Message) {
      m.Add(1, func(m *protobuf.Message) {
         m.AddBytes(1, []byte(app))
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
   authorization(req, a.Token)
   if err := x_dfe_device_id(req, a.Checkin); err != nil {
      return err
   }
   // with a new device, this needs to be included in the first request to
   // /fdfe/acquire, or you get:
   // Please open my apps to establish a connection with the server.
   // on following requests you can omit it
   if err := x_ps_rh(req, a.Checkin); err != nil {
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
   m = nil
   if err := m.Consume(data); err != nil {
      return err
   }
   if v, ok := m.Get(1); ok {
      if v, ok := v.Get(94); ok {
         if v, ok := v.Get(1); ok {
            if v, ok := v.Get(2); ok {
               if v, ok := v.Get(147291249); ok {
                  return acquire_error{v}
               }
            }
         }
      }
   }
   return nil
}

type acquire_error struct {
   m protobuf.Message
}

type Acquire struct {
   Checkin Checkin
   Token AccessToken
}

func (a acquire_error) Error() string {
   var b []byte
   for {
      i, v := a.m.Index(1)
      if i == -1 {
         break
      }
      if v, ok := v.Get(10); ok {
         if v, ok := v.Get(1); ok {
            if v, ok := v.GetBytes(1); ok {
               if b != nil {
                  b = append(b, '\n')
               }
               b = append(b, v...)
            }
         }
      }
      a.m = a.m[i+1:]
   }
   return string(b)
}
