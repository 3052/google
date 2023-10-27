package play

type Access_Token struct{}

type Acquire_Header struct{}

func (Acquire_Header) Do(doc string) error {
   return nil
}

type App struct {
   Doc string
   VC uint64
}

func (App) APK(config string) string {
   return ""
}

func (App) OBB(role uint64) string {
   return ""
}

type Checkin struct{}

type Delivery struct{}

type Delivery_Header struct{}

func (Delivery_Header) Do(App) (*Delivery, error) {
   return nil, nil
}

type Details struct{}

type Details_Header struct{}

func (Details_Header) Do(doc string) (*Details, error) {
   return nil, nil
}

type Device struct{}

type Raw_Checkin struct{}

func (Raw_Checkin) Checkin() (*Checkin, error) {
   return nil, nil
}

func (*Raw_Checkin) Do(Device) error {
   return nil
}

type Raw_Refresh_Token struct{}

func (*Raw_Refresh_Token) Do(oauth_token string) error {
   return nil
}

func (Raw_Refresh_Token) Token() (*Refresh_Token, error) {
   return nil, nil
}

type Refresh_Token struct{}

func (Refresh_Token) Do() (*Access_Token, error) {
   return nil, nil
}

type Sync_Header struct{}

func (Sync_Header) Do(Device) error {
   return nil
}
