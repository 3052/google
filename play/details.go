package play

import (
   "154.pages.dev/protobuf"
   "154.pages.dev/text"
   "errors"
   "fmt"
   "io"
   "net/http"
   "strings"
)

func (g *GoogleAuth) Details(
   checkin *GoogleCheckin, doc string, single bool,
) (*Details, error) {
   req, err := http.NewRequest("", "https://android.clients.google.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/fdfe/details"
   req.URL.RawQuery = "doc=" + doc
   authorization(req, g)
   user_agent(req, single)
   err = x_dfe_device_id(req, checkin)
   if err != nil {
      return nil, err
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   message := protobuf.Message{}
   if err = message.Unmarshal(data); err != nil {
      return nil, err
   }
   message, _ = message.Get(1)()
   message, _ = message.Get(2)()
   message, _ = message.Get(4)()
   return &Details{message}, nil
}

///

func (d *Details) Downloads() (uint64, bool) {
   v := <-d.Message.Get(13)
   v = <-v.Get(1)
   if v, ok := <-v.GetVarint(70); ok {
      return uint64(v), true
   }
   return 0, false
}

func (d *Details) field_13_1_17() chan uint64 {
   v := <-d.Message.Get(13)
   v = <-v.Get(1)
   vs := make(chan uint64)
   go func() {
      for v := range v.Get(17) {
         if v, ok := <-v.GetVarint(1); ok {
            vs <- uint64(v)
         }
      }
      close(vs)
   }()
   return vs
}

func (d *Details) field_13_1_12() (string, bool) {
   v := <-d.Message.Get(13)
   v = <-v.Get(1)
   if v, ok := <-v.GetBytes(12); ok {
      return string(v), true
   }
   return "", false
}

type Details struct {
   Message protobuf.Message
}

func (d *Details) Name() (string, bool) {
   if v, ok := <-d.Message.GetBytes(5); ok {
      return string(v), true
   }
   return "", false
}

func (d *Details) field_13_1_16() (string, bool) {
   v := <-d.Message.Get(13)
   v = <-v.Get(1)
   if v, ok := <-v.GetBytes(16); ok {
      return string(v), true
   }
   return "", false
}

func (d *Details) field_8_1() (float64, bool) {
   v := <-d.Message.Get(8)
   if v, ok := <-v.GetVarint(1); ok {
      return float64(v) / 1_000_000, true
   }
   return 0, false
}

func (d *Details) field_13_1_4() (string, bool) {
   v := <-d.Message.Get(13)
   v = <-v.Get(1)
   if v, ok := <-v.GetBytes(4); ok {
      return string(v), true
   }
   return "", false
}

func (d *Details) field_8_2() (string, bool) {
   v := <-d.Message.Get(8)
   if v, ok := <-v.GetBytes(2); ok {
      return string(v), true
   }
   return "", false
}

func (d *Details) size() (uint64, bool) {
   v := <-d.Message.Get(13)
   v = <-v.Get(1)
   if v, ok := <-v.GetVarint(9); ok {
      return uint64(v), true
   }
   return 0, false
}

func (d *Details) version_code() (uint64, bool) {
   v := <-d.Message.Get(13)
   v = <-v.Get(1)
   if v, ok := <-v.GetVarint(3); ok {
      return uint64(v), true
   }
   return 0, false
}

func (d *Details) field_13_1_82_1_1() (string, bool) {
   v := <-d.Message.Get(13)
   v = <-v.Get(1)
   v = <-v.Get(82)
   v = <-v.Get(1)
   if v, ok := <-v.GetBytes(1); ok {
      return string(v), true
   }
   return "", false
}

func (d *Details) String() string {
   var b []byte
   b = append(b, "details[8] ="...)
   if value, ok := d.field_8_1(); ok {
      b = fmt.Append(b, " ", value)
   }
   if value, ok := d.field_8_2(); ok {
      b = fmt.Append(b, " ", value)
   }
   b = append(b, "\ndetails[13][1][4] ="...)
   if value, ok := d.field_13_1_4(); ok {
      b = fmt.Append(b, " ", value)
   }
   b = append(b, "\ndetails[13][1][12] ="...)
   if value, ok := d.field_13_1_12(); ok {
      b = fmt.Append(b, " ", value)
   }
   b = append(b, "\ndetails[13][1][16] ="...)
   if value, ok := d.field_13_1_16(); ok {
      b = fmt.Append(b, " ", value)
   }
   b = append(b, "\ndetails[13][1][17] ="...)
   for file := range d.field_13_1_17() {
      if file >= 1 {
         b = append(b, " OBB"...)
      } else {
         b = append(b, " APK"...)
      }
   }
   b = append(b, "\ndetails[13][1][82][1][1] ="...)
   if value, ok := d.field_13_1_82_1_1(); ok {
      b = fmt.Append(b, " ", value)
   }
   b = append(b, "\ndownloads ="...)
   if value, ok := d.Downloads(); ok {
      b = fmt.Append(b, " ", text.Cardinal(value))
   }
   b = append(b, "\nname ="...)
   if value, ok := d.Name(); ok {
      b = fmt.Append(b, " ", value)
   }
   b = append(b, "\nsize ="...)
   if value, ok := d.size(); ok {
      b = fmt.Append(b, " ", text.Size(value))
   }
   b = append(b, "\nversion code ="...)
   if value, ok := d.version_code(); ok {
      b = fmt.Append(b, " ", value)
   }
   return string(b)
}
