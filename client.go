package amadeusgo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	test       = "https://test.api.amadeus.com"
	production = "https://api.amadeus.com"
)

const (
	auth                        = "/v1/security/oauth2/token"
	shoopingFlightOffers        = "/v2/shopping/flight-offers"
	shoopingFlightOffersPricing = "/v1/shopping/flight-offers/pricing"
	bookingFlightOrders         = "/v1/booking/flight-orders"
)

type Amadeus struct {
	Key    string
	Secret string
	ENV    string
	Token  token
}

func (a *Amadeus) New(Key, Secret, ENV string) error {

	a.Key = Key
	a.Secret = Secret

	// check env
	// set default
	// generate token, return error about credentials
	err := a.getToken()
	if err != nil {
		return err
	}

	return nil
}

func (a *Amadeus) getURL(endpoint string) string {

	switch a.ENV {
	case "TEST":
		return test + endpoint
	case "PRODUCTION":
		return production + endpoint
	}

	return ""
}

func (a *Amadeus) getToken() error {

	// this is the way to send body of mime-type: x-www-form-urlencoded
	body := url.Values{}
	body.Set("client_id", a.Key)
	body.Set("client_secret", a.Secret)
	body.Set("grant_type", "client_credentials")

	contentType := "application/x-www-form-urlencoded"
	urlStr := a.getURL(auth)
	resp, err := http.Post(urlStr, contentType, strings.NewReader(body.Encode()))

	// fetch token when time expires
	now := time.Now()

	defer resp.Body.Close()
	if err != nil {
		return err
	}

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(r, &a.Token)
	if err != nil {
		return err
	}

	a.Token.Created = now
	return nil
}
