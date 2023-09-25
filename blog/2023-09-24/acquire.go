package acquire

import "154.pages.dev/protobuf"

type acquire_error struct {
   m protobuf.Message
}

func (a acquire_error) Error() string {
   var b []byte
   m, _ := a.m.Message(1)
   m, _ = m.Message(94)
   m, _ = m.Message(1)
   m, _ = m.Message(2)
   m, _ = m.Message(175_996_169)
   for _, f := range m {
      if f.Number == 2 {
         if m, ok := f.Message(); ok {
            m, _ = m.Message(20)
            m, _ = m.Message(1)
            c, _ := m.Bytes(1)
            if b != nil {
               b = append(b, '\n')
            }
            b = append(b, c...)
         }
      }
   }
   return string(b)
}
