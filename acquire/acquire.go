package acquire

import (
   "154.pages.dev/encoding/protobuf"
   "154.pages.dev/google/play"
   "bytes"
   "errors"
   "io"
   "net/http"
   "strings"
)

// pass
//const Device_ID = "306e9f7f4192be79"

const Device_ID = "31eee6a47b6b29f3"

func Acquire(h *play.Header, doc string) error {
   var m protobuf.Message
   m.Add(1, func(m *protobuf.Message) {
      m.Add(1, func(m *protobuf.Message) {
         m.Add_String(1, doc)
         m.Add_Varint(2, 1)
         m.Add_Varint(3, 3)
      })
      m.Add_Varint(2, 1)
   })
   m.Add(4, func(m *protobuf.Message) {
      m.Add_Varint(1, 10)
   })
   m.Add_Varint(13, 1)
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/acquire",
      bytes.NewReader(m.Append(nil)),
   )
   if err != nil {
      return err
   }
   req.Header.Set("X-Dfe-Device-Id", Device_ID)
   req.Header.Set(h.Authorization())
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   if strings.Contains(string(data), connection) {
      return new_error(connection)
   }
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return nil
}

const connection = "Please open my apps to establish a connection with the server."

func new_error(s string) error {
   s = strings.TrimSuffix(s, ".")
   return errors.New(strings.ToLower(s))
}
