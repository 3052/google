package play

import (
   "154.pages.dev/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

///

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
   body, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   message := protobuf.Message{}
   if err = message.Unmarshal(body); err != nil {
      return nil, err
   }
   unknown := <-message.GetUnknown(1)
   unknown = <-unknown.Get(21)
   switch <-unknown.GetVarint(1) {
   case 3:
      return nil, errors.New("acquire")
   case 5:
      return nil, errors.New("version")
   }
   unknown = <-unknown.Get(2)
   return &Delivery{unknown}, nil
}

func (a *Apk) Url() (string, bool) {
   if v, ok := <-a.Message.GetBytes(5); ok {
      return string(v), true
   }
   return "", false
}

func (a *Apk) Field1() (string, bool) {
   if v, ok := <-a.Message.GetBytes(1); ok {
      return string(v), true
   }
   return "", false
}

type Apk struct {
   Message protobuf.Message
}

func (o *Obb) Url() (string, bool) {
   if v, ok := <-o.Message.GetBytes(4); ok {
      return string(v), true
   }
   return "", false
}

func (o *Obb) Field1() (uint64, bool) {
   if v, ok := <-o.Message.GetVarint(1); ok {
      return uint64(v), true
   }
   return 0, false
}

func (d *Delivery) Obb() chan Obb {
   vs := make(chan Obb)
   go func() {
      for v := range d.Message.Get(4) {
         vs <- Obb{v}
      }
      close(vs)
   }()
   return vs
}

func (d *Delivery) Apk() chan Apk {
   vs := make(chan Apk)
   go func() {
      for v := range d.Message.Get(15) {
         vs <- Apk{v}
      }
      close(vs)
   }()
   return vs
}

func (d *Delivery) Url() (string, bool) {
   if v, ok := <-d.Message.GetBytes(3); ok {
      return string(v), true
   }
   return "", false
}

type Delivery struct {
   Message protobuf.Message
}

type Obb struct {
   Message protobuf.Message
}
