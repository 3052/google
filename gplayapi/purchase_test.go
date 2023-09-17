package gplayapi

import (
   "154.pages.dev/google/aurora"
   "154.pages.dev/http/option"
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Acquire(t *testing.T) {
   var auth aurora.Aurora_OSS
   {
      b, err := os.ReadFile("aurora.json")
      if err != nil {
         t.Fatal(err)
      }
      if err := auth.Unmarshal(b); err != nil {
         t.Fatal(err)
      }
   }
   client := GooglePlayClient{
      AuthData: &AuthData{
         AuthToken: auth.Auth_Token,
         GsfID: "35f8806eb7f8a047",
      },
      DeviceInfo: Pixel3a,
   }
   option.No_Location()
   option.Trace()
   for _, test := range x86_tests {
      deliver, err := client.GetDeliveryResponse(test.doc, test.version)
      if err != nil {
         t.Fatal(err)
      }
      if deliver.AppDeliveryData == nil {
         fmt.Println("ACQUIRE NEEDED, TRYING NOW:")
         time.Sleep(time.Second)
         err := client.AuthData.Acquire(test.doc, uint64(test.version))
         if err != nil {
            t.Fatal(err)
         }
         fmt.Println("sleep")
         time.Sleep(9 * time.Second)
         deliver, err := client.GetDeliveryResponse(test.doc, test.version)
         if err != nil {
            t.Fatal(err)
         }
         if deliver.AppDeliveryData != nil {
            fmt.Println("ACQUIRE SUCCESS")
         } else {
            fmt.Println("ACQUIRE FAILURE")
         }
         return
      }
   }
   fmt.Println("all apps acquired, run auth again")
}

func Test_Auth(t *testing.T) {
   option.No_Location()
   option.Verbose()
   text, err := new(aurora.Aurora_OSS).Marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("aurora.json", text, 0666)
}

var x86_tests = []struct{
   doc string
   version int
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
