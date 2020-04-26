package shooping

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/structs"
)

const (

	///////////
	// 	AIR	 //
	///////////

	// Flight Offers Search
	// search for offers on given origin, destination, departure, passangers
	flightOffersURL = "/shopping/flight-offers"
)

//FlightOffersSearchRequest

// REQUEST

type FlightOffersRequest struct {
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

// sR.CurrencyCode = currency

// SetCurrency set currency code
func (sR *FlightOffersRequest) SetCurrency(currency string) *FlightOffersRequest {

	sR.CurrencyCode = currency

	return sR
}

// SetSources set sources
func (sR *FlightOffersRequest) AddSources(sources ...string) *FlightOffersRequest {

	if len(sources) != 0 {
		sR.Sources = sources
	}

	return sR
}

// AddOriginDestination add new destination to search request
func (sR *FlightOffersRequest) AddOriginDestination(origin, destination, departureDate string) *FlightOffersRequest {

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
func (sR *FlightOffersRequest) AddTravelersByType(no int, travelType string) *FlightOffersRequest {

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
func (sR *FlightOffersRequest) Oneway(origin, destination, departureDate string) *FlightOffersRequest {

	sR.AddOriginDestination(origin, destination, departureDate)

	return sR

}

// Return helper function to define return flight search
func (sR *FlightOffersRequest) Return(origin, destination, departureDate, returnDate string) *FlightOffersRequest {

	sR.AddOriginDestination(origin, destination, departureDate)

	sR.AddOriginDestination(destination, origin, returnDate)

	return sR

}

// AddTravelers helper function to add more traveler type at once
func (sR *FlightOffersRequest) AddTravelers(adult, child, infant int) *FlightOffersRequest {

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

// ParseParams parse params
func (dR *FlightOffersRequest) ParseParams(params []string) *FlightOffersRequest {

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

// SetParam set params
func (sR *FlightOffersRequest) SetParam(key, value string) {

	switch key {
	case "currency":
		sR.SetCurrency(value)
		break
	case "sources":

		//
		//	parse by ,
		//

		sR.AddSources(value)
		break

	case "originLocationCode":

		originDestinationCount := len(sR.OriginDestinations)

		switch originDestinationCount {
		case 0:
			sR.OriginDestinations = append(sR.OriginDestinations, OriginDestination{
				ID:         strconv.Itoa(originDestinationCount),
				OriginCode: value,
			})
			break
		case 1:
			sR.OriginDestinations[0].OriginCode = value
			break
		case 2:
			sR.OriginDestinations[0].OriginCode = value
			sR.OriginDestinations[0].DestinationCode = value
			break
		}

		break
	case "destinationLocationCode":

		originDestinationCount := len(sR.OriginDestinations)

		switch originDestinationCount {
		case 0:
			sR.OriginDestinations = append(sR.OriginDestinations, OriginDestination{
				ID:              strconv.Itoa(originDestinationCount),
				DestinationCode: value,
			})
			break
		case 1:
			sR.OriginDestinations[0].DestinationCode = value
			break
		case 2:
			sR.OriginDestinations[0].DestinationCode = value
			sR.OriginDestinations[0].OriginCode = value
			break
		}

		break
	case "departureDate":

		originDestinationCount := len(sR.OriginDestinations)

		switch originDestinationCount {
		case 0:
			sR.OriginDestinations = append(sR.OriginDestinations, OriginDestination{
				ID: strconv.Itoa(originDestinationCount),
				DepartureDateTimeRange: TimeRange{
					Date: value,
				},
			})
			break
		case 1:

			sR.OriginDestinations[0].DepartureDateTimeRange.Date = value

			break
		case 2:

			sR.OriginDestinations[0].DepartureDateTimeRange.Date = value

			break
		}

		break
	case "returnDate":

		originDestinationCount := len(sR.OriginDestinations)

		switch originDestinationCount {
		case 0:
			sR.OriginDestinations = append(sR.OriginDestinations, OriginDestination{
				ID: strconv.Itoa(originDestinationCount),
			})

			originDestinationCount++

			sR.OriginDestinations = append(sR.OriginDestinations, OriginDestination{
				ID: strconv.Itoa(originDestinationCount),
				DepartureDateTimeRange: TimeRange{
					Date: value,
				},
			})
			break
		case 1:

			originDestinationCount++

			sR.OriginDestinations = append(sR.OriginDestinations, OriginDestination{
				ID: strconv.Itoa(originDestinationCount),
				DepartureDateTimeRange: TimeRange{
					Date: value,
				},
			})

			break
		case 2:

			sR.OriginDestinations[1].DepartureDateTimeRange.Date = value

			break
		}

		break
	case "adults":

		// convert to int
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("ADULT number not valid", err)
		}
		sR.AddTravelersByType(num, "ADULT")
		break
	case "children":

		// convert to int
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("CHILD number not valid", err)
		}
		sR.AddTravelersByType(num, "CHILD")
		break
	case "infants":

		// convert to int
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("INFANT number not valid", err)
		}
		sR.AddTravelersByType(num, "INFANT")
		break
	}

	return
}

// ParseBody parse JSON body to request
func (sR *FlightOffersRequest) ParseBody(body string) *FlightOffersRequest {

	err := json.Unmarshal([]byte(body), sR)
	if err != nil {
		return nil
	}

	return sR
}

// GetURL returned key=value format for request on api
func (sR FlightOffersRequest) GetURL(reqType string) string {

	// set request url
	url := flightOffersURL

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

		// TODO
		// queryParams = append(queryParams, "travelClass=")
		// queryParams = append(queryParams, "includedAirlineCodes=")
		// queryParams = append(queryParams, "excludedAirlineCodes=")
		// queryParams = append(queryParams, "nonStop=")
		// queryParams = append(queryParams, "maxPrice=")
		// queryParams = append(queryParams, "max=")

		url = url + "?" + strings.Join(queryParams, "&")

		break
	case "POST":

		url = "/v2" + url
		break
	}

	return url
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (sR FlightOffersRequest) GetBody(reqType string) io.Reader {

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

type FlightOffersResponse struct {
	Meta         structs.Meta          `json:"meta,omitempty"`
	Data         []structs.FlightOffer `json:"data,omitempty"`
	Dictionaries structs.Dictionaries  `json:"dictionaries,omitempty"`
}

// Decode implement Response interface
func (dR *FlightOffersResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

// GetFlightOffer return offer from list
func (dR FlightOffersResponse) GetFlightOffer(offerNum int) structs.FlightOffer {
	return dR.Data[offerNum]
}
