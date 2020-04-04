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
	Count int `json:"count,omitempty"`
}

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

func NewSearchRequest(currency string, sources ...string) FlightOffersSearchRequest {

	var sR FlightOffersSearchRequest

	sR.CurrencyCode = currency

	if len(sources) != 0 {
		sR.Sources = sources
	}

	return sR
}

func (sR *FlightOffersSearchRequest) Oneway(origin, destination, departureDate string) {

	sR.OriginDestinations = append(sR.OriginDestinations, OriginDestination{
		ID:              "1",
		OriginCode:      origin,
		DestinationCode: destination,
		DepartureDateTimeRange: TimeRange{
			Date: departureDate,
		},
	})

}

func (sR *FlightOffersSearchRequest) Return(origin, destination, departureDate, returnDate string) {

	sR.OriginDestinations = append(sR.OriginDestinations, OriginDestination{
		ID:              "1",
		OriginCode:      origin,
		DestinationCode: destination,
		DepartureDateTimeRange: TimeRange{
			Date: departureDate,
		},
	})

	sR.OriginDestinations = append(sR.OriginDestinations, OriginDestination{
		ID:              "2",
		OriginCode:      destination,
		DestinationCode: origin,
		DepartureDateTimeRange: TimeRange{
			Date: returnDate,
		},
	})

}

// ADD ORIGIN DESTINATION

func (sR *FlightOffersSearchRequest) Multi(origin, destination, departureDate, returnDate string) {
	//TODO
}

func (sR *FlightOffersSearchRequest) AddTravelers(adult, child, infant int) {

	paxCount := 1

	if adult != 0 {

		for i := 0; i <= adult; i++ {

			sR.Travelers = append(sR.Travelers, Travelers{
				ID:           strconv.Itoa(paxCount),
				TravelerType: "ADULT",
			})

			paxCount++
		}

	}

	if child != 0 {

		for i := 0; i <= child; i++ {

			sR.Travelers = append(sR.Travelers, Travelers{
				ID:           strconv.Itoa(paxCount),
				TravelerType: "CHILD",
			})

			paxCount++
		}

	}

	if infant != 0 {

		for i := 0; i <= infant; i++ {

			sR.Travelers = append(sR.Travelers, Travelers{
				ID:           strconv.Itoa(paxCount),
				TravelerType: "INFANT",
			})

			paxCount++
		}

	}

}

// Add traveler
