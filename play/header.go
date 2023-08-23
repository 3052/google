package play

import (
   "net/http"
   "strconv"
)

func (h Header) Set_Auth(head http.Header) {
   head.Set("Authorization", "Bearer " + h.Auth.auth())
}

func (h Header) Set_Device(head http.Header) error {
   id, err  := h.Device.ID()
   if err != nil {
      return err
   }
   head.Set("X-DFE-Device-ID", strconv.FormatUint(id, 16))
   return nil
}

type Header struct {
   Auth Access_Token // Authorization
   Device Device // X-DFE-Device-ID
   Single bool
}

func (h Header) Set_Agent(head http.Header) {
   var b []byte
   // `sdk` is needed for `/fdfe/delivery`
   b = append(b, "Android-Finsky (sdk="...)
   // valid range 0 - 0x7FFF_FFFF
   b = strconv.AppendInt(b, 9, 10)
   // com.android.vending
   b = append(b, ",versionCode="...)
   if h.Single {
      // valid range 8032_0000 - 8091_9999
      b = strconv.AppendInt(b, 8091_9999, 10)
   } else {
      // valid range 8092_0000 - 0x7FFF_FFFF
      b = strconv.AppendInt(b, 9999_9999, 10)
   }
   b = append(b, ')')
   head.Set("User-Agent", string(b))
}
