package play

import (
   "fmt"
   "os"
   "testing"
)

func Test_Auth_Read(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   text, err := os.ReadFile(home + "/google/play/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   ref, err := Raw_Token.Refresh(text)
   if err != nil {
      t.Fatal(err)
   }
   acc, err := ref.Access()
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
   os.WriteFile(home + "/google/play/token.txt", text, 0666)
}
