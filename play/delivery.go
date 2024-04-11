package play

import (
   "154.pages.dev/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

func (g GoogleAuth) Delivery(
   checkin GoogleCheckin, app StoreApp, single bool,
) (*Delivery, error) {
   req, err := http.NewRequest("GET", "https://android.clients.google.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/fdfe/delivery"
   req.URL.RawQuery = url.Values{
      "doc": {app.ID},
      "vc":  {strconv.FormatUint(app.Version, 10)},
   }.Encode()
   g.authorization(req)
   user_agent(req, single)
   if err := checkin.x_dfe_device_id(req); err != nil {
      return nil, err
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   var d Delivery
   if err := d.m.Consume(data); err != nil {
      return nil, err
   }
   d.m = <-d.m.Get(1)
   d.m = <-d.m.Get(21)
   switch <-d.m.GetVarint(1) {
   case 3:
      return nil, errors.New("acquire")
   case 5:
      return nil, errors.New("version code")
   }
   d.m = <-d.m.Get(2)
   return &d, nil
}

type Delivery struct {
   m protobuf.Message
}

func (d Delivery) Field3() (string, bool) {
   if v, ok := <-d.m.GetBytes(3); ok {
      return string(v), true
   }
   return "", false
}

func (d Delivery) Field4() chan DeliveryField4 {
   vs := make(chan DeliveryField4)
   go func() {
      for v := range d.m.Get(4) {
         vs <- DeliveryField4{v}
      }
      close(vs)
   }()
   return vs
}

func (d Delivery) Field15() chan DeliveryField15 {
   vs := make(chan DeliveryField15)
   go func() {
      for v := range d.m.Get(15) {
         vs <- DeliveryField15{v}
      }
      close(vs)
   }()
   return vs
}

type DeliveryField4 struct {
   m protobuf.Message
}

func (d DeliveryField4) Field1() (uint64, bool) {
   if v, ok := <-d.m.GetVarint(1); ok {
      return uint64(v), true
   }
   return 0, false
}

func (d DeliveryField4) Field4() (string, bool) {
   if v, ok := <-d.m.GetBytes(4); ok {
      return string(v), true
   }
   return "", false
}

type DeliveryField15 struct {
   m protobuf.Message
}

func (d DeliveryField15) Field1() (string, bool) {
   if v, ok := <-d.m.GetBytes(1); ok {
      return string(v), true
   }
   return "", false
}

func (d DeliveryField15) Field5() (string, bool) {
   if v, ok := <-d.m.GetBytes(5); ok {
      return string(v), true
   }
   return "", false
}
