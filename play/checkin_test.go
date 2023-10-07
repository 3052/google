package play

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func checkin_create(id int64) error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   home += "/google/play/"
   Phone.Platform = Platforms[id]
   data, err := New_Checkin(Phone)
   if err != nil {
      return err
   }
   os.WriteFile(home + Phone.Platform + ".bin", data, 0666)
   fmt.Println("Sleep(Second)")
   time.Sleep(time.Second)
   var head Header
   head.Set_Agent(false)
   if err := head.Set_Device(data); err != nil {
      return err
   }
   return head.Sync(Phone)
}

func Test_Checkin_X86(t *testing.T) {
   err := checkin_create(0)
   if err != nil {
      t.Fatal(err)
   }
}

func Test_Checkin_ARMEABI(t *testing.T) {
   err := checkin_create(1)
   if err != nil {
      t.Fatal(err)
   }
}

func Test_Checkin_ARM64(t *testing.T) {
   err := checkin_create(2)
   if err != nil {
      t.Fatal(err)
   }
}
