package play

import (
   "fmt"
   "os"
   "reflect"
   "testing"
   "time"
)

func TestSync(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google-play"
   for _, abi := range Abis {
      var checkin GoogleCheckin
      Device.Abi = abi
      fmt.Println(abi)
      time.Sleep(time.Second)
      data, err := checkin.Marshal(&Device)
      if err != nil {
         t.Fatal(err)
      }
      err = os.WriteFile(home+"/"+abi+".txt", data, os.ModePerm)
      if err != nil {
         t.Fatal(err)
      }
      err = checkin.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
      err = Device.Sync(&checkin)
      if err != nil {
         t.Fatal(err)
      }
   }
}

func TestSize(t *testing.T) {
   size := reflect.TypeOf(&struct{}{}).Size()
   for _, test := range size_tests {
      if reflect.TypeOf(test).Size() > size {
         fmt.Printf("*%T\n", test)
      } else {
         fmt.Printf("%T\n", test)
      }
   }
}

var size_tests = []any{
   Apk{},
   Delivery{},
   Details{},
   GoogleAuth{},
   GoogleCheckin{},
   GoogleDevice{},
   GoogleToken{},
   Obb{},
   StoreApp{},
   Values{},
   acquire_error{},
}
