package play

import (
   "fmt"
   "os"
   "testing"
)

func Test_Auth_Write(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   text, err := New_Refresh_Token(code)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile(home + "/google/play/token.txt", text, 0666)
}

func Test_Auth_Read(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   token, err := os.ReadFile(home + "/google/play/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   device, err := os.ReadFile(home + "/google/play/x86.bin")
   if err != nil {
      t.Fatal(err)
   }
   head, err := New_Header(token, device, false)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(head)
}

const code = "oauth2_4/0Adeu5BVtOS_6vKpJRrBcF7xoa5V-J8XfKlMG3J1JbIj5bcaEb5IOX..."
