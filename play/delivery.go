package play

import (
   "154.pages.dev/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

type Apk struct {
   Message protobuf.Message
}

func (g *GoogleAuth) Delivery(
   checkin *GoogleCheckin, app *StoreApp, single bool,
) (*Delivery, error) {
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
   if err = x_dfe_device_id(req, checkin); err != nil {
      return nil, err
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   message := protobuf.Message{}
   if err = message.Unmarshal(data); err != nil {
      return nil, err
   }
   message, _ = message.Get(1)()
   message, _ = message.Get(21)()
   switch v, _ := message.GetVarint(1)(); v {
   case 3:
      return nil, errors.New("acquire")
   case 5:
      return nil, errors.New("version")
   }
   message, _ = message.Get(2)()
   return &Delivery{message}, nil
}

func (a *Apk) Url() (string, bool) {
   v, ok := a.Message.GetBytes(5)()
   return string(v), ok
}

func (a *Apk) Field1() (string, bool) {
   v, ok := a.Message.GetBytes(1)()
   return string(v), ok
}

type Delivery struct {
   Message protobuf.Message
}

type Obb struct {
   Message protobuf.Message
}

func (d *Delivery) Url() (string, bool) {
   v, ok := d.Message.GetBytes(3)()
   return string(v), ok
}

func (o *Obb) Url() (string, bool) {
   v, ok := o.Message.GetBytes(4)()
   return string(v), ok
}

func (o *Obb) Field1() (uint64, bool) {
   v, ok := o.Message.GetVarint(1)()
   return uint64(v), ok
}

func (d *Delivery) Obb() func() (*Obb, bool) {
   next := d.Message.Get(4)
   return func() (*Obb, bool) {
      m, ok := next()
      return &Obb{m}, ok
   }
}

func (d *Delivery) Apk() func() (*Apk, bool) {
   next := d.Message.Get(15)
   return func() (*Apk, bool) {
      m, ok := next()
      return &Apk{m}, ok
   }
}
