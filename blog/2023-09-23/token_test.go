// X-DFE-Device-Config-Token
package header

import (
	"fmt"
	"testing"
)

func Test_Token(t *testing.T) {
	m, err := token_one(token_one_in)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", m)
}
