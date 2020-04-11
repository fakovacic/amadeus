package amadeus

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (

	//
	// Base URLs
	//

	// Testing API url
	testURL = "https://test.api.amadeus.com"

	// Production API url
	productionURL = "https://api.amadeus.com"

	//
	// Requests URLs
	//

	// Authentification url
	// use to aquire token which is used in all other request
	securityOAuth2TokenURL = "/v1/security/oauth2/token"

	///////////
	// 	AIR	 //
	///////////

	//
	// Shooping
	//

	// Flight Offers Search
	// search for offers on given origin, destination, departure, passangers
	shoopingFlightOffersURL = "/shopping/flight-offers"

	// Flight Inspiration Search
	//
	shoopingFlightDestinationsURL = "/shopping/flight-destinations"

	// Flight Cheapest Date Search
	//
	shoopingFlightDatesURL = "/shopping/flight-dates"

	// Shooping Flight offers pricing
	// check certain offer if is still active, response with additional data for offer
	shoopingFlightOffersPricingURL = "/shopping/flight-offers/pricing"

	// Seatmap Display
	//
	shoopingSeatmapsURL = "/shopping/seatmaps"

	//
	// Booking
	//

	// Booking Flight orders
	// create reservation for certain offer
	bookingFlightOrdersURL = "/booking/flight-orders"

	//
	// Travel Insights
	//

	// Flight Busiest Traveling Period
	// Flight Most Booked Destinations
	// Flight Most Traveled Destinations
	travelAnalyticsAirTrafficURL = "/travel/analytics/air-traffic"

	//
	// Reference data
	//

	// Airport & City Search
	// Airport Nearest Relevant
	referenceDataLocationsURL = "/reference-data/locations"

	// Flight Check-in Links
	referenceDataUrlsCheckinLinksURL = "/reference-data/urls/checkin-links"

	// Airline Code Lookup
	referenceDataAirlinesURL = "/reference-data/airlines"

	//
	// Artificial Intelligence
	//

	// AirportPredictionsOnTime
	airportPredictionsOnTimeURL = "/airport/predictions/on-time"

	/////////////
	// 	HOTEL  //
	/////////////

	//
	// Shooping
	//

	// Hotel Search
	// search for hotel offers on destination
	shoopingHotelOffersURL = "/shopping/hotel-offers"

	// Hotel Ratings
	// get hotel reputation by hotel ids
	ereputationHotelSentimentsURL = "/e-reputation/hotel-sentiments"

	// Hotel Booking
	// allows you to book the best hotel offers you have searched from a wide choice of providers
	bookingHotelBookingsURL = "/booking/hotel-bookings"

	////////////////////
	// 	DESTINATIONS  //
	////////////////////

	// Points of Interest
	referenceDataLocationsPoisURL = "/reference-data/locations/pois"

	////////////
	// 	TRIP  //
	////////////

	// Trip Purpose Prediction
	// The Trip Purpose Prediction API allows developers to forecast traveler purpose, Business or Leisure,
	// together with the probability in the context of search & shopping
	travelPredictionTripPurposeURL = "/travel/predictions/trip-purpose"

	// AI Generated Photos
	// The AI-Generated Photos API returns a link to download a rendered image of a landscape.
	mediaFilesGeneratedPhotosURL = "/media/files/generated-photos"

	//
	// Requests && Response types
	//

	// ShoppingFlightDestination //
	ShoppingFlightDestination = iota

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

	// ReferenceDataUrlsCheckinLinks //
	ReferenceDataUrlsCheckinLinks

	// ReferenceDataAirlines //
	ReferenceDataAirlines

	// AirportPredictionsOnTime //
	AirportPredictionsOnTime

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

	// MediaFilesGeneratedPhotos //
	MediaFilesGeneratedPhotos
)

// Amadeus main struct that holds sensitive data for communicating with api
// key, secret and env for requesting token for authentification which is used in all other requests
type Amadeus struct {
	key    string
	secret string
	env    string
	token  token
}

// New creates new amadeus client for given Key, Secret & Environment
// Key & Secret are created on amadeus developers page https://developers.amadeus.com/register
func New(Key, Secret, ENV string) (*Amadeus, error) {

	var (
		a   Amadeus
		err error
	)

	err = a.setKey(Key)
	if err != nil {
		return nil, err
	}

	err = a.setSecret(Secret)
	if err != nil {
		return nil, err
	}

	err = a.setENV(ENV)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

// setKey for field key in Amadeus struct
// check if empty than return error
func (a *Amadeus) setKey(value string) error {

	if value == "" {
		return errors.New("key is empty")
	}
	a.key = value
	return nil
}

// setSecret for field secret in Amadeus struct
// check if empty than return error
func (a *Amadeus) setSecret(value string) error {

	if value == "" {
		return errors.New("secret is empty")
	}
	a.secret = value
	return nil
}

// setENV for env field in Amadeus struct
// check if empty than return error
// check if valid environment for using base const url
func (a *Amadeus) setENV(value string) error {

	if value == "" {
		return errors.New("env is empty")
	}

	switch value {
	case "TEST":
		a.env = value
		return nil
	case "PRODUCTION":
		a.env = value
		return nil
	default:
		return errors.New("env not set")
	}

}

// getURL return full url for given endpoint
// checks for environment base url and add endpoint url
func (a Amadeus) getBaseURL() (string, error) {

	switch a.env {
	case "TEST":
		return testURL, nil
	case "PRODUCTION":
		return productionURL, nil
	}

	return "", errors.New("not defined valid environment")
}

type Request interface {
	GetURL(baseURL, reqType string) string
	GetBody(version string) io.Reader
}

type Response interface {
	Decode(rsp []byte) error
}

// NewRequest return new valid Request interface
func (a *Amadeus) NewRequest(req int) (Request, Response, error) {
	switch req {

	//
	// AIR
	//

	case ShoppingFlightDestination:
		return new(ShoppingFlightDestinationRequest), new(ShoppingFlightDestinationResponse), nil
	case ShoppingFlightDates:
		return new(ShoppingFlightDatesRequest), new(ShoppingFlightDatesResponse), nil
	case ShoppingFlightOffers:
		return new(ShoppingFlightOffersRequest), new(ShoppingFlightOffersResponse), nil
	case ShoppingSeatmaps:
		return new(ShoppingSeatmapsRequest), new(ShoppingSeatmapsResponse), nil
	case ShoppingFlightPricing:
		return new(ShoppingFlightPricingRequest), new(ShoppingFlightPricingResponse), nil
	case BookingFlightOrder:
		return new(BookingFlightOrderRequest), new(BookingFlightOrderResponse), nil
	case TravelAnalyticsAirTraffic:
		return new(TravelAnalyticsAirTrafficRequest), new(TravelAnalyticsAirTrafficResponse), nil
	case ReferenceDataLocations:
		return new(ReferenceDataLocationsRequest), new(ReferenceDataLocationsResponse), nil
	case ReferenceDataLocation:
		return new(ReferenceDataLocationsRequest), new(ReferenceDataLocationResponse), nil

	case ReferenceDataUrlsCheckinLinks:
		return new(ReferenceDataUrlsCheckinLinksRequest), new(ReferenceDataUrlsCheckinLinksResponse), nil
	case ReferenceDataAirlines:
		return new(ReferenceDataAirlinesRequest), new(ReferenceDataAirlinesResponse), nil
	case AirportPredictionsOnTime:
		return new(AirportPredictionsOnTimeRequest), new(AirportPredictionsOnTimeResponse), nil

	//
	// HOTELS
	//
	case ShoopingHotelsOffers:
		return new(ShoopingHotelOffersRequest), new(ShoopingHotelsOffersResponse), nil
	case ShoopingHotelOffers:
		return new(ShoopingHotelOffersRequest), new(ShoopingHotelOffersResponse), nil
	case EReputationHotelSentiments:
		return new(EReputationHotelSentimentsRequest), new(EReputationHotelSentimentsResponse), nil
	case BookingHotelBookings:
		return new(BookingHotelBookingsRequest), new(BookingHotelBookingsResponse), nil

	//
	// DESTINATIONS
	//
	case ReferenceDataLocationsPois:
		return new(ReferenceDataLocationsPoisRequest), new(ReferenceDataLocationsPoisResponse), nil
	case ReferenceDataLocationsPoi:
		return new(ReferenceDataLocationsPoisRequest), new(ReferenceDataLocationsPoiResponse), nil

	//
	// TRIP
	//
	case TravelPredictionTripPurpose:
		return new(TravelPredictionTripPurposeRequest), new(TravelPredictionTripPurposeResponse), nil
	case MediaFilesGeneratedPhotos:
		return new(MediaFilesGeneratedPhotosRequest), new(MediaFilesGeneratedPhotosResponse), nil

	default:
		return nil, nil, errors.New("request&response not recognized")
	}
}

// Do send request to api
func (a *Amadeus) Do(req Request, resp *Response, reqType string) error {

	// Check if token is expired
	if a.token.expired() {
		err := a.GetToken()
		if err != nil {
			return err
		}
	}

	// get base api url
	baseURL, err := a.getBaseURL()
	if err != nil {
		return err
	}

	// prepare request
	r, err := http.NewRequest(
		reqType,
		req.GetURL(baseURL, reqType),
		req.GetBody(reqType),
	)
	if err != nil {
		return err
	}

	// add headers
	r.Header.Add("Authorization", a.token.getAuthorization())
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")

	// send request
	client := http.Client{}
	rsp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	// read body
	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	fmt.Println("------------------")
	fmt.Println(rsp.StatusCode)
	fmt.Println(req.GetURL(baseURL, reqType))
	fmt.Println(string(b))
	fmt.Println("------------------")

	// decode response to struct
	err = (*resp).Decode(b)

	if err != nil {
		return err
	}

	return nil

}

// Generic structs

type ErrorResponse struct {
	Code   int    `json:"code,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
	Source struct {
		Pointer string `json:"pointer,omitempty"`
		Example string `json:"example,omitempty"`
	} `json:"source,omitempty"`
	Status int `json:"status,omitempty"`
}

type Warnings struct {
	Status int    `json:"status"`
	Code   int    `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Source Source `json:"source"`
}

type Source struct {
	Pointer   string `json:"pointer"`
	Parameter string `json:"parameter"`
	Example   string `json:"example"`
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
