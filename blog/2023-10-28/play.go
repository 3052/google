package play

import "154.pages.dev/protobuf"

type Access_Token map[string]string

type Acquire struct {
   Access_Token
   Checkin
}

func (Acquire) Acquire(app string) error {
   return nil
}

type Application struct{}

type Checkin struct{}

func (Checkin) Sync(Device) error {
   return nil
}

type Delivery struct {
   Access_Token
   Application
   Checkin
   protobuf.Message
}

func (*Delivery) Delivery(single bool) error {
   return nil
}

//////////////////////////////////////////

type Details struct {
   Access_Token
   Checkin
   protobuf.Message
}

func (*Details) Details(app string, single bool) error {
   return nil
}

type Device struct{}

func (Device) Checkin() (Raw_Checkin, error) {
   return nil, nil
}

type Raw_Checkin []byte

func (Raw_Checkin) Checkin() (*Checkin, error) {
   return nil, nil
}

type Raw_Refresh_Token []byte

func Auth(oauth_token string) (Raw_Refresh_Token, error) {
   return nil, nil
}

func (Raw_Refresh_Token) Refresh_Token() (Refresh_Token, error) {
   return nil, nil
}

type Refresh_Token map[string]string

func (Refresh_Token) Auth() (Access_Token, error) {
   return nil, nil
}
