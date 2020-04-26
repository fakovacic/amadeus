package shooping

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"

	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/structs"
)

const (

	///////////
	// 	AIR	 //
	///////////

	// Flight Inspiration Search
	//
	flightDestinationsURL = "/shopping/flight-destinations"
)

type FlightDestinationRequest struct {
	Origin string `json:"origin"`
	//DepartureDateFrom string `json:"departureDateFrom"`
	//DepartureDateTo   string `json:"departureDateTo"`
	DepartureDate string `json:"departureDate"`
	OneWay        bool   `json:"oneWay"`
	Duration      string `json:"duration"`
	//DurationFrom      string `json:"durationFrom"`
	//DurationTo        string `json:"durationTo"`
	NonStop bool `json:"nonStop"`

	// ViewBy possible options
	// DATE, DESTINATION, DURATION, WEEK, or COUNTRY.
	ViewBy   string `json:"viewBy"`
	MaxPrice int    `json:"maxPrice"`
}

// SetOrigin set origin
func (dR *FlightDestinationRequest) SetOrigin(origin string) *FlightDestinationRequest {

	dR.Origin = origin

	return dR
}

// SetDepartureDate set departure
func (dR *FlightDestinationRequest) SetDepartureDate(date string) *FlightDestinationRequest {

	dR.DepartureDate = date

	return dR
}

// IsOneWay set oneway request
func (dR *FlightDestinationRequest) IsOneWay(oneWay string) *FlightDestinationRequest {

	if oneWay == "true" {
		dR.OneWay = true
	}

	return dR
}

// IsNonStop set nonstop request
func (dR *FlightDestinationRequest) IsNonStop(nonStop string) *FlightDestinationRequest {

	if nonStop == "true" {
		dR.NonStop = true
	}

	return dR
}

// // AddDeparture add range of departure dates
// func (dR *FlightDestinationRequest) AddDeparture(from, to string) *FlightDestinationRequest {

// 	dR.DepartureDateFrom = from
// 	dR.DepartureDateTo = to

// 	return dR
// }

// SetDuration add range of duration in days
func (dR *FlightDestinationRequest) SetDuration(duration string) *FlightDestinationRequest {

	dR.Duration = duration

	return dR
}

// // AddDuration add range of duration in days
// func (dR *FlightDestinationRequest) AddDuration(from, to string) *FlightDestinationRequest {

// 	dR.DurationFrom = from
// 	dR.DurationTo = to

// 	return dR
// }

// SetViewBy add grouping option of returned offers
func (dR *FlightDestinationRequest) SetViewBy(view string) *FlightDestinationRequest {

	// check if one of values
	// DATE, DESTINATION, DURATION, WEEK, or COUNTRY.

	dR.ViewBy = view

	return dR
}

// SetMaxPrice add max price of retured offers
func (dR *FlightDestinationRequest) SetMaxPrice(price string) *FlightDestinationRequest {

	// set to int
	//dR.MaxPrice = price

	return dR
}

// ParseParams parse params
func (dR *FlightDestinationRequest) ParseParams(params []string) *FlightDestinationRequest {

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
func (dR *FlightDestinationRequest) SetParam(key, value string) {

	switch key {
	case "origin":
		dR.SetOrigin(value)
		break
	case "departureDate":
		dR.SetDepartureDate(value)
		break
		/*
			case "departureDate":
				dR.SetDestinationLocationCode(value)
				break
			case "departureDate":
				dR.SetDepartureDate(value)
				break
		*/
	case "oneWay":
		dR.IsOneWay(value)
		break

	case "duration":
		dR.SetDuration(value)
		break
		/*
			case "durationFrom":
				dR.SetSearchDate(value)
				break
			case "durationTo":
				dR.SetSearchDate(value)
				break
		*/

	case "nonStop":
		dR.IsNonStop(value)
		break
	case "viewBy":
		dR.SetViewBy(value)
		break
	case "maxPrice":
		dR.SetMaxPrice(value)
		break
	}
}

// GetURL returned key=value format for request on api
func (dR FlightDestinationRequest) GetURL(reqType string) string {

	// set request url
	url := flightDestinationsURL

	// add version
	switch reqType {
	case "GET":

		url = "/v1" + url

		// define query params
		queryParams := []string{}

		queryParams = append(queryParams, "origin="+dR.Origin)

		if dR.OneWay {
			queryParams = append(queryParams, "oneway=true")
		}

		if dR.NonStop {
			queryParams = append(queryParams, "nonStop=true")
		}

		if dR.ViewBy != "" {
			queryParams = append(queryParams, "viewBy="+dR.ViewBy)
		}

		if dR.MaxPrice != 0 {
			queryParams = append(queryParams, "maxPrice="+strconv.Itoa(dR.MaxPrice))
		}

		if dR.DepartureDate != "" {
			queryParams = append(queryParams, "departureDate="+dR.DepartureDate)
		}

		if dR.Duration != "" {
			queryParams = append(queryParams, "duration="+dR.Duration)
		}

		if dR.MaxPrice != 0 {
			queryParams = append(queryParams, "maxPrice="+strconv.Itoa(dR.MaxPrice))
		}

		// if dR.DepartureDateFrom != "" && dR.DepartureDateTo != "" {
		// 	queryParams = append(queryParams, "departureDate="+dR.DepartureDateFrom+","+dR.DepartureDateTo)
		// } else if dR.DepartureDateFrom != "" {
		// 	queryParams = append(queryParams, "departureDate="+dR.DepartureDateFrom)
		// } else if dR.DepartureDateTo != "" {
		// 	queryParams = append(queryParams, "departureDate="+dR.DepartureDateTo)
		// }

		url = url + "?" + strings.Join(queryParams, "&")

		break
	}

	return url
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (dR FlightDestinationRequest) GetBody(reqType string) io.Reader {
	return nil
}

type FlightDestinationResponse struct {
	Meta         structs.Meta         `json:"meta,omitempty"`
	Data         []structs.Data       `json:"data,omitempty"`
	Dictionaries structs.Dictionaries `json:"dictionaries,omitempty"`
}

// Decode implement Response interface
func (dR *FlightDestinationResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}
