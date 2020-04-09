package amadeus

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type ShoppingFlightDatesRequest struct {
	Origin            string `json:"origin"`
	Destination       string `json:"destination"`
	DepartureDateFrom string `json:"departureDateFrom"`
	DepartureDateTo   string `json:"departureDateTo"`
	OneWay            bool   `json:"oneWay"`
	DurationFrom      string `json:"durationFrom"`
	DurationTo        string `json:"durationTo"`
	NonStop           bool   `json:"nonStop"`

	// ViewBy possible options
	// DATE, DURATION,or WEEK.
	ViewBy   string `json:"viewBy"`
	MaxPrice int    `json:"maxPrice"`
}

// SetOrigin set origin
func (dR *ShoppingFlightDatesRequest) SetOrigin(origin string) *ShoppingFlightDatesRequest {

	dR.Origin = origin

	return dR
}

// SetDestination set destination
func (dR *ShoppingFlightDatesRequest) SetDestination(destination string) *ShoppingFlightDatesRequest {

	dR.Destination = destination

	return dR
}

// IsOneWay set oneway request
func (dR *ShoppingFlightDatesRequest) IsOneWay(oneWay bool) *ShoppingFlightDatesRequest {

	if oneWay {
		dR.OneWay = true
	}

	return dR
}

// IsNonStop set nonstop request
func (dR *ShoppingFlightDatesRequest) IsNonStop(nonStop bool) *ShoppingFlightDatesRequest {

	if nonStop {
		dR.NonStop = true
	}

	return dR
}

// AddDeparture add range of departure dates
func (dR *ShoppingFlightDatesRequest) AddDeparture(from, to string) *ShoppingFlightDatesRequest {

	dR.DepartureDateFrom = from
	dR.DepartureDateTo = to

	return dR
}

// AddDuration add range of duration in days
func (dR *ShoppingFlightDatesRequest) AddDuration(from, to string) *ShoppingFlightDatesRequest {

	dR.DurationFrom = from
	dR.DurationTo = to

	return dR
}

// AddViewBy add grouping option of returned offers
func (dR *ShoppingFlightDatesRequest) AddViewBy(view string) *ShoppingFlightDatesRequest {

	// check if one of values
	// DATE, DESTINATION, DURATION, WEEK, or COUNTRY.

	dR.ViewBy = view

	return dR
}

// AddMaxPrice add max price of retured offers
func (dR *ShoppingFlightDatesRequest) AddMaxPrice(price int) *ShoppingFlightDatesRequest {

	dR.MaxPrice = price

	return dR
}

// GetURL returned key=value format for request on api
func (dR ShoppingFlightDatesRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := shoopingFlightDatesURL

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

		if dR.DepartureDateFrom != "" && dR.DepartureDateTo != "" {
			queryParams = append(queryParams, "departureDate="+dR.DepartureDateFrom+","+dR.DepartureDateTo)
		} else if dR.DepartureDateFrom != "" {
			queryParams = append(queryParams, "departureDate="+dR.DepartureDateFrom)
		} else if dR.DepartureDateTo != "" {
			queryParams = append(queryParams, "departureDate="+dR.DepartureDateTo)
		}

		url = url + "?" + strings.Join(queryParams, "&")

		break
	}

	return baseURL + url
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (dR ShoppingFlightDatesRequest) GetBody(reqType string) io.Reader {
	return nil
}

type ShoppingFlightDatesResponse struct {
	Meta         Meta            `json:"meta"`
	Data         []Data          `json:"data"`
	Dictionaries Dictionaries    `json:"dictionaries,omitempty"`
	Errors       []ErrorResponse `json:"errors,omitempty"`
}

// Decode implement Response interface
func (dR *ShoppingFlightDatesResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}
