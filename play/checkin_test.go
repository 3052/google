package play

import (
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
   // checkin
   data, err := Checkin()
   if err != nil {
      return err
   }
   // upload
   var head Header
   head.Set_Agent(false)
   {
      b, err := os.ReadFile(home + "token.txt")
      if err != nil {
         return err
      }
      head.Set_Authorization(b)
   }
   Phone.Platform = Platforms[id]
   {
      b, err := os.ReadFile(home + Phone.Platform + ".bin")
      if err != nil {
         return err
      }
      head.Set_Device(b)
   }
   if err := head.upload_device(Phone); err != nil {
      return err
   }
   // write
   os.WriteFile(home + Phone.Platform + ".bin", data, 0666)
   time.Sleep(9*time.Second)
   return nil
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
