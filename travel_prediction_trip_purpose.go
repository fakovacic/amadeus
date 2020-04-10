package amadeus

import (
	"encoding/json"
	"io"
	"strings"
)

type TravelPredictionTripPurposeRequest struct {
	OriginLocationCode      string `json:"originLocationCode"`
	DestinationLocationCode string `json:"destinationLocationCode"`
	DepartureDate           string `json:"departureDate"`
	ReturnDate              string `json:"returnDate"`
	SearchDate              string `json:"searchDate"`
}

// SetOriginLocationCode set origin
func (dR *TravelPredictionTripPurposeRequest) SetOriginLocationCode(code string) *TravelPredictionTripPurposeRequest {

	dR.OriginLocationCode = code

	return dR
}

// SetDestinationLocationCode set destination
func (dR *TravelPredictionTripPurposeRequest) SetDestinationLocationCode(code string) *TravelPredictionTripPurposeRequest {

	dR.DestinationLocationCode = code

	return dR
}

// SetDepartureDate set departure date
func (dR *TravelPredictionTripPurposeRequest) SetDepartureDate(date string) *TravelPredictionTripPurposeRequest {

	dR.DepartureDate = date

	return dR
}

// SetReturnDate set retrun date
func (dR *TravelPredictionTripPurposeRequest) SetReturnDate(date string) *TravelPredictionTripPurposeRequest {

	dR.ReturnDate = date

	return dR
}

// SetSearchDate set search date
func (dR *TravelPredictionTripPurposeRequest) SetSearchDate(date string) *TravelPredictionTripPurposeRequest {

	dR.SearchDate = date

	return dR
}

//
// Helper functions
//

// ReturnFlight helper function to define return flight prediction
func (dR *TravelPredictionTripPurposeRequest) ReturnFlight(
	origin, destination, departureDate, returnDate string) *TravelPredictionTripPurposeRequest {

	dR.SetOriginLocationCode(origin)
	dR.SetDestinationLocationCode(destination)
	dR.SetDepartureDate(departureDate)
	dR.SetReturnDate(returnDate)

	return dR

}

// GetURL returned key=value format for request on api
func (dR TravelPredictionTripPurposeRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := travelPredictionTripPurposeURL

	// add version
	switch reqType {
	case "GET":

		// define query params
		queryParams := []string{}

		url = baseURL + "/v1" + url

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
func (dR TravelPredictionTripPurposeRequest) GetBody(reqType string) io.Reader {
	return nil
}

type TravelPredictionTripPurposeResponse struct {
	Data   TripPurposeData `json:"data,omitempty"`
	Errors []ErrorResponse `json:"errors,omitempty"`
}

// Decode implement Response interface
func (dR *TravelPredictionTripPurposeResponse) Decode(rsp []byte) error {

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
