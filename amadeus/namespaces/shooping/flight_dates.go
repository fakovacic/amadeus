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

	// Flight Cheapest Date Search
	//
	flightDatesURL = "/shopping/flight-dates"
)

type FlightDatesRequest struct {
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	DepartureDate string `json:"departureDate"`
	// DepartureDateFrom string `json:"departureDateFrom"`
	// DepartureDateTo   string `json:"departureDateTo"`
	OneWay   bool   `json:"oneWay"`
	Duration string `json:"duration"`
	// DurationFrom string `json:"durationFrom"`
	// DurationTo   string `json:"durationTo"`
	NonStop bool `json:"nonStop"`

	// ViewBy possible options
	// DATE, DURATION,or WEEK.
	ViewBy   string `json:"viewBy"`
	MaxPrice int    `json:"maxPrice"`
}

// SetOrigin set origin
func (dR *FlightDatesRequest) SetOrigin(origin string) *FlightDatesRequest {

	dR.Origin = origin

	return dR
}

// SetDestination set destination
func (dR *FlightDatesRequest) SetDestination(destination string) *FlightDatesRequest {

	dR.Destination = destination

	return dR
}

// SetDepartureDate set departureDate
func (dR *FlightDatesRequest) SetDepartureDate(date string) *FlightDatesRequest {

	dR.DepartureDate = date

	return dR
}

// SetDuration set departureDate
func (dR *FlightDatesRequest) SetDuration(duration string) *FlightDatesRequest {

	dR.Duration = duration

	return dR
}

// IsOneWay set oneway request
func (dR *FlightDatesRequest) IsOneWay(oneWay string) *FlightDatesRequest {

	if oneWay == "true" {
		dR.OneWay = true
	}

	return dR
}

// IsNonStop set nonstop request
func (dR *FlightDatesRequest) IsNonStop(nonStop string) *FlightDatesRequest {

	if nonStop == "true" {
		dR.NonStop = true
	}

	return dR
}

// AddDeparture add range of departure dates
// func (dR *FlightDatesRequest) AddDeparture(from, to string) *FlightDatesRequest {

// 	dR.DepartureDateFrom = from
// 	dR.DepartureDateTo = to

// 	return dR
// }

// // AddDuration add range of duration in days
// func (dR *FlightDatesRequest) AddDuration(from, to string) *FlightDatesRequest {

// 	dR.DurationFrom = from
// 	dR.DurationTo = to

// 	return dR
// }

// SetViewBy add grouping option of returned offers
func (dR *FlightDatesRequest) SetViewBy(view string) *FlightDatesRequest {

	// check if one of values
	// DATE, DESTINATION, DURATION, WEEK, or COUNTRY.

	dR.ViewBy = view

	return dR
}

// AddMaxPrice add max price of retured offers
func (dR *FlightDatesRequest) SetMaxPrice(price string) *FlightDatesRequest {

	// parse to int
	//dR.MaxPrice = price

	return dR
}

// ParseParams parse params
func (dR *FlightDatesRequest) ParseParams(params []string) *FlightDatesRequest {

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
func (dR *FlightDatesRequest) SetParam(key, value string) {

	switch key {
	case "origin":
		dR.SetOrigin(value)
		break
	case "destination":
		dR.SetDestination(value)
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
func (dR FlightDatesRequest) GetURL(reqType string) string {

	// set request url
	url := flightDatesURL

	// add version
	switch reqType {
	case "GET":

		url = "/v1" + url

		// define query params
		queryParams := []string{}

		queryParams = append(queryParams, "origin="+dR.Origin)
		queryParams = append(queryParams, "destination="+dR.Destination)

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
func (dR FlightDatesRequest) GetBody(reqType string) io.Reader {
	return nil
}

type FlightDatesResponse struct {
	Meta         structs.Meta         `json:"meta"`
	Data         []structs.Data       `json:"data"`
	Dictionaries structs.Dictionaries `json:"dictionaries,omitempty"`
}

// Decode implement Response interface
func (dR *FlightDatesResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}
