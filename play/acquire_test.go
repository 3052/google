package play

import (
   "154.pages.dev/http"
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Acquire(t *testing.T) {
   http.No_Location()
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google/play/"
   var token play.Refresh_Token
   token.Raw, err = os.ReadFile(home + "token.txt")
   if err != nil {
      return err
   }
   token.Unmarshal()
   var acq play.Acquire
   acq.Token.Refresh(token)
   time.Sleep(time.Second)
   for _, app := range apps {
      acq.Checkin, err = func() (*play.Checkin, error) {
         s := func() string {
            var b strings.Builder
            b.WriteString(home)
            b.WriteString("/google/play/")
            b.WriteString(play.Platforms[f.platform])
            b.WriteString(".bin")
            return b.String()
         }()
         b, err := os.ReadFile(s)
         if err != nil {
            return nil, err
         }
         return play.Raw_Checkin.Checkin(b)
      }()
      if err != nil {
         return err
      }
      fmt.Println(app.doc)
      if err := acq.Acquire(f.app.ID); err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
