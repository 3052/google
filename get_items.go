package googleplay

import (
   "2a.pages.dev/rosso/protobuf"
   "bytes"
   "encoding/base64"
   "net/http"
   "os"
)

func (h Header) Get_Items(app string) (*Response, error) {
   body := protobuf.Message{
      2: protobuf.Message{
         1: protobuf.Message{
            1: protobuf.Message{
               1: protobuf.String(app),
            },
         },
      },
   }.Marshal()
   req, err := http.NewRequest(
      "POST", "https://play-fe.googleapis.com/fdfe/getItems",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   h.Set_Device(req.Header)
   field := protobuf.Message{
      // valid range 0xC0 - 0xFFFFFFFF
      3: protobuf.Bytes{0xFF, 0xFF, 0xFF, 0xFF},
      // valid range 0xC0 - 0xFFFFFFFF
      4: protobuf.Bytes{0xFF, 0xFF, 0xFF, 0xFF},
   }.Marshal()
   mask := base64.StdEncoding.EncodeToString(field)
   req.Header.Set("X-Dfe-Item-Field-Mask", mask)
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   return &Response{res}, nil
}

type Items struct {
   protobuf.Message
}

func Open_Items(name string) (*Items, error) {
   data, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   message, err := protobuf.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   message = message.Get(11).Get(2)
   return &Items{message}, nil
}

func (i Items) Num_Downloads() (uint64, error) {
   return i.Get(3).Get(8).Get_Varint(3)
}

func (i Items) Version_Code() (uint64, error) {
   return i.Get(3).Get(2).Get_Varint(1)
}

func (i Items) Version() (string, error) {
   return i.Get(3).Get(2).Get_String(2)
}

func (i Items) Upload_Date() (string, error) {
   return i.Get(3).Get(9).Get_String(2)
}

func (i Items) Title() (string, error) {
   return i.Get(2).Get(1).Get_String(1)
}

func (i Items) Category() (string, error) {
   return i.Get(2).Get(30).Get_String(1)
}

func (i Items) Creator() (string, error) {
   return i.Get(3).Get(14).Get_String(1)
}

func (i Items) Offer() (string, error) {
   return i.Get(2).Get(10).Get(1).Get(1).Get(2).Get(1).Get_String(2)
}
