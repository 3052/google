package play

import (
   "41.neocities.org/protobuf"
   "41.neocities.org/x/stringer"
   "errors"
   "fmt"
   "io"
   "iter"
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
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
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
   for value := range value.Get(1) {
      for value := range value.Get(2) {
         for value := range value.Get(4) {
            return &Details{value}, nil
         }
      }
   }
   return nil, errors.New(req.URL.Path)
}

type Details [1]protobuf.Message

func (d Details) Name() string {
   for data := range d[0].GetBytes(5) {
      return string(data)
   }
   return ""
}

// com.google.android.youtube.tvkids
func (d Details) field_15_18() string {
   for data := range d[0].Get(15) {
      for data := range data.GetBytes(18) {
         return string(data)
      }
   }
   return ""
}

func (d Details) field_8_1() float64 {
   for data := range d[0].Get(8) {
      for data := range data.GetVarint(1) {
         return float64(data) / 1_000_000
      }
   }
   return 0
}

func (d Details) field_8_2() string {
   for data := range d[0].Get(8) {
      for data := range data.GetBytes(2) {
         return string(data)
      }
   }
   return ""
}

func (d Details) Downloads() uint64 {
   for data := range d[0].Get(13) {
      for data := range data.Get(1) {
         for data := range data.GetVarint(70) {
            return uint64(data)
         }
      }
   }
   return 0
}

func (d Details) field_13_1_4() string {
   for data := range d[0].Get(13) {
      for data := range data.Get(1) {
         for data := range data.GetBytes(4) {
            return string(data)
         }
      }
   }
   return ""
}

func (d Details) field_13_1_16() string {
   for data := range d[0].Get(13) {
      for data := range data.Get(1) {
         for data := range data.GetBytes(16) {
            return string(data)
         }
      }
   }
   return ""
}

func (d Details) size() uint64 {
   for data := range d[0].Get(13) {
      for data := range data.Get(1) {
         for data := range data.GetVarint(9) {
            return uint64(data)
         }
      }
   }
   return 0
}

func (d Details) version_code() uint64 {
   for data := range d[0].Get(13) {
      for data := range data.Get(1) {
         for data := range data.GetVarint(3) {
            return uint64(data)
         }
      }
   }
   return 0
}

func (d Details) field_13_1_82_1_1() string {
   for data := range d[0].Get(13) {
      for data := range data.Get(1) {
         for data := range data.Get(82) {
            for data := range data.Get(1) {
               for data := range data.GetBytes(1) {
                  return string(data)
               }
            }
         }
      }
   }
   return ""
}

func (d Details) field_13_1_17() iter.Seq[uint64] {
   return func(yield func(uint64) bool) {
      for data := range d[0].Get(13) {
         for data := range data.Get(1) {
            for data := range data.Get(17) {
               for data := range data.GetVarint(1) {
                  if !yield(uint64(data)) {
                     return
                  }
               }
            }
         }
      }
   }
}

func (d Details) String() string {
   var b []byte
   b = fmt.Appendln(b, "details[8] =", d.field_8_1(), d.field_8_2())
   b = fmt.Appendln(b, "details[13][1][4] =", d.field_13_1_4())
   b = fmt.Appendln(b, "details[13][1][16] =", d.field_13_1_16())
   b = append(b, "details[13][1][17] ="...)
   for data := range d.field_13_1_17() {
      if data >= 1 {
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
