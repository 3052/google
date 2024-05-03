package play

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func TestAcquire(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google-play"
   var token GoogleToken
   token.Data, err = os.ReadFile(home + "/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   if err := token.Unmarshal(); err != nil {
      t.Fatal(err)
   }
   var auth GoogleAuth
   if err := auth.Auth(token); err != nil {
      t.Fatal(err)
   }
   time.Sleep(time.Second)
   for _, app := range apps {
      name := fmt.Sprint(home, "/", ABI[app.abi], ".bin")
      var checkin GoogleCheckin
      checkin.Data, err = os.ReadFile(name)
      if err != nil {
         t.Fatal(err)
      }
      if err := checkin.Unmarshal(); err != nil {
         t.Fatal(err)
      }
      fmt.Println(app.id)
      if err := checkin.Acquire(auth, app.id); err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
