package play

import (
   "41.neocities.org/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

func (g GoogleAuth) Delivery(
   check GoogleCheckin, app *StoreApp, single bool,
) (*GoogleDelivery, error) {
   req, err := http.NewRequest("", "https://android.clients.google.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/fdfe/delivery"
   req.URL.RawQuery = url.Values{
      "doc": {app.Id},
      "vc":  {strconv.FormatUint(app.Version, 10)},
   }.Encode()
   authorization(req, g)
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
   return &GoogleDelivery{message}, nil
}

func (g GoogleDelivery) Url() string {
   value, _ := g.Message.GetBytes(3)()
   return string(value)
}

type GoogleDelivery struct {
   Message protobuf.Message
}

type Apk func() protobuf.Message

func (a Apk) Field1() string {
   data, _ := a().GetBytes(1)()
   return string(data)
}

func (o Obb) Url() string {
   data, _ := o().GetBytes(4)()
   return string(data)
}

func (o Obb) Field1() uint64 {
   data, _ := o().GetVarint(1)()
   return uint64(data)
}

type Obb func() protobuf.Message

func (a Apk) Url() string {
   data, _ := a().GetBytes(5)()
   return string(data)
}

func (g GoogleDelivery) Apk() func() (Apk, bool) {
   values := g.Message.Get(15)
   return func() (Apk, bool) {
      value, ok := values()
      return Apk{value}, ok
   }
}

func (g GoogleDelivery) Obb() func() (Obb, bool) {
   values := g.Message.Get(4)
   return func() (Obb, bool) {
      value, ok := values()
      return Obb{value}, ok
   }
}
