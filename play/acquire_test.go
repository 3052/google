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
   var acq Acquire
   acq.Token, err = func() (Access_Token, error) {
      b, err := os.ReadFile(home + "/google/play/token.txt")
      if err != nil {
         return nil, err
      }
      m, err := Raw_Refresh_Token.Refresh_Token(b)
      if err != nil {
         return nil, err
      }
      return m.Auth()
   }()
   if err != nil {
      return err
   }
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
