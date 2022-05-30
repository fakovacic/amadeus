package amadeus

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

//FlightOffersSearchRequest

// REQUEST

type ShoppingFlightOffersRequest struct {
	CurrencyCode       string              `json:"currencyCode,omitempty"`
	CarrierCode        string              `json:"operating,omitempty"`
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

// SetCarrier set carrier code
func (sR *ShoppingFlightOffersRequest) SetCarrier(carrier string) *ShoppingFlightOffersRequest {

	sR.CarrierCode = carrier

	return sR
}

// SetCurrency set currency code
func (sR *ShoppingFlightOffersRequest) SetCurrency(currency string) *ShoppingFlightOffersRequest {

	sR.CurrencyCode = currency

	return sR
}

// SetSources set sources
func (sR *ShoppingFlightOffersRequest) SetSources(sources ...string) *ShoppingFlightOffersRequest {

	if len(sources) != 0 {
		sR.Sources = sources
	}

	return sR
}

// AddOriginDestination add new destination to search request
func (sR *ShoppingFlightOffersRequest) AddOriginDestination(origin, destination, departureDate string) *ShoppingFlightOffersRequest {

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

// AddTravelersByType add travelers of certain type
// Traveler type: ADULT CHILD INFANT
func (sR *ShoppingFlightOffersRequest) AddTravelersByType(no int, travelType string) *ShoppingFlightOffersRequest {

	if no == 0 {
		return sR
	}

	paxCount := len(sR.Travelers)

	for i := 0; i < no; i++ {

		paxCount++

		sR.Travelers = append(sR.Travelers, Travelers{
			ID:           strconv.Itoa(paxCount),
			TravelerType: travelType,
		})

	}

	return sR
}

//
// Helper functions
//

// Oneway helper function to define oneway flight search
func (sR *ShoppingFlightOffersRequest) Oneway(origin, destination, departureDate string) *ShoppingFlightOffersRequest {

	sR.AddOriginDestination(origin, destination, departureDate)

	return sR

}

// Return helper function to define return flight search
func (sR *ShoppingFlightOffersRequest) Return(origin, destination, departureDate, returnDate string) *ShoppingFlightOffersRequest {

	sR.AddOriginDestination(origin, destination, departureDate)

	sR.AddOriginDestination(destination, origin, returnDate)

	return sR

}

// AddTravelers helper function to add more traveler type at once
func (sR *ShoppingFlightOffersRequest) AddTravelers(adult, child, infant int) *ShoppingFlightOffersRequest {

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

// GetURL returned key=value format for request on api
func (sR ShoppingFlightOffersRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := shoopingFlightOffersURL

	// add version
	switch reqType {
	case "GET":

		url = "/v2" + url

		// define query params
		queryParams := []string{}

		oDCouunt := len(sR.OriginDestinations)

		if oDCouunt != 0 {

			oD := sR.OriginDestinations[0]

			// add origin
			queryParams = append(queryParams, "originLocationCode="+oD.OriginCode)

			// add destination
			queryParams = append(queryParams, "destinationLocationCode="+oD.DestinationCode)

			// add departure date
			queryParams = append(queryParams, "departureDate="+oD.DepartureDateTimeRange.Date)

			// check if request is for return flight
			if oDCouunt == 2 {

				oDReturn := sR.OriginDestinations[1]

				if oD.OriginCode == oDReturn.DestinationCode && oD.DestinationCode == oDReturn.OriginCode {

					// add return date
					queryParams = append(queryParams, "returnDate="+oDReturn.DepartureDateTimeRange.Date)

				}

			}

		}

		travelerCount := len(sR.Travelers)

		if travelerCount != 0 {

			var adults int
			var children int
			var infants int

			for _, traveler := range sR.Travelers {

				switch traveler.TravelerType {
				case "ADULT":
					adults++
					break
				case "CHILD":
					children++
					break
				case "INFANT":
					infants++
					break
				}

			}

			queryParams = append(queryParams, "adults="+fmt.Sprintf("%v", adults))
			queryParams = append(queryParams, "children="+fmt.Sprintf("%v", children))
			queryParams = append(queryParams, "infants="+fmt.Sprintf("%v", infants))

		}

		if sR.CurrencyCode != "" {
			queryParams = append(queryParams, "currencyCode="+sR.CurrencyCode)
		}

		if sR.CarrierCode != "" {
			queryParams = append(queryParams, "includedAirlineCodes="+sR.CarrierCode)
		}

		// TODO
		// queryParams = append(queryParams, "travelClass=true")
		// queryParams = append(queryParams, "includedAirlineCodes=true")
		// queryParams = append(queryParams, "excludedAirlineCodes="+dR.ViewBy)
		// queryParams = append(queryParams, "nonStop="+dR.ViewBy)
		// queryParams = append(queryParams, "nonStop="+dR.ViewBy)
		// queryParams = append(queryParams, "maxPrice="+dR.ViewBy)
		// queryParams = append(queryParams, "max="+dR.ViewBy)

		url = url + "?" + strings.Join(queryParams, "&")

		break
	case "POST":

		url = "/v2" + url
		break
	}

	return baseURL + url
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (sR ShoppingFlightOffersRequest) GetBody(reqType string) io.Reader {

	switch reqType {
	case "POST":
		reqPayload, err := json.Marshal(sR)
		if err != nil {
			return nil
		}

		return strings.NewReader(string(reqPayload))
	}

	return nil
}

// RESPONSE

type ShoppingFlightOffersResponse struct {
	Meta         Meta            `json:"meta,omitempty"`
	Data         []FlightOffer   `json:"data,omitempty"`
	Dictionaries Dictionaries    `json:"dictionaries,omitempty"`
	Errors       []ErrorResponse `json:"errors,omitempty"`
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

// Decode implement Response interface
func (dR *ShoppingFlightOffersResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

// GetOffer return offer from list
func (dR ShoppingFlightOffersResponse) GetOffer(offerNum int) FlightOffer {
	return dR.Data[offerNum]
}
