package play

import (
   "fmt"
   "os"
   "reflect"
   "testing"
   "time"
)

func TestSize(t *testing.T) {
   a := reflect.TypeOf(&struct{}{}).Size()
   for _, test := range size_tests {
      b := reflect.TypeOf(test).Size()
      if b > a {
         fmt.Printf("*%T\n", test)
      } else {
         fmt.Printf("%T\n", test)
      }
   }
}

var size_tests = []any{
   Apk{},
   GoogleAuth{},
   GoogleCheckin{},
   GoogleDelivery{},
   GoogleDetails{},
   GoogleDevice{},
   GoogleToken{},
   Obb{},
   StoreApp{},
   Values{},
   acquire_error{},
}

func TestSync(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google-play"
   for _, abi := range Abis {
      fmt.Println(abi)
      Device.Abi = abi
      time.Sleep(time.Second)
      var checkin GoogleCheckin
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
      err = Device.Sync(checkin)
      if err != nil {
         t.Fatal(err)
      }
   }
}
