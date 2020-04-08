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

// GetToken send request to amadeus api  aquire token
func (a *Amadeus) GetToken() error {

	// prepare request body
	body := url.Values{}
	body.Set("client_id", a.key)
	body.Set("client_secret", a.secret)
	body.Set("grant_type", "client_credentials")

	// get url string
	baseURL, err := a.getBaseURL()
	if err != nil {
		return err
	}

	// send request to api
	resp, err := http.Post(
		baseURL+"/"+securityOAuth2TokenURL,
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
	baseURL, err := a.getBaseURL()
	if err != nil {
		return err
	}

	// send request to api
	resp, err := http.Get(baseURL + "/" + securityOAuth2TokenURL + "/" + a.token.AccessToken)
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
