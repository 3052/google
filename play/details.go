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

type Details struct {
   Message protobuf.Message
}

func (d Details) Name() (string, bool) {
   if v, ok := <-d.Message.GetBytes(5); ok {
      return string(v), true
   }
   return "", false
}

func (d Details) String() string {
   var b []byte
   b = append(b, "details[6] ="...)
   if v, ok := d.field_6(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\ndetails[8] ="...)
   if v, ok := d.field_8_1(); ok {
      b = fmt.Append(b, " ", v)
   }
   if v, ok := d.field_8_2(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\ndetails[13][1][4] ="...)
   if v, ok := d.field_13_1_4(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\ndetails[13][1][16] ="...)
   if v, ok := d.field_13_1_16(); ok {
      b = fmt.Append(b, " ", v)
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
   if v, ok := d.field_13_1_82_1_1(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\ndownloads ="...)
   if v, ok := d.Downloads(); ok {
      b = fmt.Append(b, " ", text.Cardinal(v))
   }
   b = append(b, "\nname ="...)
   if v, ok := d.Name(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\nsize ="...)
   if v, ok := d.size(); ok {
      b = fmt.Append(b, " ", text.Size(v))
   }
   b = append(b, "\nversion code ="...)
   if v, ok := d.version_code(); ok {
      b = fmt.Append(b, " ", v)
   }
   return string(b)
}

func (d Details) field_6() (string, bool) {
   if v, ok := <-d.Message.GetBytes(6); ok {
      return string(v), true
   }
   return "", false
}

func (d Details) Downloads() (uint64, bool) {
   if v, ok := <-d.Message.Get(13); ok {
      if v, ok := <-v.Get(1); ok {
         if v, ok := v.GetVarint(70); ok {
            return uint64(v), true
         }
      }
   }
   return 0, false
}

func (d Details) field_8_1() (float64, bool) {
   if v, ok := <-d.Message.Get(8); ok {
      if v, ok := v.GetVarint(1); ok {
         return float64(v) / 1_000_000, true
      }
   }
   return 0, false
}

/////////

func (d Details) field_8_2() (string, bool) {
   d.Message = <-d.Message.Get(8)
   if v, ok := <-d.Message.GetBytes(2); ok {
      return string(v), true
   }
   return "", false
}

func (d Details) field_13_1_4() (string, bool) {
   d.Message = <-d.Message.Get(13)
   d.Message = <-d.Message.Get(1)
   if v, ok := <-d.Message.GetBytes(4); ok {
      return string(v), true
   }
   return "", false
}

func (d Details) field_13_1_16() (string, bool) {
   d.Message = <-d.Message.Get(13)
   d.Message = <-d.Message.Get(1)
   if v, ok := <-d.Message.GetBytes(16); ok {
      return string(v), true
   }
   return "", false
}

func (d Details) field_13_1_17() chan uint64 {
   vs := make(chan uint64)
   d.Message = <-d.Message.Get(13)
   d.Message = <-d.Message.Get(1)
   go func() {
      for v := range d.Message.Get(17) {
         if v, ok := <-v.GetVarint(1); ok {
            vs <- uint64(v)
         }
      }
      close(vs)
   }()
   return vs
}

func (d Details) field_13_1_82_1_1() (string, bool) {
   d.Message = <-d.Message.Get(13)
   d.Message = <-d.Message.Get(1)
   d.Message = <-d.Message.Get(82)
   d.Message = <-d.Message.Get(1)
   if v, ok := <-d.Message.GetBytes(1); ok {
      return string(v), true
   }
   return "", false
}

func (d Details) size() (uint64, bool) {
   d.Message = <-d.Message.Get(13)
   d.Message = <-d.Message.Get(1)
   if v, ok := <-d.Message.GetVarint(9); ok {
      return uint64(v), true
   }
   return 0, false
}

func (d Details) version_code() (uint64, bool) {
   d.Message = <-d.Message.Get(13)
   d.Message = <-d.Message.Get(1)
   if v, ok := <-d.Message.GetVarint(3); ok {
      return uint64(v), true
   }
   return 0, false
}

func (g GoogleCheckin) Details(
   auth GoogleAuth, doc string, single bool,
) (*Details, error) {
   req, err := http.NewRequest("", "https://android.clients.google.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/fdfe/details"
   req.URL.RawQuery = "doc=" + doc
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
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   var d Details
   if err := d.Message.Consume(data); err != nil {
      return nil, err
   }
   d.Message = <-d.Message.Get(1)
   d.Message = <-d.Message.Get(2)
   d.Message = <-d.Message.Get(4)
   return &d, nil
}
