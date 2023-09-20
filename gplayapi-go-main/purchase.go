package gplayapi

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Juby210/gplayapi-go/gpproto"
)

var ErrMissingAppDeliveryData = errors.New("buy response is missing AppDeliveryData")

func (client *GooglePlayClient) GetBuyResponse(packageName string, version int) (*gpproto.BuyResponse, error) {
	params := &url.Values{}
	params.Set("ot", "1")
	params.Set("doc", packageName)
	params.Set("vc", strconv.Itoa(version))

	r, _ := http.NewRequest("POST", UrlPurchase+"?"+params.Encode(), nil)
	payload, err := client.doAuthedReq(r)
	if err != nil {
		return nil, err
	}
	if payload == nil {
		return nil, ErrNilPayload
	}
	return payload.BuyResponse, nil
}

func (client *GooglePlayClient) GetDeliveryResponse(packageName string, version int) (*gpproto.DeliveryResponse, error) {
	params := &url.Values{}
	params.Set("ot", "1")
	params.Set("doc", packageName)
	params.Set("vc", strconv.Itoa(version))

	r, _ := http.NewRequest("GET", UrlDelivery+"?"+params.Encode(), nil)
	payload, err := client.doAuthedReq(r)
	if err != nil {
		return nil, err
	}
	if payload == nil {
		return nil, ErrNilPayload
	}
	return payload.DeliveryResponse, nil
}

func (client *GooglePlayClient) Purchase(packageName string, version int) (*gpproto.AndroidAppDeliveryData, error) {
	if version == 0 {
		app, err := client.GetAppDetails(packageName)
		if err != nil {
			return nil, err
		}
		version = app.VersionCode
	}

	_, _ = client.GetBuyResponse(packageName, version)

	res, err := client.GetDeliveryResponse(packageName, version)
	if err != nil {
		return nil, err
	}

	appDeliveryData := res.AppDeliveryData
	if appDeliveryData == nil {
		return nil, ErrMissingAppDeliveryData
	}
	return appDeliveryData, nil
}
