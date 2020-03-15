package amadeusgolang

import (
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
