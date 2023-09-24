package header

import (
	"fmt"
	"testing"
)

func Test_Cookie(t *testing.T) {
	m, err := cookie_two(cookie_two_in)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", m)
}
