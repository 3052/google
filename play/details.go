package play

import (
   "154.pages.dev/protobuf"
   "fmt"
   "io"
   "net/http"
)

func (h Header) Details(doc string) (*Details, error) {
   req, err := http.NewRequest(
      "GET", "https://android.clients.google.com/fdfe/details?doc=" + doc, nil,
   )
   if err != nil {
      return nil, err
   }
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
   response_wrapper, err := protobuf.Consume(data) // ResponseWrapper
   if err != nil {
      return nil, err
   }
   mes, ok := response_wrapper.Message(1) // Payload payload
   if !ok {
      return nil, fmt.Errorf("payload not found")
   }
   mes, _ = mes.Message(2) // Details.DetailsResponse detailsResponse
   mes, _ = mes.Message(4) // DocV2 docV2
   return &Details{mes}, nil
}

