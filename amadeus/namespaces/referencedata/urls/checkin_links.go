package urls

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/structs"
)

const (

	// Flight Check-in Links
	checkinLinksURL = "/reference-data/urls/checkin-links"
)

type CheckInLinksRequest struct {
	AirlineCode string `json:"airlineCode"`
	Language    string `json:"language"`
}

// SetAirlineCode add airlines code
func (dR *CheckInLinksRequest) SetAirlineCode(airlineCode string) *CheckInLinksRequest {

	dR.AirlineCode = airlineCode

	return dR
}

// SetLanguage set language
func (dR *CheckInLinksRequest) SetLanguage(language string) *CheckInLinksRequest {

	dR.Language = language

	return dR
}

// ParseParams parse params
func (dR *CheckInLinksRequest) ParseParams(params []string) *CheckInLinksRequest {

	if len(params) == 0 {
		return dR
	}

	for _, param := range params {
		p := strings.Split(param, "=")

		if len(p) != 2 {
			continue
		}

		dR.SetParam(p[0], p[1])

	}

	return dR
}

// SetParam set param
func (dR *CheckInLinksRequest) SetParam(key, value string) {

	switch key {
	case "airlineCode":
		dR.SetAirlineCode(value)
	case "language":
		dR.SetLanguage(value)
		break
	}
}

// GetURL returned key=value format for request on api
func (dR CheckInLinksRequest) GetURL(reqType string) string {

	// set request url
	url := checkinLinksURL

	// add version
	switch reqType {
	case "GET":

		url = "/v2" + url

		// define query params
		queryParams := []string{}

		queryParams = append(queryParams, "airlineCode="+dR.AirlineCode)

		if dR.Language != "" {
			queryParams = append(queryParams, "language="+dR.Language)
		}

		url = url + "?" + strings.Join(queryParams, "&")

		break
	}

	return url
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (dR CheckInLinksRequest) GetBody(reqType string) io.Reader {
	return nil
}

type CheckInLinksResponse struct {
	Meta     structs.Meta       `json:"meta,omitempty"`
	Data     []UrlData          `json:"data,omitempty"`
	Warnings []structs.Warnings `json:"warnings,omitempty"`
}

// Decode implement Response interface
func (dR *CheckInLinksResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type UrlData struct {
	Type       string     `json:"type,omitempty"`
	ID         string     `json:"id,omitempty"`
	Href       string     `json:"href,omitempty"`
	Channel    string     `json:"channel,omitempty"`
	Parameters Parameters `json:"parameters,omitempty"`
}

type Parameters struct {
	LAST LAST `json:"LAST,omitempty"`
	PNR  PNR  `json:"PNR,omitempty"`
}

type LAST struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
}

type PNR struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Format      string `json:"format,omitempty"`
}
