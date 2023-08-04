package play

import (
   "154.pages.dev/encoding/protobuf"
   "net/url"
   "strings"
)

// Checkin$AndroidCheckinResponse
type Device struct {
   m protobuf.Message
}

// godocs.io/encoding#BinaryMarshaler
func (d Device) MarshalBinary() ([]byte, error) {
   return d.m.Append(nil), nil
}

// godocs.io/encoding#BinaryUnmarshaler
func (d *Device) UnmarshalBinary(data []byte) error {
   var err error
   d.m, err = protobuf.Consume(data)
   if err != nil {
      return err
   }
   return nil
}

type Auth map[string]string

// godocs.io/flag#Value
func (a Auth) String() string {
   // cs.opensource.google/go/go/+/refs/tags/go1.20.7:src/net/url/url.go
   var buf strings.Builder
   for k, v := range a {
      if buf.Len() >= 1 {
         buf.WriteByte('\n')
      }
      buf.WriteString(url.QueryEscape(k))
      buf.WriteByte('=')
      buf.WriteString(url.QueryEscape(v))
   }
   return buf.String()
}

// godocs.io/flag#Value
func (a Auth) Set(query string) (err error) {
   // cs.opensource.google/go/go/+/refs/tags/go1.20.7:src/net/url/url.go
   for query != "" {
      var key string
      key, query, _ = strings.Cut(query, "\n")
      if key == "" {
         continue
      }
      key, value, _ := strings.Cut(key, "=")
      key, err1 := url.QueryUnescape(key)
      if err1 != nil {
         if err == nil {
            err = err1
         }
         continue
      }
      value, err1 = url.QueryUnescape(value)
      if err1 != nil {
         if err == nil {
            err = err1
         }
         continue
      }
      a[key] = value
   }
   return err
}

