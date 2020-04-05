package amadeus

import (
	"encoding/json"
	"strconv"
)

//FlightOffersSearchRequest

// REQUEST

type FlightOffersSearchRequest struct {
	CurrencyCode       string              `json:"currencyCode,omitempty"`
	OriginDestinations []OriginDestination `json:"originDestinations,omitempty"`
	Travelers          []Travelers         `json:"travelers,omitempty"`
	Sources            []string            `json:"sources,omitempty"`
	SearchCriteria     SearchCriteria      `json:"searchCriteria,omitempty"`
}

type OriginDestination struct {
	ID                     string    `json:"id,omitempty"`
	OriginCode             string    `json:"originLocationCode,omitempty"`
	DestinationCode        string    `json:"destinationLocationCode,omitempty"`
	DepartureDateTimeRange TimeRange `json:"departureDateTimeRange,omitempty"`
}

type TimeRange struct {
	Date string `json:"date,omitempty"`
	Time string `json:"time,omitempty"`
}

type Travelers struct {
	ID           string `json:"id,omitempty"`
	TravelerType string `json:"travelerType,omitempty"`
}

type SearchCriteria struct {
	MaxFlightOffers int           `json:"maxFlightOffers,omitempty"`
	FlightFilters   FlightFilters `json:"flightFilters,omitempty"`
}

type FlightFilters struct {
	CabinRestrictions   []CabinRestrictions `json:"cabinRestrictions,omitempty"`
	CarrierRestrictions CarrierRestrictions `json:"carrierRestrictions,omitempty"`
}

type CabinRestrictions struct {
	Cabin                string   `json:"cabin,omitempty"`
	Coverage             string   `json:"coverage,omitempty"`
	OriginDestinationIds []string `json:"originDestinationIds,omitempty"`
}
type CarrierRestrictions struct {
	ExcludedCarrierCodes []string `json:"excludedCarrierCodes,omitempty"`
}

// RESPONSE

type FlightOffersSearchResponse struct {
	Meta   Meta            `json:"meta,omitempty"`
	Data   []FlightOffer   `json:"data,omitempty"`
	Errors []ErrorResponse `json:"errors,omitempty"`
}

type Meta struct {
	Count    int      `json:"count,omitempty"`
	Currency string   `json:"currency,omitempty"`
	Links    Links    `json:"links,omitempty"`
	Defaults Defaults `json:"defaults,omitempty"`
}

type Defaults struct {
	DepartureDate string `json:"departureDate,omitempty"`
	OneWay        bool   `json:"oneWay,omitempty"`
	Duration      string `json:"duration,omitempty"`
	NonStop       bool   `json:"nonStop,omitempty"`
	ViewBy        string `json:"viewBy,omitempty"`
}

// FlightOffers send request to api to retrive flight offers
func (a *Amadeus) FlightOffers(request FlightOffersSearchRequest) (FlightOffersSearchResponse, error) {

	var response FlightOffersSearchResponse

	urlStr, err := a.getURL(shoopingFlightOffers)
	if err != nil {
		return response, err
	}

	reqPayload, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	resp, err := a.postRequest(string(reqPayload), urlStr)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// NewSearchRequest construct flight search request
func NewSearchRequest(currency string, sources ...string) *FlightOffersSearchRequest {

	var sR FlightOffersSearchRequest

	sR.CurrencyCode = currency

	if len(sources) != 0 {
		sR.Sources = sources
	}

	return &sR
}

// Oneway helper function to define oneway flight search
func (sR *FlightOffersSearchRequest) Oneway(origin, destination, departureDate string) *FlightOffersSearchRequest {

	sR.AddOriginDestination(origin, destination, departureDate)

	return sR

}

// Return helper function to define return flight search
func (sR *FlightOffersSearchRequest) Return(origin, destination, departureDate, returnDate string) *FlightOffersSearchRequest {

	sR.AddOriginDestination(origin, destination, departureDate)

	sR.AddOriginDestination(origin, destination, returnDate)

	return sR

}

// AddOriginDestination add new destination to search request
func (sR *FlightOffersSearchRequest) AddOriginDestination(origin, destination, departureDate string) *FlightOffersSearchRequest {

	// check origin

	// check destination

	// check departureDate

	originDestinationCount := len(sR.OriginDestinations)

	originDestinationCount++

	sR.OriginDestinations = append(sR.OriginDestinations, OriginDestination{
		ID:              strconv.Itoa(originDestinationCount),
		OriginCode:      origin,
		DestinationCode: destination,
		DepartureDateTimeRange: TimeRange{
			Date: departureDate,
		},
	})

	return sR

}

// AddTravelers helper function to add more traveler type
func (sR *FlightOffersSearchRequest) AddTravelers(adult, child, infant int) *FlightOffersSearchRequest {

	if adult != 0 {

		sR.AddTravelersByType(adult, "ADULT")

	}

	if child != 0 {

		sR.AddTravelersByType(child, "CHILD")

	}

	if infant != 0 {

		sR.AddTravelersByType(infant, "INFANT")

	}

	return sR
}

// AddTravelersByType add travelers of certain type
// Traveler type: ADULT CHILD INFANT
func (sR *FlightOffersSearchRequest) AddTravelersByType(no int, travelType string) *FlightOffersSearchRequest {

	if no == 0 {
		return sR
	}

	paxCount := len(sR.Travelers)

	for i := 0; i <= no; i++ {

		paxCount++

		sR.Travelers = append(sR.Travelers, Travelers{
			ID:           strconv.Itoa(paxCount),
			TravelerType: travelType,
		})

	}

	return sR
}
