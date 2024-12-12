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
   data, err := os.ReadFile(home + "/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token GoogleToken
   err = token.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   auth, err := token.Auth()
   if err != nil {
      t.Fatal(err)
   }
   time.Sleep(time.Second)
   for _, app := range apps {
      data, err = os.ReadFile(home + "/" + app.abi + ".txt")
      if err != nil {
         t.Fatal(err)
      }
      var checkin GoogleCheckin
      err = checkin.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(app.id)
      err = auth.Acquire(checkin, app.id)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
