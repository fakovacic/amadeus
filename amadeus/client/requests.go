package client

import (
	"errors"
	"io"

	airportPredictions "github.com/fakovacic/amadeus-golang/amadeus/namespaces/airport/predictions"
	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/booking"
	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/ereputation"
	files "github.com/fakovacic/amadeus-golang/amadeus/namespaces/media/files"
	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/referencedata"
	pois "github.com/fakovacic/amadeus-golang/amadeus/namespaces/referencedata/locations"
	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/referencedata/urls"
	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/security"
	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/shooping"
	travelAnalytics "github.com/fakovacic/amadeus-golang/amadeus/namespaces/travel/analytics"
	travelPredictions "github.com/fakovacic/amadeus-golang/amadeus/namespaces/travel/predictions"
)

type ApiCall interface {
	Request
	Response
	Get(params ...string) (*Response, error)
	Post(body string) (*Response, error)
	Delete(params ...string) (*Response, error)
}

type Request interface {
	GetURL(reqType string) string
	GetBody(version string) io.Reader
	SetParam(key, value string)
}

type Response interface {
	Decode(rsp []byte) error
}

const (

	//
	// Requests && Response types
	//

	// SecurityToken //
	SecurityToken = iota

	// ShoppingFlightDestination //
	ShoppingFlightDestination

	// ShoppingFlightDates //
	ShoppingFlightDates

	// ShoppingFlightOffers //
	ShoppingFlightOffers

	// ShoppingSeatmaps //
	ShoppingSeatmaps

	// ShoppingFlightPricing //
	ShoppingFlightPricing

	// BookingFlightOrder //
	BookingFlightOrder

	// TravelAnalyticsAirTraffic //
	TravelAnalyticsAirTraffic

	// ReferenceDataLocations //
	ReferenceDataLocations

	// ReferenceDataLocation //
	ReferenceDataLocation

	// ReferenceDataUrlsCheckInLinks //
	ReferenceDataUrlsCheckInLinks

	// ReferenceDataAirlines //
	ReferenceDataAirlines

	// ShoopingHotelsOffers //
	ShoopingHotelsOffers

	// ShoopingHotelOffers //
	ShoopingHotelOffers

	// EReputationHotelSentiments //
	EReputationHotelSentiments

	// BookingHotelBookings //
	BookingHotelBookings

	// ReferenceDataLocationsPois //
	ReferenceDataLocationsPois

	// ReferenceDataLocationsPoi //
	ReferenceDataLocationsPoi

	// TravelPredictionTripPurpose //
	TravelPredictionTripPurpose

	AirportPredictionsOnTime

	MediaFilesGeneratedPhotos
)

// NewRequest return new valid Request interface
func NewRequest(req int) (Request, Response, error) {
	switch req {

	//
	// AIR
	//

	case SecurityToken:
		return new(security.TokenRequest), new(security.TokenResponse), nil

	//
	// AIR
	//

	case ShoppingFlightDestination:
		return new(shooping.FlightDestinationRequest), new(shooping.FlightDestinationResponse), nil
	case ShoppingFlightDates:
		return new(shooping.FlightDatesRequest), new(shooping.FlightDatesResponse), nil
	case ShoppingFlightOffers:
		return new(shooping.FlightOffersRequest), new(shooping.FlightOffersResponse), nil
	case ShoppingSeatmaps:
		return new(shooping.SeatmapsRequest), new(shooping.SeatmapsResponse), nil
	case ShoppingFlightPricing:
		return new(shooping.FlightPricingRequest), new(shooping.FlightPricingResponse), nil
	case BookingFlightOrder:
		return new(booking.FlightOrderRequest), new(booking.FlightOrderResponse), nil
	case TravelAnalyticsAirTraffic:
		return new(travelAnalytics.AirTrafficRequest), new(travelAnalytics.AirTrafficResponse), nil
	case ReferenceDataLocations:
		return new(referencedata.LocationsRequest), new(referencedata.LocationsResponse), nil
	case ReferenceDataLocation:
		return new(referencedata.LocationsRequest), new(referencedata.LocationResponse), nil

	case ReferenceDataUrlsCheckInLinks:
		return new(urls.CheckInLinksRequest), new(urls.CheckInLinksResponse), nil
	case ReferenceDataAirlines:
		return new(referencedata.AirlinesRequest), new(referencedata.AirlinesResponse), nil
	case AirportPredictionsOnTime:
		return new(airportPredictions.OnTimeRequest), new(airportPredictions.OnTimeResponse), nil

	//
	// HOTELS
	//

	case ShoopingHotelsOffers:
		return new(shooping.HotelOffersRequest), new(shooping.HotelsOffersResponse), nil
	case ShoopingHotelOffers:
		return new(shooping.HotelOffersRequest), new(shooping.HotelOffersResponse), nil
	case EReputationHotelSentiments:
		return new(ereputation.HotelSentimentsRequest), new(ereputation.HotelSentimentsResponse), nil
	case BookingHotelBookings:
		return new(booking.HotelBookingsRequest), new(booking.HotelBookingsResponse), nil

	//
	// DESTINATIONS
	//
	case ReferenceDataLocationsPois:
		return new(pois.PoisRequest), new(pois.PoisResponse), nil
	case ReferenceDataLocationsPoi:
		return new(pois.PoisRequest), new(pois.PoiResponse), nil

	//
	// TRIP
	//
	case TravelPredictionTripPurpose:
		return new(travelPredictions.TripPurposeRequest), new(travelPredictions.TripPurposeResponse), nil
	case MediaFilesGeneratedPhotos:
		return new(files.GeneratedPhotosRequest), new(files.GeneratedPhotosResponse), nil

	default:
		return nil, nil, errors.New("request&response not recognized")
	}
}
