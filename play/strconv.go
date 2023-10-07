package play

import (
   "154.pages.dev/strconv"
   "fmt"
)

func (d Details) String() string {
   var b []byte
   b = append(b, "downloads:"...)
   if v, ok := d.Downloads(); ok {
      b = fmt.Append(b, " ", strconv.Cardinal(v))
   }
   b = append(b, "\nfiles:"...)
   for _, file := range d.Files() {
      if file >= 1 {
         b = append(b, " OBB"...)
      } else {
         b = append(b, " APK"...)
      }
   }
   b = append(b, "\nname:"...)
   if v, ok := d.Name(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\noffered by:"...)
   if v, ok := d.Offered_By(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\nprice:"...)
   if v, ok := d.Price(); ok {
      b = fmt.Append(b, " ", v)
   }
   if v, ok := d.Price_Currency(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\nrequires:"...)
   if v, ok := d.Requires(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\nsize:"...)
   if v, ok := d.Size(); ok {
      b = fmt.Append(b, " ", strconv.Size(v))
   }
   b = append(b, "\nupdated on:"...)
   if v, ok := d.Updated_On(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\nversion code:"...)
   if v, ok := d.Version_Code(); ok {
      b = fmt.Append(b, " ", v)
   }
   b = append(b, "\nversion name:"...)
   if v, ok := d.Version_Name(); ok {
      b = fmt.Append(b, " ", v)
   }
   return string(b)
}
