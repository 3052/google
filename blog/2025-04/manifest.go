package manifest

import (
   "iter"
   "strings"
)

func (i *intent_filter) String() string {
   var b strings.Builder
   b.WriteString("action.name = ")
   b.WriteString(i.Action.Name)
   for _, category := range i.Category {
      b.WriteString("\ncategory.name = ")
      b.WriteString(category.Name)
   }
   for _, data := range i.Data {
      if data.Host != "" {
         b.WriteString("\ndata.host = ")
         b.WriteString(data.Host)
      }
      if data.PathPattern != "" {
         b.WriteString("\ndata.pathPattern = ")
         b.WriteString(data.PathPattern)
      }
      if data.PathPrefix != "" {
         b.WriteString("\ndata.pathPrefix = ")
         b.WriteString(data.PathPrefix)
      }
      if data.Scheme != "" {
         b.WriteString("\ndata.scheme = ")
         b.WriteString(data.Scheme)
      }
   }
   return b.String()
}

func (m manifest) intent_filter() iter.Seq[intent_filter] {
   return func(yield func(intent_filter) bool) {
      for _, activity := range m.Application.Activity {
         for _, intent := range activity.IntentFilter {
            if intent.Action.Name == "android.intent.action.VIEW" {
               if !yield(intent) {
                  return
               }
            }
         }
      }
   }
}

type manifest struct {
   Application struct {
      Activity []struct {
         IntentFilter []intent_filter `xml:"intent-filter"`
      } `xml:"activity"`
   } `xml:"application"`
}
 
type intent_filter struct {
   Action     struct {
      Name string `xml:"name,attr"`
   } `xml:"action"`
   Category []struct {
      Name string `xml:"name,attr"`
   } `xml:"category"`
   Data []struct {
      Scheme      string `xml:"scheme,attr"`
      Host        string `xml:"host,attr"`
      PathPattern string `xml:"pathPattern,attr"`
      PathPrefix  string `xml:"pathPrefix,attr"`
   } `xml:"data"`
}
