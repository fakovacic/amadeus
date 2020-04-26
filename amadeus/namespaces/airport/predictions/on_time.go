package predictions

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/structs"
)

const (

	//
	// Artificial Intelligence
	//

	// PredictionsOnTime
	onTimeURL = "/airport/predictions/on-time"
)

type OnTimeRequest struct {
	AirportCode string `json:"aiportCode"`
	Date        string `json:"date"`
}

// SetAirportCode set airport code
func (dR *OnTimeRequest) SetAirportCode(airlineCode string) {

	dR.AirportCode = airlineCode
}

// SetDate set date
func (dR *OnTimeRequest) SetDate(date string) {

	dR.Date = date

}

// ParseParams parse params
func (dR *OnTimeRequest) ParseParams(params []string) *OnTimeRequest {

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
func (dR *OnTimeRequest) SetParam(key, value string) {

	switch key {
	case "airportCode":
		dR.SetAirportCode(value)
		break
	case "date":
		dR.SetDate(value)
		break
	}
}

// GetURL returned key=value format for request on api
func (dR OnTimeRequest) GetURL(reqType string) string {

	// set request url
	url := onTimeURL

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

	return url
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (dR OnTimeRequest) GetBody(reqType string) io.Reader {
	return nil
}

type OnTimeResponse struct {
	Meta structs.Meta `json:"meta,omitempty"`
	Data AirportData  `json:"data,omitempty"`
}

// Decode implement Response interface
func (dR *OnTimeResponse) Decode(rsp []byte) error {

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
