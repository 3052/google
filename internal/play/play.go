package main

import (
   "41.neocities.org/google/play"
   "41.neocities.org/text"
   "net/http"
   "os"
   "strings"
)

func (f *flags) do_checkin() error {
   if f.leanback {
      play.Device.Feature = append(play.Device.Feature, play.Leanback)
   }
   data, err := play.Device.Checkin(nil)
   if err != nil {
      return err
   }
   return os.WriteFile(f.device_path(), data, os.ModePerm)
}

func (f *flags) do_auth() error {
   data, err := (*play.GoogleToken).New(nil, f.token)
   if err != nil {
      return err
   }
   os.Mkdir(f.home, os.ModePerm)
   return os.WriteFile(f.home + "/token.txt", data, os.ModePerm)
}

func (f *flags) do_sync() error {
   data, err := os.ReadFile(f.device_path())
   if err != nil {
      return err
   }
   var checkin play.GoogleCheckin
   err = checkin.Unmarshal(data)
   if err != nil {
      return err
   }
   if f.leanback {
      play.Device.Feature = append(play.Device.Feature, play.Leanback)
   }
   return play.Device.Sync(&checkin)
}

func (f *flags) do_delivery() error {
   checkin := &play.GoogleCheckin{}
   auth, err := f.client(checkin)
   if err != nil {
      return err
   }
   deliver, err := auth.Delivery(checkin, &f.app, f.single)
   if err != nil {
      return err
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
   err = download(deliver.Url(), f.app.Apk(""))
   if err != nil {
      return err
   }
   return nil
}

func (f *flags) do_acquire() error {
   checkin := &play.GoogleCheckin{}
   auth, err := f.client(checkin)
   if err != nil {
      return err
   }
   return auth.Acquire(checkin, f.app.Id)
}

func (f *flags) device_path() string {
   var b strings.Builder
   b.WriteString(f.home)
   b.WriteByte('/')
   b.WriteString(play.Device.Abi)
   if f.leanback {
      b.WriteString("-leanback")
   }
   b.WriteString(".txt")
   return b.String()
}
func (f *flags) client(checkin *play.GoogleCheckin) (*play.GoogleAuth, error) {
   data, err := os.ReadFile(f.home + "/token.txt")
   if err != nil {
      return nil, err
   }
   var token play.GoogleToken
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

func download(address, name string) error {
   dst, err := os.Create(name)
   if err != nil {
      return err
   }
   defer dst.Close()
   resp, err := http.Get(address)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   var meter text.ProgressMeter
   meter.Set(1)
   _, err = dst.ReadFrom(meter.Reader(resp))
   if err != nil {
      return err
   }
   return nil
}

func (f *flags) do_details() (*play.Details, error) {
   checkin := &play.GoogleCheckin{}
   auth, err := f.client(checkin)
   if err != nil {
      return nil, err
   }
   return auth.Details(checkin, f.app.Id, f.single)
}
