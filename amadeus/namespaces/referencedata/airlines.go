package referencedata

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/structs"
)

const (

	// Airline Code Lookup
	airlinesURL = "/reference-data/airlines"
)

type AirlinesRequest struct {
	AirlineCodes []string `json:"airlineCode"`
}

// AddAirlineCode add airlines code
func (dR *AirlinesRequest) AddAirlineCode(airlineCode ...string) *AirlinesRequest {

	dR.AirlineCodes = airlineCode

	return dR
}

// ParseParams parse params
func (dR *AirlinesRequest) ParseParams(params []string) *AirlinesRequest {

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
func (dR *AirlinesRequest) SetParam(key, value string) {

	switch key {
	case "airlineCode":

		//
		// split by ,
		//

		dR.AddAirlineCode(value)
		break
	}
}

// GetURL returned key=value format for request on api
func (dR AirlinesRequest) GetURL(reqType string) string {

	// set request url
	url := airlinesURL

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

	return url
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (dR AirlinesRequest) GetBody(reqType string) io.Reader {
	return nil
}

type AirlinesResponse struct {
	Meta     structs.Meta       `json:"meta,omitempty"`
	Data     []AirlineData      `json:"data,omitempty"`
	Warnings []structs.Warnings `json:"warnings,omitempty"`
}

// Decode implement Response interface
func (dR *AirlinesResponse) Decode(rsp []byte) error {

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
