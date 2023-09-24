package play

import (
   "154.pages.dev/strconv"
   "fmt"
)

func (d Details) String() string {
   var b []byte
   b = append(b, "creator: "...)
   {
      v, _ := d.Creator()
      b = append(b, v...)
   }
   b = append(b, "\nfile:"...)
   for _, file := range d.File() {
      if v, _ := file.File_Type(); v >= 1 {
         b = append(b, " OBB"...)
      } else {
         b = append(b, " APK"...)
      }
   }
   b = append(b, "\ninstallation size: "...)
   {
      v, _ := d.Installation_Size()
      b = fmt.Append(b, strconv.Size(v))
   }
   b = append(b, "\ndownloads: "...)
   {
      v, _ := d.Num_Downloads()
      b = fmt.Append(b, strconv.Cardinal(v))
   }
   b = append(b, "\noffer: "...)
   {
      v, _ := d.Micros()
      b = fmt.Append(b, v)
   }
   b = append(b, ' ')
   {
      v, _ := d.Currency_Code()
      b = append(b, v...)
   }
   b = append(b, "\ntitle: "...)
   {
      v, _ := d.Title()
      b = append(b, v...)
   }
   b = append(b, "\nupload date: "...)
   {
      v, _ := d.Upload_Date()
      b = append(b, v...)
   }
   b = append(b, "\nversion: "...)
   {
      v, _ := d.Version()
      b = append(b, v...)
   }
   b = append(b, "\nversion code: "...)
   {
      v, _ := d.Version_Code()
      b = fmt.Append(b, v)
   }
   return string(b)
}
