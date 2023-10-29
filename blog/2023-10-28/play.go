package play

import (
   "154.pages.dev/protobuf"
   "net/url"
)

var Phone = Device{}

type Access_Token struct{}

func (*Access_Token) Refresh(Refresh_Token) error {
   return nil
}

type Acquire struct {
   Access_Token
   Checkin
}

func (Acquire) Acquire(app string) error {
   return nil
}

type Application struct{}

type Checkin struct {
   Raw []byte
   protobuf.Message
}

func (*Checkin) Unmarshal() error {
   return nil
}

type Device struct{}

type Delivery struct {
   Access_Token
   Application
   Checkin
}

type Details struct {
   Access_Token
   Checkin
}

func (Device) Checkin() (*Checkin, error) {
   return nil, nil
}

type Refresh_Token struct {
   Raw []byte
   url.Values
}

func (*Refresh_Token) Unmarshal() error {
   return nil
}

func Exchange(oauth_token string) (*Refresh_Token, error) {
   return nil, nil
}

func (Device) Sync(*Checkin) error {
   return nil
}

func (*Details) Details(app string, single bool) error {
   return nil
}

func (*Delivery) Delivery(single bool) error {
   return nil
}
