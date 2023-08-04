package play

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "net/url"
   "strings"
)

type Auth map[string]string

// cs.opensource.google/go/go/+/refs/tags/go1.20.7:src/net/url/url.go
func (a Auth) MarshalText() ([]byte, error) {
   var b bytes.Buffer
   for k, v := range a {
      if b.Len() >= 1 {
         b.WriteByte('\n')
      }
      b.WriteString(url.QueryEscape(k))
      b.WriteByte('=')
      b.WriteString(url.QueryEscape(v))
   }
   return b.Bytes(), nil
}

// cs.opensource.google/go/go/+/refs/tags/go1.20.7:src/net/url/url.go
func (a Auth) UnmarshalText(text []byte) error {
   query := string(text)
   for query != "" {
      var line string
      line, query, _ = strings.Cut(query, "\n")
      key, value, _ := strings.Cut(line, "=")
      var err error
      key, err = url.QueryUnescape(key)
      if err != nil {
         return err
      }
      value, err = url.QueryUnescape(value)
      if err != nil {
         return err
      }
      a[key] = value
   }
   return nil
}

// Checkin$AndroidCheckinResponse
type Device struct {
   m protobuf.Message
}

func (d Device) MarshalBinary() ([]byte, error) {
   return d.m.Append(nil), nil
}

func (d *Device) UnmarshalBinary(data []byte) error {
   var err error
   d.m, err = protobuf.Consume(data)
   if err != nil {
      return err
   }
   return nil
}
