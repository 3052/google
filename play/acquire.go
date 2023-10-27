package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

func (h Header) Acquire(doc string) error {
   var m protobuf.Message
   m.Add(1, func(m *protobuf.Message) {
      m.Add(1, func(m *protobuf.Message) {
         m.Add_String(1, doc)
         m.Add_Varint(2, 1)
         m.Add_Varint(3, 3)
      })
      m.Add_Varint(2, 1)
   })
   m.Add_Varint(13, 1)
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/acquire",
      bytes.NewReader(m.Append(nil)),
   )
   if err != nil {
      return err
   }
   req.Header.Set(h.Authorization())
   req.Header.Set(h.X_DFE_Device_ID())
   // with a new device, this needs to be included in the first request to
   // /fdfe/acquire, or you get:
   // Please open my apps to establish a connection with the server.
   // on following requests you can omit it
   req.Header.Set(h.X_PS_RH())
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   m, err = func() (protobuf.Message, error) {
      b, err := io.ReadAll(res.Body)
      if err != nil {
         return nil, err
      }
      return protobuf.Consume(b)
   }()
   if err != nil {
      return err
   }
   m.Message(1)
   m.Message(94)
   m.Message(1)
   m.Message(2)
   if m.Message(147291249) {
      return acquire_error{m}
   }
   return nil
}

type acquire_error struct {
   m protobuf.Message
}

func (a acquire_error) Error() string {
   var b []byte
   for _, f := range a.m {
      if f.Number == 1 {
         if m, ok := f.Message(); ok {
            m.Message(10)
            m.Message(1)
            if c, ok := m.Bytes(1); ok {
               if b != nil {
                  b = append(b, '\n')
               }
               b = append(b, c...)
            }
         }
      }
   }
   return string(b)
}
