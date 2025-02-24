package main

import (
   "41.neocities.org/google/play"
   "net/http"
   "os"
   "strings"
   xhttp "41.neocities.org/x/http"
)

func (f *flags) do_details() (*play.Details, error) {
   var checkin play.Checkin
   auth, err := f.client(&checkin)
   if err != nil {
      return nil, err
   }
   return auth.Details(checkin, f.app.Id, f.single)
}

func (f *flags) do_auth() error {
   data, err := play.Token{}.Marshal(f.token)
   if err != nil {
      return err
   }
   os.Mkdir(f.home, os.ModePerm)
   return os.WriteFile(f.home+"/token.txt", data, os.ModePerm)
}

func (f *flags) do_checkin() error {
   if f.leanback {
      play.DefaultDevice.Feature = append(
         play.DefaultDevice.Feature, play.Leanback,
      )
   }
   data, err := play.Checkin{}.Marshal(&play.DefaultDevice)
   if err != nil {
      return err
   }
   return os.WriteFile(f.device_path(), data, os.ModePerm)
}

func (f *flags) do_sync() error {
   data, err := os.ReadFile(f.device_path())
   if err != nil {
      return err
   }
   var checkin play.Checkin
   err = checkin.Unmarshal(data)
   if err != nil {
      return err
   }
   if f.leanback {
      play.DefaultDevice.Feature = append(
         play.DefaultDevice.Feature, play.Leanback,
      )
   }
   return play.DefaultDevice.Sync(checkin)
}

func download(address, name string) error {
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   resp, err := http.Get(address)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   var progress xhttp.ProgressBytes
   progress.Set(resp)
   _, err = file.ReadFrom(&progress)
   if err != nil {
      return err
   }
   return nil
}

func (f *flags) do_delivery() error {
   var checkin play.Checkin
   auth, err := f.client(&checkin)
   if err != nil {
      return err
   }
   deliver, err := auth.Delivery(checkin, &f.app, f.single)
   if err != nil {
      return err
   }
   obbs := deliver.Obb()
   for {
      obb, ok := obbs()
      if !ok {
         break
      }
      err := download(
         obb.Url(), f.app.Obb(obb.Field1()),
      )
      if err != nil {
         return err
      }
   }
   apks := deliver.Apk()
   for {
      apk, ok := apks()
      if !ok {
         break
      }
      err := download(
         apk.Url(), f.app.Apk(apk.Field1()),
      )
      if err != nil {
         return err
      }
   }
   err = download(deliver.Url(), f.app.Apk(""))
   if err != nil {
      return err
   }
   return nil
}

func (f *flags) do_acquire() error {
   var checkin play.Checkin
   auth, err := f.client(&checkin)
   if err != nil {
      return err
   }
   return auth.Acquire(checkin, f.app.Id)
}

func (f *flags) device_path() string {
   var b strings.Builder
   b.WriteString(f.home)
   b.WriteByte('/')
   b.WriteString(play.DefaultDevice.Abi)
   if f.leanback {
      b.WriteString("-leanback")
   }
   b.WriteString(".txt")
   return b.String()
}

func (f *flags) client(checkin *play.Checkin) (*play.Auth, error) {
   data, err := os.ReadFile(f.home + "/token.txt")
   if err != nil {
      return nil, err
   }
   var token play.Token
   err = token.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   auth, err := token.Auth()
   if err != nil {
      return nil, err
   }
   data, err = os.ReadFile(f.device_path())
   if err != nil {
      return nil, err
   }
   err = checkin.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   return auth, nil
}
