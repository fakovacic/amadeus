package amadeus

import (
	"encoding/json"
	"io"
	"strings"
)

type ReferenceDataAirlinesRequest struct {
	AirlineCodes []string `json:"airlineCode"`
}

// AddAirlineCode add airlines code
func (dR *ReferenceDataAirlinesRequest) AddAirlineCode(airlineCode ...string) *ReferenceDataAirlinesRequest {

	dR.AirlineCodes = airlineCode

	return dR
}

// GetURL returned key=value format for request on api
func (dR ReferenceDataAirlinesRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := referenceDataAirlinesURL

	// add version
	switch reqType {
	case "GET":

		url = "/v1" + url

		// define query params
		queryParams := []string{}

		queryParams = append(queryParams, "airlineCodes="+strings.Join(dR.AirlineCodes, ","))

		url = url + "?" + strings.Join(queryParams, "&")

		break
	}

	return baseURL + url
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (dR ReferenceDataAirlinesRequest) GetBody(reqType string) io.Reader {
	return nil
}

type ReferenceDataAirlinesResponse struct {
	Meta     Meta            `json:"meta,omitempty"`
	Data     []AirlineData   `json:"data,omitempty"`
	Warnings []Warnings      `json:"warnings,omitempty"`
	Errors   []ErrorResponse `json:"errors,omitempty"`
}

// Decode implement Response interface
func (dR *ReferenceDataAirlinesResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type AirlineData struct {
	Type         string `json:"typ,omitemptye"`
	IataCode     string `json:"iataCode,omitempty"`
	IcaoCode     string `json:"icaoCode,omitempty"`
	BusinessName string `json:"businessName,omitempty"`
	CommonName   string `json:"commonName,omitempty"`
}
