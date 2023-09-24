package play

import (
   "154.pages.dev/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

func (h Header) Delivery(doc string, vc uint64) (*Delivery, error) {
   req, err := http.NewRequest(
      "GET", "https://play-fe.googleapis.com/fdfe/delivery", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "doc": {doc},
      "vc": {strconv.FormatUint(vc, 10)},
   }.Encode()
   req.Header.Set(h.Agent())
   req.Header.Set(h.Authorization())
   req.Header.Set(h.Device_ID())
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   mes, err := protobuf.Consume(data) // ResponseWrapper
   if err != nil {
      return nil, err
   }
   mes, _ = mes.Message(1) // payload
   mes, _ = mes.Message(21) // deliveryResponse
   status, ok := mes.Varint(1)
   if ok {
      switch status {
      case 3:
         return nil, errors.New("acquire")
      case 5:
         return nil, errors.New("version")
      }
   }
   mes, ok = mes.Message(2)
   if !ok {
      return nil, errors.New("appDeliveryData not found")
   }
   return &Delivery{mes}, nil
}

