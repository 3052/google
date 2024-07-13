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

func (a Apk) Field1() (string, bool) {
   if v, ok := <-a.Message.GetBytes(1); ok {
      return string(v), true
   }
   return "", false
}

func (a Apk) Url() (string, bool) {
   if v, ok := <-a.Message.GetBytes(5); ok {
      return string(v), true
   }
   return "", false
}

type Delivery struct {
   Message protobuf.Message
}

func (d Delivery) Apk() chan Apk {
   vs := make(chan Apk)
   go func() {
      for v := range d.Message.Get(15) {
         vs <- Apk{v}
      }
      close(vs)
   }()
   return vs
}

func (d Delivery) Obb() chan Obb {
   vs := make(chan Obb)
   go func() {
      for v := range d.Message.Get(4) {
         vs <- Obb{v}
      }
      close(vs)
   }()
   return vs
}

func (d Delivery) Url() (string, bool) {
   if v, ok := <-d.Message.GetBytes(3); ok {
      return string(v), true
   }
   return "", false
}

type Obb struct {
   Message protobuf.Message
}

func (o Obb) Field1() (uint64, bool) {
   if v, ok := <-o.Message.GetVarint(1); ok {
      return uint64(v), true
   }
   return 0, false
}

func (o Obb) Url() (string, bool) {
   if v, ok := <-o.Message.GetBytes(4); ok {
      return string(v), true
   }
   return "", false
}

func (g GoogleCheckin) Delivery(
   auth GoogleAuth, app StoreApp, single bool,
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
   auth.authorization(req)
   user_agent(req, single)
   if err := g.x_dfe_device_id(req); err != nil {
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
   var m protobuf.Message
   err = m.Consume(data)
   if err != nil {
      return nil, err
   }
   m = <-m.Get(1)
   m = <-m.Get(21)
   switch <-m.GetVarint(1) {
   case 3:
      return nil, errors.New("acquire")
   case 5:
      return nil, errors.New("version")
   }
   m = <-m.Get(2)
   return &Delivery{m}, nil
}
