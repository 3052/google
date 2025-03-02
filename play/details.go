package play

import (
   "41.neocities.org/protobuf"
   "41.neocities.org/x/stringer"
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
   var value protobuf.Message
   err = value.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   value, _ = value.Get(1)()
   value, _ = value.Get(2)()
   value, _ = value.Get(4)()
   return &Details{value}, nil
}

func (d Details) Name() string {
   value, _ := d[0].GetBytes(5)()
   return string(value)
}

func (d Details) Downloads() uint64 {
   data, _ := d[0].Get(13)()
   data, _ = data.Get(1)()
   value, _ := data.GetVarint(70)()
   return uint64(value)
}

func (d Details) field_8_1() float64 {
   data, _ := d[0].Get(8)()
   value, _ := data.GetVarint(1)()
   return float64(value) / 1_000_000
}

func (d Details) field_8_2() string {
   data, _ := d[0].Get(8)()
   value, _ := data.GetBytes(2)()
   return string(value)
}

func (d Details) field_13_1_4() string {
   data, _ := d[0].Get(13)()
   data, _ = data.Get(1)()
   value, _ := data.GetBytes(4)()
   return string(value)
}

type Details [1]protobuf.Message

func (d Details) field_13_1_16() string {
   data, _ := d[0].Get(13)()
   data, _ = data.Get(1)()
   value, _ := data.GetBytes(16)()
   return string(value)
}

func (d Details) field_13_1_82_1_1() string {
   data, _ := d[0].Get(13)()
   data, _ = data.Get(1)()
   data, _ = data.Get(82)()
   data, _ = data.Get(1)()
   value, _ := data.GetBytes(1)()
   return string(value)
}

func (d Details) size() uint64 {
   data, _ := d[0].Get(13)()
   data, _ = data.Get(1)()
   value, _ := data.GetVarint(9)()
   return uint64(value)
}

func (d Details) version_code() uint64 {
   data, _ := d[0].Get(13)()
   data, _ = data.Get(1)()
   value, _ := data.GetVarint(3)()
   return uint64(value)
}

// com.google.android.youtube.tvkids
func (d Details) field_15_18() string {
   data, _ := d[0].Get(15)()
   value, _ := data.GetBytes(18)()
   return string(value)
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
      b, "downloads =", stringer.Cardinal(d.Downloads()),
   )
   b = fmt.Appendln(b, "name =", d.Name())
   b = fmt.Appendln(
      b, "size =", stringer.Size(d.size()),
   )
   b = fmt.Append(b, "version code = ", d.version_code())
   return string(b)
}

func (d Details) field_13_1_17() func() (uint64, bool) {
   data, _ := d[0].Get(13)()
   data, _ = data.Get(1)()
   next := data.Get(17)
   return func() (uint64, bool) {
      data, _ = next()
      value, ok := data.GetVarint(1)()
      return uint64(value), ok
   }
}
