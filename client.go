package amadeus

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

	//
	// Base URLs
	//

	// Testing API url
	test = "https://test.api.amadeus.com"

	// Production API url
	production = "https://api.amadeus.com"

	//
	// Requests URLs
	//

	// Authentification url
	// use to aquire token which is used in all other request
	auth = "/v1/security/oauth2/token"

	///////////
	// 	AIR	 //
	///////////

	//
	// Shooping requests
	//

	// Flight Offers Search
	// search for offers on given origin, destination, departure, passangers
	shoopingFlightOffers = "/v2/shopping/flight-offers"

	// Flight Inspiration Search
	// check certain offer if is still active, response with additional data for offer
	shoopingFlightDestinations = "/v1/shopping/flight-destinations"

	// Shooping Flight offers pricing
	// check certain offer if is still active, response with additional data for offer
	shoopingFlightOffersPricing = "/v1/shopping/flight-offers/pricing"

	//
	// Booking requests
	//

	// Booking Flight orders
	// create reservation for certain offer
	bookingFlightOrders = "/v1/booking/flight-orders"
)

// Amadeus main struct that holds sensitive data for communicating with api
// key, secret and env for requesting token for authentification
// which is used in all other requests
type Amadeus struct {
	key    string
	secret string
	env    string
	token  token
}

// New creates new amadeus client for given Key, Secret & Environment
// Key & Secret are created on amadeus developers page
// https://developers.amadeus.com/register
func New(Key, Secret, ENV string) (*Amadeus, error) {

	var (
		a   Amadeus
		err error
	)

	err = a.setKey(Key)
	if err != nil {
		return nil, err
	}

	err = a.setSecret(Secret)
	if err != nil {
		return nil, err
	}

	err = a.setENV(ENV)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

// setKey for field key in Amadeus struct
// check if empty than return error
func (a *Amadeus) setKey(value string) error {

	if value == "" {
		return errors.New("key is empty")
	}
	a.key = value
	return nil
}

// setSecret for field secret in Amadeus struct
// check if empty than return error
func (a *Amadeus) setSecret(value string) error {

	if value == "" {
		return errors.New("secret is empty")
	}
	a.secret = value
	return nil
}

// setENV for env field in Amadeus struct
// check if empty than return error
// check if valid environment for using base const url
func (a *Amadeus) setENV(value string) error {

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

// getURL return full url for given endpoint
// checks for environment base url and add endpoint url
func (a Amadeus) getURL(endpoint string) (string, error) {

	switch a.env {
	case "TEST":
		return test + endpoint, nil
	case "PRODUCTION":
		return production + endpoint, nil
	}

	return "", errors.New("not defined valid environment")
}

// GetToken send request to amadeus api  aquire token
func (a *Amadeus) GetToken() error {

	// prepare request body
	body := url.Values{}
	body.Set("client_id", a.key)
	body.Set("client_secret", a.secret)
	body.Set("grant_type", "client_credentials")

	// get url string
	urlStr, err := a.getURL(auth)
	if err != nil {
		return err
	}

	// send request to api
	resp, err := http.Post(
		urlStr,
		"application/x-www-form-urlencoded",
		strings.NewReader(body.Encode()),
	)
	if err != nil {
		return err
	}

	// fetch token when time expires
	now := time.Now()

	defer resp.Body.Close()

	// check if status code valid
	if resp.StatusCode != 200 {
		return errors.New("requesting token failed")
	}

	// read body response
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// convert json struct to Token struct
	err = json.Unmarshal(r, &a.token)
	if err != nil {
		return err
	}

	// add created time
	a.token.Created = now

	// check expiry on retrived token
	if a.token.ExpiresIn == 0 {
		return errors.New("returned token not valid")
	}

	return nil
}

// CheckToken send request to amadeus api to check if token is still valid
func (a *Amadeus) CheckToken() error {

	// get url string
	urlStr, err := a.getURL(auth)
	if err != nil {
		return err
	}

	// send request to api
	resp, err := http.Get(urlStr + "/" + a.token.getAuthorization())
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// check if status code valid
	if resp.StatusCode != 200 {
		return errors.New("requesting token failed")
	}

	// read body response
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// convert json struct to Token struct
	err = json.Unmarshal(r, &a.token)
	if err != nil {
		return err
	}

	// check expiry on retrived token
	if a.token.ExpiresIn == 0 {
		return errors.New("returned token not valid")
	}

	return nil
}

// ErrorResponse struct for error response from api
type ErrorResponse struct {
	Code   int    `json:"code,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
	Source struct {
		Pointer string `json:"pointer,omitempty"`
		Example string `json:"example,omitempty"`
	} `json:"source,omitempty"`
	Status int `json:"status,omitempty"`
}

// requests send POST request to api with given payload
func (a *Amadeus) getRequest(url string, queryParams []string) ([]byte, error) {

	if a.token.expired() {
		err := a.GetToken()
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", a.token.getAuthorization())
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// requests send POST request to api with given payload
func (a *Amadeus) postRequest(reqPayload, url string) ([]byte, error) {

	if a.token.expired() {
		err := a.GetToken()
		if err != nil {
			return nil, err
		}
	}

	payload := strings.NewReader(reqPayload)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", a.token.getAuthorization())
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

type token struct {
	Type             string        `json:"type,omitempty"`
	Username         string        `json:"username,omitempty"`
	AppName          string        `json:"application_name,omitempty"`
	ClientID         string        `json:"client_id,omitempty"`
	TokenType        string        `json:"token_type,omitempty"`
	AccessToken      string        `json:"access_token,omitempty"`
	ExpiresIn        time.Duration `json:"expires_in,omitempty"`
	State            string        `json:"state,omitempty"`
	Scope            string        `json:"scope,omitempty"`
	Error            string        `json:"error,omitempty"`
	ErrorDescription string        `json:"error_description,omitempty"`
	Code             int           `json:"code,omitempty"`
	Title            string        `json:"title,omitempty"`
	Created          time.Time
}

// getAuthorization return token type and token
func (t *token) getAuthorization() string {
	return t.TokenType + " " + t.AccessToken
}

// expired check if token is expired
func (t *token) expired() bool {

	if t.ExpiresIn == 0 {
		return true
	}

	if time.Now().Sub(t.Created) < t.ExpiresIn {
		return true
	}

	return false
}
