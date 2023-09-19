package acquire

import (
   "154.pages.dev/google/play"
   "154.pages.dev/protobuf"
   "errors"
   //"fmt"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

func Delivery(h *play.Header, doc string, version uint64) error {
   req := new(http.Request)
   req.Method = "GET"
   req.Header = make(http.Header)
   req.Header["User-Agent"] = []string{"Android-Finsky/22.8.44-21%20%5B0%5D%20%5BPR%5D%20342964500 (api=3,versionCode=82284410,sdk=27,device=generic_x86,hardware=ranchu,product=sdk_gphone_x86,platformVersionRelease=8.1.0,model=Android%20SDK%20built%20for%20x86,buildId=OSM1.180201.037,isWideScreen=0,supportedAbis=x86)"}
   req.Header["X-Dfe-Device-Id"] = []string{device_ID}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "play-fe.googleapis.com"
   req.URL.Path = "/fdfe/delivery"
   val := make(url.Values)
   req.URL.Scheme = "https"
   req.Header.Set(h.Authorization())
   val["doc"] = []string{doc}
   val["vc"] = []string{strconv.FormatUint(version, 10)}
   //val["da"] = []string{"3"}
   //val["fdcf"] = []string{"1", "2"}
   //val["ia"] = []string{"false"}
   //val["isid"] = []string{"oLHwp38pRtqEky5tTkuimg"}
   //val["ot"] = []string{"1"}
   //val["st"] = []string{"EMbEnagGGdo8npHYQdlB"}
   req.URL.RawQuery = val.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   mes, err := protobuf.Consume(data) // ResponseWrapper
   if err != nil {
      return err
   }
   //fmt.Printf("%#v\n", mes)
   mes, ok := mes.Message(1)
   if !ok {
      return errors.New("payload")
   }
   mes, ok = mes.Message(21)
   if !ok {
      return errors.New("deliveryResponse")
   }
   status, ok := mes.Varint(1)
   if ok {
      switch status {
      case 3:
         return errors.New("acquire")
      case 5:
         return errors.New("version")
      }
   }
   _, ok = mes.Message(2)
   if !ok {
      return errors.New("appDeliveryData")
   }
   return nil
}
