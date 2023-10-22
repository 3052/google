package play

import (
	"154.pages.dev/protobuf"
	"bytes"
	"errors"
	"io"
	"net/http"
)

// this could be a method, but we want to match the signature of upload.
// device is Pixel 2
func New_Checkin(d Device) ([]byte, error) {
	var m protobuf.Message
	m.Add(4, func(m *protobuf.Message) {
		m.Add(1, func(m *protobuf.Message) {
			// single APK valid range 14 - 28
			// multiple APK valid range 14 - math.MaxInt32
			m.Add_Varint(10, 26)
		})
	})
	m.Add_Varint(14, 3)
	m.Add(18, func(m *protobuf.Message) {
		m.Add_Varint(1, 3)
		m.Add_Varint(2, 2)
		m.Add_Varint(3, 1)
		m.Add_Varint(4, 2)
		m.Add_Varint(5, 1)
		m.Add_Varint(6, 0)
		m.Add_Varint(7, 420)
		m.Add_Varint(8, 196609)
		for _, library := range d.Library {
			m.Add_String(9, library)
		}
		m.Add_String(11, d.Platform)
		for _, texture := range d.Texture {
			m.Add_String(15, texture)
		}
		// you cannot swap the next two lines:
		for _, feature := range d.Feature {
			m.Add(26, func(m *protobuf.Message) {
				m.Add_String(1, feature)
			})
		}
	})
	res, err := http.Post(
		"https://android.clients.google.com/checkin",
		"application/x-protobuffer",
		bytes.NewReader(m.Append(nil)),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	return io.ReadAll(res.Body)
}

type Checkin struct {
	m protobuf.Message
}

func (c Checkin) device_ID() (uint64, error) {
	if f, ok := c.m.Fixed64(7); ok {
		return f, nil
	}
	return 0, errors.New("device ID")
}
