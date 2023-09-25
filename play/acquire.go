package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "fmt"
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
      m.Add_Bytes(3, nil)
   })
   m.Add_Bytes(8, nil)
   m.Add_Varint(13, 1)
   m.Add_Varint(15, 0)
   m.Add_String(22, "nonce=qlIhuESfWlLW7GI7k6fWej7z403Mynf3o0dl5B9RYfWxmHxGGSGqBARF_TxpzL5RfVPW3oFX0zAHhSETtuUBv7TvrzWOx5hEgolPjFDs1lr_Po9lyH1HJ8UskVSkyMe_gImmY0-hA-I0SSaBfUXyInciRuuMtSNiXsMclJwWoW1bPgjYsQKCn5szQnDPlMvSqz4hBCbIGxGKiWe6L9f6IcmfIwlz8eSRQUA02YN633zvXDbptBIKfrpwE9_P_N0sWrOhc3k9LAQlrn4f4kXor4g98ZQ6BN6U3us7US-2ES-xiCaFvdlbIMiMvpp7_AsnLon1KyvxS_rujvoDaUyOOQ")
   m.Add_Varint(25, 2)
   m.Add(30, func(m *protobuf.Message) {
      m.Add_Varint(1, 2)
   })
   req, err := http.NewRequest(
      "POST", "https://play-fe.googleapis.com/fdfe/acquire",
      bytes.NewReader(m.Append(nil)),
   )
   if err != nil {
      return err
   }
   //req.Header.Set("Accept-Language", "en-US")
   //req.Header.Set("Connection", "Keep-Alive")
   //req.Header.Set("Content-Type", "application/x-protobuf")
   //req.Header.Set("X-Dfe-Client-Id", "am-unknown")
   //req.Header.Set("X-Dfe-Mccmnc", "310260")
   //req.Header.Set("X-Dfe-Network-Type", "4")
   //req.Header.Set("X-Dfe-Request-Params", "timeoutMs=35000")
   //req.URL.RawQuery = "theme=2"
   req.Header.Set(h.Authorization())
   req.Header.Set(h.Device_Config())
   req.Header.Set(h.Device_ID())
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return fmt.Errorf(res.Status)
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
   m, _ = m.Message(1)
   m, _ = m.Message(94)
   m, _ = m.Message(1)
   m, _ = m.Message(2)
   if m, ok := m.Message(175996169); ok {
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
      if f.Number == 2 {
         if m, ok := f.Message(); ok {
            m, _ = m.Message(20)
            m, _ = m.Message(1)
            c, _ := m.Bytes(1)
            if b != nil {
               b = append(b, '\n')
            }
            b = append(b, c...)
         }
      }
   }
   return string(b)
}
