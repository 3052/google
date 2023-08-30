package acquire

import (
	"154.pages.dev/encoding/protobuf"
	"154.pages.dev/google/play"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const device_ID = "306e9f7f4192be79"

func New_Delivery(h *play.Header, doc string, vc uint64) error {
	req, err := http.NewRequest(
		"GET", "https://play-fe.googleapis.com/fdfe/delivery", nil,
	)
	if err != nil {
		return err
	}
	req.URL.RawQuery = url.Values{
		"doc": {doc},
		"vc":  {strconv.FormatUint(vc, 10)},
	}.Encode()
	req.Header.Set(h.Agent())
	req.Header.Set(h.Authorization())
	req.Header.Set("X-DFE-Device-Id", device_ID)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// ResponseWrapper
	mes, err := protobuf.Consume(data)
	if err != nil {
		return err
	}
	// payload
	mes, _ = mes.Message(1)
	// deliveryResponse
	mes, _ = mes.Message(21)
	status, err := mes.Varint(1)
	if err != nil {
		return err
	}
	switch status {
	case 3:
		return errors.New("purchase required")
	case 5:
		return errors.New("invalid version")
	}
	if _, err := mes.Message(2); err != nil {
		return errors.New("appDeliveryData not found")
	}
	return nil
}
