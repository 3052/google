package play

import (
   "41.neocities.org/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

func (d Delivery) Url() string {
   value, _ := d[0].GetBytes(3)()
   return string(value)
}

func (a Apk) Field1() string {
   value, _ := a[0].GetBytes(1)()
   return string(value)
}

func (a Apk) Url() string {
   value, _ := a[0].GetBytes(5)()
   return string(value)
}

func (o Obb) Field1() uint64 {
   value, _ := o[0].GetVarint(1)()
   return uint64(value)
}

func (o Obb) Url() string {
   value, _ := o[0].GetBytes(4)()
   return string(value)
}

func (a Auth) Delivery(
   check Checkin, app0 *App, single bool,
) (*Delivery, error) {
   req, _ := http.NewRequest("", "https://android.clients.google.com", nil)
   req.URL.Path = "/fdfe/delivery"
   req.URL.RawQuery = url.Values{
      "doc": {app0.Id},
      "vc":  {strconv.FormatUint(app0.Version, 10)},
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
   message := protobuf.Message{}
   err = message.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   message, _ = message.Get(1)()
   message, _ = message.Get(21)()
   switch err, _ := message.GetVarint(1)(); err {
   case 2:
      return nil, errors.New("version")
   case 3:
      return nil, errors.New("acquire")
   }
   message, _ = message.Get(2)()
   return &Delivery{message}, nil
}

type Delivery [1]protobuf.Message

type Apk [1]protobuf.Message

type Obb [1]protobuf.Message

func (d Delivery) Obb() func() (Obb, bool) {
   next := d[0].Get(4)
   return func() (Obb, bool) {
      message, ok := next()
      return Obb{message}, ok
   }
}

func (d Delivery) Apk() func() (Apk, bool) {
   next := d[0].Get(15)
   return func() (Apk, bool) {
      message, ok := next()
      return Apk{message}, ok
   }
}
