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

func (d *Details) field_13_1_17() func() (uint64, bool) {
   m, _ := d.Message.Get(13)()
   m, _ = m.Get(1)()
   next := m.Get(17)
   return func() (uint64, bool) {
      m, _ = next()
      v, ok := m.GetVarint(1)()
      return uint64(v), ok
   }
}

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

type Details struct {
   Message protobuf.Message
}

func (d *Details) Downloads() (uint64, bool) {
   m, _ := d.Message.Get(13)()
   m, _ = m.Get(1)()
   v, ok := m.GetVarint(70)()
   return uint64(v), ok
}

func (d *Details) field_13_1_12() (string, bool) {
   m, _ := d.Message.Get(13)()
   m, _ = m.Get(1)()
   v, ok := m.GetBytes(12)()
   return string(v), ok
}

func (d *Details) Name() (string, bool) {
   v, ok := d.Message.GetBytes(5)()
   return string(v), ok
}

func (d *Details) field_13_1_16() (string, bool) {
   m, _ := d.Message.Get(13)()
   m, _ = m.Get(1)()
   v, ok := m.GetBytes(16)()
   return string(v), ok
}

func (d *Details) field_13_1_4() (string, bool) {
   m, _ := d.Message.Get(13)()
   m, _ = m.Get(1)()
   v, ok := m.GetBytes(4)()
   return string(v), ok
}

func (d *Details) field_8_2() (string, bool) {
   m, _ := d.Message.Get(8)()
   v, ok := m.GetBytes(2)()
   return string(v), ok
}

func (d *Details) size() (uint64, bool) {
   m, _ := d.Message.Get(13)()
   m, _ = m.Get(1)()
   v, ok := m.GetVarint(9)()
   return uint64(v), ok
}

func (d *Details) version_code() (uint64, bool) {
   m, _ := d.Message.Get(13)()
   m, _ = m.Get(1)()
   v, ok := m.GetVarint(3)()
   return uint64(v), ok
}

func (d *Details) field_13_1_82_1_1() (string, bool) {
   m, _ := d.Message.Get(13)()
   m, _ = m.Get(1)()
   m, _ = m.Get(82)()
   m, _ = m.Get(1)()
   v, ok := m.GetBytes(1)()
   return string(v), ok
}

func (d *Details) field_8_1() (float64, bool) {
   m, _ := d.Message.Get(8)()
   v, ok := m.GetVarint(1)()
   return float64(v) / 1_000_000, ok
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
   next := d.field_13_1_17()
   for {
      value, ok := next()
      if !ok {
         break
      }
      if value >= 1 {
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
