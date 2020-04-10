package amadeus

import (
	"encoding/json"
	"io"
	"strings"
)

type AirportPredictionsOnTimeRequest struct {
	AirportCode string `json:"aiportCode"`
	Date        string `json:"date"`
}

// SetAirportCode set airport code
func (dR *AirportPredictionsOnTimeRequest) SetAirportCode(airlineCode string) *AirportPredictionsOnTimeRequest {

	dR.AirportCode = airlineCode

	return dR
}

// SetDate set date
func (dR *AirportPredictionsOnTimeRequest) SetDate(date string) *AirportPredictionsOnTimeRequest {

	dR.Date = date

	return dR
}

// GetURL returned key=value format for request on api
func (dR AirportPredictionsOnTimeRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := airportPredictionsOnTimeURL

	// add version
	switch reqType {
	case "GET":

		url = "/v1" + url

		// define query params
		queryParams := []string{}

		queryParams = append(queryParams, "airportCode="+dR.AirportCode)

		queryParams = append(queryParams, "date="+dR.Date)

		url = url + "?" + strings.Join(queryParams, "&")

		break
	}

	return baseURL + url
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (dR AirportPredictionsOnTimeRequest) GetBody(reqType string) io.Reader {
	return nil
}

type AirportPredictionsOnTimeResponse struct {
	Meta   Meta            `json:"meta,omitempty"`
	Data   AirportData     `json:"data,omitempty"`
	Errors []ErrorResponse `json:"errors,omitempty"`
}

// Decode implement Response interface
func (dR *AirportPredictionsOnTimeResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type AirportData struct {
	ID          string `json:"id,omitempty"`
	Probability string `json:"probability,omitempty"`
	Result      string `json:"result,omitempty"`
	SubType     string `json:"subType,omitempty"`
	Type        string `json:"type,omitempty"`
}
