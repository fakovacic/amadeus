package amadeus

import (
	"encoding/json"
	"io"
	"strings"
)

type ReferenceDataUrlsCheckinLinksRequest struct {
	AirlineCode string `json:"airlineCode"`
	Language    string `json:"language"`
}

// SetAirlineCode add airlines code
func (dR *ReferenceDataUrlsCheckinLinksRequest) SetAirlineCode(airlineCode string) *ReferenceDataUrlsCheckinLinksRequest {

	dR.AirlineCode = airlineCode

	return dR
}

// SetLanguage set language
func (dR *ReferenceDataUrlsCheckinLinksRequest) SetLanguage(language string) *ReferenceDataUrlsCheckinLinksRequest {

	dR.Language = language

	return dR
}

// GetURL returned key=value format for request on api
func (dR ReferenceDataUrlsCheckinLinksRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := referenceDataUrlsCheckinLinksURL

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

	return baseURL + url
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (dR ReferenceDataUrlsCheckinLinksRequest) GetBody(reqType string) io.Reader {
	return nil
}

type ReferenceDataUrlsCheckinLinksResponse struct {
	Meta     Meta            `json:"meta,omitempty"`
	Data     []UrlData       `json:"data,omitempty"`
	Warnings []Warnings      `json:"warnings,omitempty"`
	Errors   []ErrorResponse `json:"errors,omitempty"`
}

// Decode implement Response interface
func (dR *ReferenceDataUrlsCheckinLinksResponse) Decode(rsp []byte) error {

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
