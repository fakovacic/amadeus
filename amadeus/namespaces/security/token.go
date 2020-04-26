package security

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
	"time"
)

type TokenRequest struct {
	AccessToken string `json:"access_token,omitempty"`
	Key         string `json:"client_id,omitempty"`
	Secret      string `json:"client_secret,omitempty"`
}

// SetKey set access token
func (dR *TokenRequest) SetKey(key string) *TokenRequest {

	dR.Key = key

	return dR
}

// SetSecret set access token
func (dR *TokenRequest) SetSecret(secret string) *TokenRequest {

	dR.Secret = secret

	return dR
}

// SetAccessToken set access token
func (dR *TokenRequest) SetAccessToken(accessToken string) *TokenRequest {

	dR.AccessToken = accessToken

	return dR
}

// SetParam set params
func (dR *TokenRequest) SetParam(key, value string) {
	return
}

// GetURL returned key=value format for request on api
func (dR TokenRequest) GetURL(reqType string) string {

	// set request url
	url := oAuth2TokenURL

	// add version
	switch reqType {
	case "GET":
		return "/v1" + url + "/" + dR.AccessToken
	case "POST":
		return "/v1" + url
	}

	return ""
}

// GetBody implementation for Request
func (dR TokenRequest) GetBody(reqType string) io.Reader {

	switch reqType {
	case "POST":
		body := url.Values{}
		body.Set("client_id", dR.Key)
		body.Set("client_secret", dR.Secret)
		body.Set("grant_type", "client_credentials")

		return strings.NewReader(body.Encode())
	}

	return nil
}

type TokenResponse struct {
	Type        string        `json:"type,omitempty"`
	Username    string        `json:"username,omitempty"`
	AppName     string        `json:"application_name,omitempty"`
	ClientID    string        `json:"client_id,omitempty"`
	TokenType   string        `json:"token_type,omitempty"`
	AccessToken string        `json:"access_token,omitempty"`
	ExpiresIn   time.Duration `json:"expires_in,omitempty"`
	State       string        `json:"state,omitempty"`
	Scope       string        `json:"scope,omitempty"`
}

// Decode implement Response interface
func (dR *TokenResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

// GetAuthorization return Authorization string
func (dR TokenResponse) GetAuthorization() string {
	return dR.TokenType + " " + dR.AccessToken
}
