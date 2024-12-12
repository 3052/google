package play

import (
   "41.neocities.org/protobuf"
   "41.neocities.org/text"
   "errors"
   "fmt"
   "io"
   "net/http"
   "strings"
)

func (g GoogleAuth) Details(
   check GoogleCheckin, doc string, single bool,
) (GoogleDetails, error) {
   req, err := http.NewRequest("", "https://android.clients.google.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/fdfe/details"
   req.URL.RawQuery = "doc=" + doc
   authorization(req, g)
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
   message := protobuf.Message{}
   err = message.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   message, _ = message.Get(1)()
   message, _ = message.Get(2)()
   message, _ = message.Get(4)()
   return func() protobuf.Message {
      return message
   }, nil
}

type GoogleDetails func() protobuf.Message

func (g GoogleDetails) Downloads() uint64 {
   data, _ := g().Get(13)()
   data, _ = data.Get(1)()
   value, _ := data.GetVarint(70)()
   return uint64(value)
}

func (g GoogleDetails) Name() string {
   data, _ := g().GetBytes(5)()
   return string(data)
}

func (g GoogleDetails) String() string {
   var data []byte
   data = fmt.Appendln(data, "details[8] =", g.field_8_1(), g.field_8_2())
   data = fmt.Appendln(data, "details[13][1][4] =", g.field_13_1_4())
   data = fmt.Appendln(data, "details[13][1][16] =", g.field_13_1_16())
   data = append(data, "details[13][1][17] ="...)
   values := g.field_13_1_17()
   for {
      value, ok := values()
      if !ok {
         break
      }
      if value >= 1 {
         data = append(data, " OBB"...)
      } else {
         data = append(data, " APK"...)
      }
   }
   data = append(data, '\n')
   data = fmt.Appendln(data, "details[13][1][82][1][1] =", g.field_13_1_82_1_1())
   data = fmt.Appendln(data, "details[15][18] =", g.field_15_18())
   data = fmt.Appendln(
      data, "downloads =", text.Cardinal(g.Downloads()),
   )
   data = fmt.Appendln(data, "name =", g.Name())
   data = fmt.Appendln(
      data, "size =", text.Size(g.size()),
   )
   data = fmt.Append(data, "version code = ", g.version_code())
   return string(data)
}

func (g GoogleDetails) field_8_1() float64 {
   data, _ := g().Get(8)()
   value, _ := data.GetVarint(1)()
   return float64(value) / 1_000_000
}

func (g GoogleDetails) field_8_2() string {
   data, _ := g().Get(8)()
   value, _ := data.GetBytes(2)()
   return string(value)
}

func (g GoogleDetails) field_13_1_4() string {
   data, _ := g().Get(13)()
   data, _ = data.Get(1)()
   value, _ := data.GetBytes(4)()
   return string(value)
}

func (g GoogleDetails) field_13_1_16() string {
   data, _ := g().Get(13)()
   data, _ = data.Get(1)()
   value, _ := data.GetBytes(16)()
   return string(value)
}

func (g GoogleDetails) field_13_1_17() func() (uint64, bool) {
   value, _ := g().Get(13)()
   value, _ = value.Get(1)()
   values := value.Get(17)
   return func() (uint64, bool) {
      value, _ = values()
      data, ok := value.GetVarint(1)()
      return uint64(data), ok
   }
}

func (g GoogleDetails) field_13_1_82_1_1() string {
   data, _ := g().Get(13)()
   data, _ = data.Get(1)()
   data, _ = data.Get(82)()
   data, _ = data.Get(1)()
   value, _ := data.GetBytes(1)()
   return string(value)
}

// com.google.android.youtube.tvkids
func (g GoogleDetails) field_15_18() string {
   data, _ := g().Get(15)()
   value, _ := data.GetBytes(18)()
   return string(value)
}

func (g GoogleDetails) size() uint64 {
   data, _ := g().Get(13)()
   data, _ = data.Get(1)()
   value, _ := data.GetVarint(9)()
   return uint64(value)
}

func (g GoogleDetails) version_code() uint64 {
   data, _ := g().Get(13)()
   data, _ = data.Get(1)()
   value, _ := data.GetVarint(3)()
   return uint64(value)
}
