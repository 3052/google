package play

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_Details(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	home += "/google/play/"
	var head Header
	b, err := os.ReadFile(home + "token.txt")
	if err != nil {
		t.Fatal(err)
	}
	head.Set_Authorization(b)
	head.Set_Agent(false)
	for i, app := range apps {
		b, err := os.ReadFile(home + Platforms[app.platform] + ".bin")
		if err != nil {
			t.Fatal(err)
		}
		head.Set_Device(b)
		d, err := head.Details(app.doc)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := d.Downloads(); !ok {
			t.Fatal(err)
		}
		if f := d.Files(); len(f) == 0 {
			t.Fatal(f)
		}
		if _, ok := d.Name(); !ok {
			t.Fatal("name")
		}
		if _, ok := d.Offered_By(); !ok {
			t.Fatal("offered by")
		}
		if _, ok := d.Price(); !ok {
			t.Fatal("price")
		}
		if _, ok := d.Price_Currency(); !ok {
			t.Fatal("price currency")
		}
		if _, ok := d.Requires(); !ok {
			t.Fatal("requires")
		}
		if _, ok := d.Size(); !ok {
			t.Fatal("size")
		}
		{
			s, ok := d.Updated_On()
			if !ok {
				t.Fatal("updated on")
			}
			d, err := time.Parse("Jan 2, 2006", s)
			if err != nil {
				t.Fatal(err)
			}
			apps[i].date = d.Format("2006-01-02")
		}
		if _, ok := d.Version_Code(); !ok {
			t.Fatal("version code")
		}
		if _, ok := d.Version_Name(); !ok {
			t.Fatal("version name")
		}
		fmt.Printf("%#v,\n", app)
		time.Sleep(99 * time.Millisecond)
	}
}
