package main

import (
   "154.pages.dev/google/acquire"
   "154.pages.dev/google/aurora"
   "154.pages.dev/google/play"
   "154.pages.dev/http/option"
   "flag"
   "fmt"
   "os"
   "time"
)

func main() {
   flag.StringVar(&acquire.Device_ID, "id", "", "device ID")
   flag.Parse()
   if acquire.Device_ID == "" {
      flag.Usage()
      return
   }
   head := new(play.Header)
   head.Set_Agent(false)
   option.No_Location()
   option.Trace()
   {
      b, err := os.ReadFile("aurora.json")
      if err != nil {
         panic(err)
      }
      var v aurora.Aurora_OSS
      if err := v.Unmarshal(b); err != nil {
         panic(err)
      }
      head.Authorization = func() (string, string) {
         return "Authorization", "Bearer " + v.Auth_Token
      }
   }
   for _, test := range x86_tests {
      err := acquire.New_Delivery(head, test.doc, test.version)
      if err == nil {
         continue
      }
      if err.Error() != "acquire" {
         continue
      }
      fmt.Println("ACQUIRE NEEDED, TRYING NOW:")
      time.Sleep(time.Second)
      if err := acquire.Acquire(head, test.doc); err != nil {
         panic(err)
      }
      fmt.Println("sleep")
      time.Sleep(9 * time.Second)
      err = acquire.New_Delivery(head, test.doc, test.version)
      if err != nil {
         panic(err)
      }
      fmt.Println("ACQUIRE SUCCESS")
      return
   }
   fmt.Println("all apps acquired, run auth again")
}

var x86_tests = []struct{
   doc string
   version uint64
}{
   {"app.source.getcontact", 2739},
   {"br.com.rodrigokolb.realdrum", 317},
   {"com.amctve.amcfullepisodes", 28021790},
   {"com.cabify.rider", 17144463},
   {"com.clearchannel.iheartradio.controller", 710310201},
   {"com.google.android.apps.walletnfcrel", 930739964},
   {"com.google.android.youtube", 1539962304},
   {"com.instagram.android", 370010808},
   {"com.pinterest", 11338030},
   {"kr.sira.metal", 74},
   {"org.thoughtcrime.securesms", 132708},
   {"org.videolan.vlc", 13050405},
}
