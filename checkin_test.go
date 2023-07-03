package googleplay

import (
   "os"
   "testing"
   "time"
)

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

func checkin_create(id int64) error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   home += "/2a/googleplay/"
   res, err := Phone.Checkin(Platforms[id])
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if err := res.Write_File(home + Platforms[id] + ".bin"); err != nil {
      return err
   }
   time.Sleep(Sleep)
   return nil
}
