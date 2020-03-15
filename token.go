package amadeusgolang

import (
	"time"
)

type token struct {
	Type             string        `json:"type"`
	Username         string        `json:"username"`
	AppName          string        `json:"application_name"`
	ClientID         string        `json:"client_id"`
	TokenType        string        `json:"token_type"`
	AccessToken      string        `json:"access_token"`
	ExpiresIn        time.Duration `json:"expires_in"`
	State            string        `json:"state"`
	Scope            string        `json:"scope"`
	Created          time.Time
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Code             int    `json:"code"`
	Title            string `json:"title"`
}

func (t *token) getBearer() string {
	return t.TokenType + " " + t.AccessToken
}

func (t *token) expired() bool {

	if t.ExpiresIn == 0 {
		return true
	}

	if time.Now().Sub(t.Created) < t.ExpiresIn {
		return true
	}

	return false
}
