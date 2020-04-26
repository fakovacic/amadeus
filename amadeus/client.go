package amadeus

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fakovacic/amadeus-golang/amadeus/client"
	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/security"
)

const (

	//
	// Base URLs
	//

	// Testing API url
	test = "https://test.api.amadeus.com"

	// Production API url
	production = "https://api.amadeus.com"
)

var a *Amadeus

// Amadeus main struct that holds sensitive data for communicating with api
// key, secret and env for requesting token for authentification which is used in all other requests
type Amadeus struct {
	key    string
	secret string
	env    string
	Token  *security.TokenResponse
	Airport
	Media
	ReferenceData
	Travel
	EReputation
	Shooping
}

// New creates new amadeus client for given Key, Secret & Environment
// Key & Secret are created on amadeus developers page https://developers.amadeus.com/register
func New(Key, Secret string) (*Amadeus, error) {

	var err error

	if a == nil {
		a = new(Amadeus)
	}

	err = a.setKey(Key)
	if err != nil {
		return nil, err
	}

	err = a.setSecret(Secret)
	if err != nil {
		return nil, err
	}

	err = a.SetENV("test")
	if err != nil {
		return nil, err
	}

	return a, nil
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

// SetENV for env field in Amadeus struct
// check if empty than return error
// check if valid environment for using base const url
func (a *Amadeus) SetENV(value string) error {

	if value == "" {
		return errors.New("env is empty")
	}

	switch value {
	case "test":
		a.env = value
		return nil
	case "production":
		a.env = value
		return nil
	default:
		return errors.New("env not set")
	}

}

func (a Amadeus) GetAuthorization() string {
	return a.Token.GetAuthorization()
}

// getURL return full url for given endpoint
// checks for environment base url and add endpoint url
func (a Amadeus) GetBaseURL() (string, error) {

	switch a.env {
	case "test":
		return test, nil
	case "production":
		return production, nil
	}

	return "", errors.New("not defined valid environment")
}

func (amadeus *Amadeus) GetToken() error {

	req, res, err := client.NewRequest(client.SecurityToken)

	req.(*security.TokenRequest).SetKey(amadeus.key)
	req.(*security.TokenRequest).SetSecret(amadeus.secret)

	// get base api url
	baseURL, err := amadeus.GetBaseURL()
	if err != nil {
		return err
	}

	// prepare request
	r, err := http.NewRequest(
		"POST",
		baseURL+req.GetURL("POST"),
		req.GetBody("POST"),
	)
	if err != nil {
		return err
	}

	// add headers
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Accept", "application/json")

	// send request
	client := http.Client{}
	rsp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	// read body
	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	// check status code
	// return error if not 200

	fmt.Println("------------------")
	fmt.Println(rsp.StatusCode)
	fmt.Println(req.GetURL("POST"))
	fmt.Println(string(b))
	fmt.Println("------------------")

	// decode response to struct
	err = res.Decode(b)

	if err != nil {
		return err
	}

	amadeus.Token = res.(*security.TokenResponse)

	return nil

}

func (amadeus *Amadeus) CheckToken() error {

	if amadeus.Token == nil {
		return errors.New("token empty")
	}

	if amadeus.Token.AccessToken == "" {
		return errors.New("access token empty")
	}

	req, res, err := client.NewRequest(client.SecurityToken)

	req.(*security.TokenRequest).SetAccessToken(amadeus.Token.AccessToken)

	// get base api url
	baseURL, err := amadeus.GetBaseURL()
	if err != nil {
		return err
	}

	// prepare request
	r, err := http.NewRequest(
		"GET",
		baseURL+req.GetURL("GET"),
		req.GetBody("GET"),
	)
	if err != nil {
		return err
	}

	// add headers
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Accept", "application/json")

	// send request
	client := http.Client{}
	rsp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	// read body
	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	// check status code
	// return error if not 200

	fmt.Println("------------------")
	fmt.Println(rsp.StatusCode)
	fmt.Println(req.GetURL("GET"))
	fmt.Println(string(b))
	fmt.Println("------------------")

	// decode response to struct
	err = res.Decode(b)

	if err != nil {
		return err
	}

	amadeus.Token = res.(*security.TokenResponse)

	return nil

}
