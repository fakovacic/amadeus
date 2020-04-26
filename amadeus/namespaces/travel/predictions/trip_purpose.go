package predictions

import (
	"encoding/json"
	"io"
	"strings"
)

const (

	// Trip Purpose Prediction
	// The Trip Purpose Prediction API allows developers to forecast traveler purpose, Business or Leisure,
	// together with the probability in the context of search & shopping
	tripPurposeURL = "/travel/predictions/trip-purpose"
)

type TripPurposeRequest struct {
	OriginLocationCode      string `json:"originLocationCode"`
	DestinationLocationCode string `json:"destinationLocationCode"`
	DepartureDate           string `json:"departureDate"`
	ReturnDate              string `json:"returnDate"`
	SearchDate              string `json:"searchDate"`
}

// SetOriginLocationCode set origin
func (dR *TripPurposeRequest) SetOriginLocationCode(code string) *TripPurposeRequest {

	dR.OriginLocationCode = code

	return dR
}

// SetDestinationLocationCode set destination
func (dR *TripPurposeRequest) SetDestinationLocationCode(code string) *TripPurposeRequest {

	dR.DestinationLocationCode = code

	return dR
}

// SetDepartureDate set departure date
func (dR *TripPurposeRequest) SetDepartureDate(date string) *TripPurposeRequest {

	dR.DepartureDate = date

	return dR
}

// SetReturnDate set retrun date
func (dR *TripPurposeRequest) SetReturnDate(date string) *TripPurposeRequest {

	dR.ReturnDate = date

	return dR
}

// SetSearchDate set search date
func (dR *TripPurposeRequest) SetSearchDate(date string) *TripPurposeRequest {

	dR.SearchDate = date

	return dR
}

//
// Helper functions
//

// ReturnFlight helper function to define return flight prediction
func (dR *TripPurposeRequest) ReturnFlight(
	origin, destination, departureDate, returnDate string) *TripPurposeRequest {

	dR.SetOriginLocationCode(origin)
	dR.SetDestinationLocationCode(destination)
	dR.SetDepartureDate(departureDate)
	dR.SetReturnDate(returnDate)

	return dR

}

// ParseParams parse params
func (dR *TripPurposeRequest) ParseParams(params []string) *TripPurposeRequest {

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
func (dR *TripPurposeRequest) SetParam(key, value string) {

	switch key {
	case "originLocationCode":
		dR.SetOriginLocationCode(value)
		break
	case "destinationLocationCode":
		dR.SetDestinationLocationCode(value)
		break
	case "departureDate":
		dR.SetDepartureDate(value)
		break
	case "returnDate":
		dR.SetReturnDate(value)
		break
	case "searchDate":
		dR.SetSearchDate(value)
		break
	}
}

// GetURL returned key=value format for request on api
func (dR TripPurposeRequest) GetURL(reqType string) string {

	// set request url
	url := tripPurposeURL

	// add version
	switch reqType {
	case "GET":

		// define query params
		queryParams := []string{}

		url = "/v1" + url

		queryParams = append(queryParams, "originLocationCode="+dR.OriginLocationCode)
		queryParams = append(queryParams, "destinationLocationCode="+dR.DestinationLocationCode)
		queryParams = append(queryParams, "departureDate="+dR.DepartureDate)
		queryParams = append(queryParams, "returnDate="+dR.ReturnDate)

		if dR.SearchDate != "" {
			queryParams = append(queryParams, "searchDate="+dR.SearchDate)
		}

		return url + "?" + strings.Join(queryParams, "&")
	}

	return ""
}

// GetBody implementation for Request
func (dR TripPurposeRequest) GetBody(reqType string) io.Reader {
	return nil
}

type TripPurposeResponse struct {
	Data TripPurposeData `json:"data,omitempty"`
}

// Decode implement Response interface
func (dR *TripPurposeResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type TripPurposeData struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	Subtype     string `json:"subtype,omitempty"`
	Result      string `json:"result,omitempty"`
	Probability string `json:"probability,omitempty"`
}
