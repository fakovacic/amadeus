package amadeus

import (
	"encoding/json"
	"strconv"
)

type FlightDestinationsRequest struct {
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

type FlightDestinationResponse struct {
	Data []Data `json:"data,omitempty"`
	Meta Meta   `json:"meta,omitempty"`
}

type Data struct {
	Type          string `json:"type,omitempty"`
	Origin        string `json:"origin,omitempty"`
	Destination   string `json:"destination,omitempty"`
	DepartureDate string `json:"departureDate,omitempty"`
	ReturnDate    string `json:"returnDate,omitempty"`
	Price         Price  `json:"price,omitempty"`
	Links         Links  `json:"links,omitempty"`
}

type Links struct {
	FlightDates  string `json:"flightDates,omitempty"`
	FlightOffers string `json:"flightOffers,omitempty"`
	Self         string `json:"self,omitempty"`
}

// FlightDestinations send request to api to retrive flight inspiration
func (a *Amadeus) FlightDestinations(request FlightDestinationsRequest) (FlightDestinationResponse, error) {

	var response FlightDestinationResponse

	urlStr, err := a.getURL(shoopingFlightDestinations)
	if err != nil {
		return response, err
	}

	// prepare request params
	qParams := request.PrepareQueryParamsForRequest()

	resp, err := a.getRequest(urlStr, qParams)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// NewFlightDestinationsRequest construct flight destination request
func NewFlightDestinationsRequest(origin string, oneWay, nonStop bool) *FlightDestinationsRequest {

	var dR FlightDestinationsRequest

	dR.Origin = origin

	if oneWay {
		dR.OneWay = true
	}

	if nonStop {
		dR.NonStop = true
	}

	return &dR
}

// AddDeparture add range of departure dates
func (dR *FlightDestinationsRequest) AddDeparture(from, to string) *FlightDestinationsRequest {

	dR.DepartureDateFrom = from
	dR.DepartureDateTo = to

	return dR
}

// AddDuration add range of duration in days
func (dR *FlightDestinationsRequest) AddDuration(from, to string) *FlightDestinationsRequest {

	dR.DurationFrom = from
	dR.DurationTo = to

	return dR
}

// AddViewBy add grouping option of returned offers
func (dR *FlightDestinationsRequest) AddViewBy(view string) *FlightDestinationsRequest {

	// check if one of values
	// DATE, DESTINATION, DURATION, WEEK, or COUNTRY.

	dR.ViewBy = view

	return dR
}

// AddMaxPrice add max price of retured offers
func (dR *FlightDestinationsRequest) AddMaxPrice(price int) *FlightDestinationsRequest {

	dR.MaxPrice = price

	return dR
}

// PrepareQueryParamsForRequest prepare struct values to slice
// returned key=value format for request on api
func (dR FlightDestinationsRequest) PrepareQueryParamsForRequest() []string {

	var queryParams []string

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

	return queryParams
}
