package play

import (
   "41.neocities.org/protobuf"
   "errors"
   "io"
   "iter"
   "net/http"
   "net/url"
   "strconv"
)

type Delivery [1]protobuf.Message

type Apk [1]protobuf.Message

type Obb [1]protobuf.Message

func (a Auth) Delivery(
   check Checkin, app1 *App, single bool,
) (*Delivery, error) {
   req, _ := http.NewRequest("", "https://android.clients.google.com", nil)
   req.URL.Path = "/fdfe/delivery"
   req.URL.RawQuery = url.Values{
      "doc": {app1.Id},
      "vc":  {strconv.FormatUint(app1.Version, 10)},
   }.Encode()
   authorization(req, a)
   user_agent(req, single)
   x_dfe_device_id(req, check)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   var value protobuf.Message
   err = value.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   for value := range value.Get(1) {
      for value := range value.Get(21) {
         for err := range value.GetVarint(1) {
            switch err {
            case 2:
               return nil, errors.New("version")
            case 3:
               return nil, errors.New("acquire")
            }
         }
         for value := range value.Get(2) {
            return &Delivery{value}, nil
         }
      }
   }
   return nil, errors.New(req.URL.Path)
}

func (d Delivery) Url() string {
   for data := range d[0].GetBytes(3) {
      return string(data)
   }
   return ""
}

func (d Delivery) Obb() iter.Seq[Obb] {
   return func(yield func(Obb) bool) {
      for message := range d[0].Get(4) {
         if !yield(Obb{message}) {
            break
         }
      }
   }
}

func (d Delivery) Apk() iter.Seq[Apk] {
   return func(yield func(Apk) bool) {
      for message := range d[0].Get(15) {
         if !yield(Apk{message}) {
            break
         }
      }
   }
}

func (a Apk) Field1() string {
   for data := range a[0].GetBytes(1) {
      return string(data)
   }
   return ""
}

func (a Apk) Url() string {
   for data := range a[0].GetBytes(5) {
      return string(data)
   }
   return ""
}

func (o Obb) Field1() uint64 {
   for data := range o[0].GetVarint(1) {
      return uint64(data)
   }
   return 0
}

func (o Obb) Url() string {
   for data := range o[0].GetBytes(4) {
      return string(data)
   }
   return ""
}
