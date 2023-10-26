package play

type Refresh_Token map[string]string

func (Refresh_Token) Parse(string) {}

func (Refresh_Token) Encode(string) string {
   return ""
}
