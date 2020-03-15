package amadeusgolang

import (
	"encoding/json"
	"errors"
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

func New(Key, Secret, ENV string) (amadeus, error) {

	var (
		a   amadeus
		err error
	)

	err = a.setKey(Key)
	if err != nil {
		return a, err
	}

	err = a.setSecret(Secret)
	if err != nil {
		return a, err
	}

	err = a.setENV(ENV)
	if err != nil {
		return a, err
	}

	return a, nil
}

type amadeus struct {
	key    string
	secret string
	env    string
	token  token
}

func (a *amadeus) setKey(value string) error {

	if value == "" {
		return errors.New("key is empty")
	}
	a.key = value
	return nil
}

func (a *amadeus) setSecret(value string) error {

	if value == "" {
		return errors.New("secret is empty")
	}
	a.secret = value
	return nil
}

func (a *amadeus) setENV(value string) error {

	if value == "" {
		return errors.New("env is empty")
	}

	switch value {
	case "TEST":
		a.env = value
		return nil
	case "PRODUCTION":
		a.env = value
		return nil
	default:
		return errors.New("env not set")
	}

}

func (a *amadeus) getURL(endpoint string) string {

	switch a.env {
	case "TEST":
		return test + endpoint
	case "PRODUCTION":
		return production + endpoint
	}

	return ""
}

func (a *amadeus) GetToken() error {

	err := a.requestToken()
	if err != nil {
		return err
	}

	if a.token.ExpiresIn == 0 {
		return errors.New("token not valid")
	}

	return nil
}

func (a *amadeus) requestToken() error {

	// this is the way to send body of mime-type: x-www-form-urlencoded
	body := url.Values{}
	body.Set("client_id", a.key)
	body.Set("client_secret", a.secret)
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

	/*
		if resp.StatusCode != 200 {
			return err
		}
	*/

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(r, &a.token)
	if err != nil {
		return err
	}

	a.token.Created = now
	return nil
}
