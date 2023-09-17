package gplayapi

import (
   "errors"
   "net/http"
   "net/url"
   "strconv"
)

import "154.pages.dev/google/gplayapi/gpproto"

var ErrMissingAppDeliveryData = errors.New("buy response is missing AppDeliveryData")

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
