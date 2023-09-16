package main

import (
   "encoding/json"
   "fmt"
)

type fail struct {
   Auth_Token string `json:"authToken"`
}

func (v *fail) UnmarshalText(text []byte) error {
   return json.Unmarshal(text, v)
}

type pass struct {
   Auth_Token string `json:"authToken"`
}

func (v *pass) UnmarshalText(text []byte) error {
   type t pass
   var p t
   err := json.Unmarshal(text, &p)
   if err != nil {
      return err
   }
   *v = pass(p)
   return nil
}

func main() {
   text := []byte(`{"authToken": "hello world"}`)
   {
      var v fail
      err := v.UnmarshalText(text)
      // {} json: cannot unmarshal object into Go value of type *main.fail
      fmt.Println(v, err)
   }
   {
      var v pass
      err := v.UnmarshalText(text)
      // {hello world} <nil>
      fmt.Println(v, err)
   }
}
