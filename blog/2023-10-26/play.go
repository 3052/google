package play

import "154.pages.dev/protobuf"

type Access_Token map[string]string

func (Access_Token) Auth() string {
   return ""
}

type Raw_Refresh_Token string

func (*Raw_Refresh_Token) Post(oauth_token string) error {
   return nil
}

func (Raw_Refresh_Token) Token() (Refresh_Token, error) {
   return nil, nil
}

type Refresh_Token map[string]string

func (Refresh_Token) Post() (Access_Token, error) {
   return nil, nil
}

func (Refresh_Token) Token() string {
   return ""
}

type Device struct{}

type Raw_Checkin []byte

func (*Raw_Checkin) Post(Device) error {
   return nil
}

func (Raw_Checkin) Checkin() (*Checkin, error) {
   return nil, nil
}

type Checkin struct {
   Message protobuf.Message
}

func (Checkin) Device_ID() (uint64, bool) {
   return 0, false
}

// 1. to
// 2. from
// 3. field/method

type Details struct {
   Message protobuf.Message
}

func (*Details) Get(doc string) error {
   return nil
}

func (Details) Downloads() (uint64, bool) {
   return 0, false
}
