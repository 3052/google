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
   data, err := Phone.Checkin(Platforms[id])
   if err != nil {
      return err
   }
   os.WriteFile(home + Platforms[id] + ".bin", data, 0666)
   time.Sleep(Sleep)
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
