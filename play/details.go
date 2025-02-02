package play

import (
   "41.neocities.org/protobuf"
   "41.neocities.org/x/strconv"
   "errors"
   "fmt"
   "io"
   "net/http"
   "strings"
)

func (a Auth) Details(check Checkin, doc string, single bool) (*Details, error) {
   req, _ := http.NewRequest("", "https://android.clients.google.com", nil)
   req.URL.Path = "/fdfe/details"
   req.URL.RawQuery = "doc=" + doc
   authorization(req, a)
   user_agent(req, single)
   x_dfe_device_id(req, check)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var data strings.Builder
      resp.Write(&data)
      return nil, errors.New(data.String())
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
   message, _ = message.Get(2)()
   message, _ = message.Get(4)()
   return &Details{message}, nil
}

type Details struct {
   Message protobuf.Message
}

func (d Details) String() string {
   var b []byte
   b = fmt.Appendln(b, "details[8] =", d.field_8_1(), d.field_8_2())
   b = fmt.Appendln(b, "details[13][1][4] =", d.field_13_1_4())
   b = fmt.Appendln(b, "details[13][1][16] =", d.field_13_1_16())
   b = append(b, "details[13][1][17] ="...)
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
   b = append(b, '\n')
   b = fmt.Appendln(b, "details[13][1][82][1][1] =", d.field_13_1_82_1_1())
   b = fmt.Appendln(b, "details[15][18] =", d.field_15_18())
   b = fmt.Appendln(
      b, "downloads =", strconv.Cardinal(d.Downloads()),
   )
   b = fmt.Appendln(b, "name =", d.Name())
   b = fmt.Appendln(
      b, "size =", strconv.Size(d.size()),
   )
   b = fmt.Append(b, "version code = ", d.version_code())
   return string(b)
}

func (d Details) Name() string {
   value, _ := d.Message.GetBytes(5)()
   return string(value)
}

func (d Details) Downloads() uint64 {
   m, _ := d.Message.Get(13)()
   m, _ = m.Get(1)()
   value, _ := m.GetVarint(70)()
   return uint64(value)
}

func (d Details) field_8_1() float64 {
   m, _ := d.Message.Get(8)()
   value, _ := m.GetVarint(1)()
   return float64(value) / 1_000_000
}

func (d Details) field_8_2() string {
   m, _ := d.Message.Get(8)()
   value, _ := m.GetBytes(2)()
   return string(value)
}

func (d Details) field_13_1_4() string {
   m, _ := d.Message.Get(13)()
   m, _ = m.Get(1)()
   value, _ := m.GetBytes(4)()
   return string(value)
}

func (d Details) field_13_1_16() string {
   m, _ := d.Message.Get(13)()
   m, _ = m.Get(1)()
   value, _ := m.GetBytes(16)()
   return string(value)
}

func (d Details) field_13_1_82_1_1() string {
   m, _ := d.Message.Get(13)()
   m, _ = m.Get(1)()
   m, _ = m.Get(82)()
   m, _ = m.Get(1)()
   value, _ := m.GetBytes(1)()
   return string(value)
}

func (d Details) size() uint64 {
   m, _ := d.Message.Get(13)()
   m, _ = m.Get(1)()
   value, _ := m.GetVarint(9)()
   return uint64(value)
}

func (d Details) version_code() uint64 {
   m, _ := d.Message.Get(13)()
   m, _ = m.Get(1)()
   value, _ := m.GetVarint(3)()
   return uint64(value)
}

func (d Details) field_13_1_17() func() (uint64, bool) {
   m, _ := d.Message.Get(13)()
   m, _ = m.Get(1)()
   next := m.Get(17)
   return func() (uint64, bool) {
      m, _ = next()
      value, ok := m.GetVarint(1)()
      return uint64(value), ok
   }
}

// com.google.android.youtube.tvkids
func (d Details) field_15_18() string {
   m, _ := d.Message.Get(15)()
   value, _ := m.GetBytes(18)()
   return string(value)
}
