package amadeus

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type ShoppingFlightDestinationRequest struct {
	Origin            string `json:"origin"`
	DepartureDateFrom string `json:"departureDateFrom"`
	DepartureDateTo   string `json:"departureDateTo"`
	OneWay            bool   `json:"oneWay"`
	DurationFrom      string `json:"durationFrom"`
	DurationTo        string `json:"durationTo"`
	NonStop           bool   `json:"nonStop"`

	// ViewBy possible options
	// DATE, DESTINATION, DURATION, WEEK, or COUNTRY.
	ViewBy   string `json:"viewBy"`
	MaxPrice int    `json:"maxPrice"`
}

// SetOrigin set origin
func (dR *ShoppingFlightDestinationRequest) SetOrigin(origin string) *ShoppingFlightDestinationRequest {

	dR.Origin = origin

	return dR
}

// IsOneWay set oneway request
func (dR *ShoppingFlightDestinationRequest) IsOneWay(oneWay bool) *ShoppingFlightDestinationRequest {

	if oneWay {
		dR.OneWay = true
	}

	return dR
}

// IsNonStop set nonstop request
func (dR *ShoppingFlightDestinationRequest) IsNonStop(nonStop bool) *ShoppingFlightDestinationRequest {

	if nonStop {
		dR.NonStop = true
	}

	return dR
}

// AddDeparture add range of departure dates
func (dR *ShoppingFlightDestinationRequest) AddDeparture(from, to string) *ShoppingFlightDestinationRequest {

	dR.DepartureDateFrom = from
	dR.DepartureDateTo = to

	return dR
}

// AddDuration add range of duration in days
func (dR *ShoppingFlightDestinationRequest) AddDuration(from, to string) *ShoppingFlightDestinationRequest {

	dR.DurationFrom = from
	dR.DurationTo = to

	return dR
}

// AddViewBy add grouping option of returned offers
func (dR *ShoppingFlightDestinationRequest) AddViewBy(view string) *ShoppingFlightDestinationRequest {

	// check if one of values
	// DATE, DESTINATION, DURATION, WEEK, or COUNTRY.

	dR.ViewBy = view

	return dR
}

// AddMaxPrice add max price of retured offers
func (dR *ShoppingFlightDestinationRequest) AddMaxPrice(price int) *ShoppingFlightDestinationRequest {

	dR.MaxPrice = price

	return dR
}

// GetURL returned key=value format for request on api
func (dR ShoppingFlightDestinationRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := shoopingFlightDestinationsURL

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
func (dR ShoppingFlightDestinationRequest) GetBody(reqType string) io.Reader {
	return nil
}

type ShoppingFlightDestinationResponse struct {
	Meta         Meta            `json:"meta,omitempty"`
	Data         []Data          `json:"data,omitempty"`
	Dictionaries Dictionaries    `json:"dictionaries,omitempty"`
	Errors       []ErrorResponse `json:"errors,omitempty"`
}

// Decode implement Response interface
func (dR *ShoppingFlightDestinationResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}
