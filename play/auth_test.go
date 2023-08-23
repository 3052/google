package play

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Auth_Read(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   raw, err := os.ReadFile(home + "/google/play/auth.txt")
   if err != nil {
      t.Fatal(err)
   }
   ref, err := New_Refresh_Token(raw)
   if err != nil {
      t.Fatal(err)
   }
   acc, err := ref.Refresh()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(acc)
}

const code = "oauth2_4/0Adeu5BVtOS_6vKpJRrBcF7xoa5V-J8XfKlMG3J1JbIj5bcaEb5IOX..."

func Test_Auth_Write(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   text, err := New_Raw_Token(code)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile(home + "/google/play/auth.txt", text, 0666)
}
