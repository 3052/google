package play

type Access_Token struct{}

func (Access_Token) Authorization() (string, string) {
   return "", ""
}

type Acquire_Request struct {
   Access_Token
   Checkin
}

type Checkin struct{}

func (Acquire_Request) Do(app string) error {
   return nil
}

type Application struct {
   ID string
   Version uint64
}

func (Application) APK(config string) string {
   return ""
}

func (Application) OBB(role uint64) string {
   return ""
}

type Delivery struct{}

type Delivery_Request struct {
   Access_Token
   Checkin
   Application
}

func (Checkin) X_DFE_Device_ID() (string, string) {
   return "", ""
}

func User_Agent(single bool) (string, string) {
   return "", ""
}

func (Delivery_Request) Do(single bool) (*Delivery, error) {
   return nil, nil
}

type Details struct{}

type Details_Request struct {
   Access_Token
   Checkin
}

func (Details_Request) Do(app string, single bool) (*Details, error) {
   return nil, nil
}

type Device struct{}

////////////////////////////////////////////////////////////

func (Device) Checkin() (*Raw_Checkin, error) {
   return nil, nil
}

type Raw_Checkin struct{}

func (Raw_Checkin) Checkin() (*Checkin, error) {
   return nil, nil
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

type Sync_Request struct {
   Device
   Checkin
}

func (Sync_Request) Do() error {
   return nil
}
