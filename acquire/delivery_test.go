package acquire

import (
   "154.pages.dev/google/aurora"
   "154.pages.dev/google/play"
   "154.pages.dev/http/option"
   "os"
   "testing"
   "time"
)

/*
Authorization aurora
X-Dfe-Device-Id mine
pass

Authorization aurora
X-Dfe-Device-Id studio
fail

Authorization mine
X-Dfe-Device-Id studio
fail

Authorization mine
X-Dfe-Device-Id aurora
fail
*/
func Test_Delivery(t *testing.T) {
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
   var head play.Header
   head.Authorization = func() (string, string) {
      return "Authorization", "Bearer " + auth.Auth_Token
   }
   option.No_Location()
   //option.Trace()
   option.Verbose()
   for _, app := range x86_apps {
      err := Delivery(&head, app.doc, app.version)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

var x86_apps = []struct{
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
