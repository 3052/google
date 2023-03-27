package googleplay

import (
   "os"
   "testing"
   "time"
)

func checkin_create(id int64) error {
   platform := Platforms[id]
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   res, err := Phone.Checkin(platform)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   platform += ".bin"
   if err := res.Create(home + "/googleplay/" + platform); err != nil {
      return err
   }
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
