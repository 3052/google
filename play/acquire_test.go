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
   home += "/google-play/"
   var token RefreshToken
   token.Data, err = os.ReadFile(home + "token.txt")
   if err != nil {
      t.Fatal(err)
   }
   if err := token.Unmarshal(); err != nil {
      t.Fatal(err)
   }
   var client Acquire
   if err := client.Token.Refresh(token); err != nil {
      t.Fatal(err)
   }
   time.Sleep(time.Second)
   for _, app := range apps {
      name := fmt.Sprint(home, app.platform, ".bin")
      client.Checkin.Data, err = os.ReadFile(name)
      if err != nil {
         t.Fatal(err)
      }
      if err := client.Checkin.Unmarshal(); err != nil {
         t.Fatal(err)
      }
      fmt.Println(app.id)
      if err := client.Do(app.id); err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
