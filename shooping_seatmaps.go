package amadeus

import (
	"encoding/json"
	"io"
	"strings"
)

type ShoppingSeatmapsRequest struct {
	FlightOrderID string        `json:"flightOrderID,omitempty"`
	FlightOffers  []FlightOffer `json:"data,omitempty"`
}

// SetFlightOrderID set flight order ID
func (sR *ShoppingSeatmapsRequest) SetFlightOrderID(id string) *ShoppingSeatmapsRequest {

	sR.FlightOrderID = id

	return sR
}

// AddOffer add flight offer to request
func (sR *ShoppingSeatmapsRequest) AddOffer(offer FlightOffer) *ShoppingSeatmapsRequest {

	sR.FlightOffers = append(sR.FlightOffers, offer)

	return sR
}

// GetURL returned key=value format for request on api
func (sR ShoppingSeatmapsRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := shoopingSeatmapsURL

	// add version
	switch reqType {
	case "GET":

		// define query params
		queryParams := []string{}

		url = baseURL + "/v1" + url

		if sR.FlightOrderID != "" {

			queryParams = append(queryParams, "flight-orderId="+sR.FlightOrderID)
			return url + "?" + strings.Join(queryParams, "&")

		}

		return ""

	case "POST":

		return baseURL + "/v1" + url
	}

	return ""
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (sR ShoppingSeatmapsRequest) GetBody(reqType string) io.Reader {

	switch reqType {
	case "GET":
		return nil
	case "POST":
		reqPayload, err := json.Marshal(sR)
		if err != nil {
			return nil
		}

		return strings.NewReader(string(reqPayload))
	}

	return nil
}

type ShoppingSeatmapsResponse struct {
	Meta   Meta            `json:"meta,omitempty"`
	Data   []SeatData      `json:"data,omitempty"`
	Errors []ErrorResponse `json:"errors,omitempty"`
}

// Decode implement Response interface
func (dR *ShoppingSeatmapsResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type SeatData struct {
	Type          string    `json:"type,omitempty"`
	FlightOfferid string    `json:"flightOfferid,omitempty"`
	Segmentid     string    `json:"segmentid,omitempty"`
	CarrierCode   string    `json:"carrierCode,omitempty"`
	Number        string    `json:"number,omitempty"`
	Class         string    `json:"class,omitempty"`
	Aircraft      Aircraft  `json:"aircraft,omitempty"`
	Departure     Departure `json:"departure,omitempty"`
	Arrival       Arrival   `json:"arrival,omitempty"`
	Decks         []Decks   `json:"decks,omitempty"`
}

type Departure struct {
	IataCode string `json:"iataCode,omitempty"`
	At       string `json:"at,omitempty"`
}
type Arrival struct {
	IataCode string `json:"iataCode"`
}

type Decks struct {
	DeckType          string            `json:"deckType,omitempty"`
	DeckConfiguration DeckConfiguration `json:"deckConfiguration,omitempty"`
	Facilities        []Facilities      `json:"facilities,omitempty"`
	Seats             []Seats           `json:"seats,omitempty"`
}

type DeckConfiguration struct {
	Width         int   `json:"width,omitempty"`
	Length        int   `json:"length,omitempty"`
	StartseatRow  int   `json:"startseatRow,omitempty"`
	EndSeatRow    int   `json:"endSeatRow,omitempty"`
	StartWingsRow int   `json:"startWingsRow,omitempty"`
	EndWingsRow   int   `json:"endWingsRow,omitempty"`
	StartWingsX   int   `json:"startWingsX,omitempty"`
	EndWingsX     int   `json:"endWingsX,omitempty"`
	ExitRowsX     []int `json:"exitRowsX,omitempty"`
}

type Facilities struct {
	Code        string      `json:"code,omitempty"`
	Column      string      `json:"column,omitempty"`
	Row         string      `json:"row,omitempty"`
	Position    string      `json:"position,omitempty"`
	Coordinates Coordinates `json:"coordinates,omitempty"`
}

type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type SeatPrice struct {
	Currency string `json:"currency"`
	Total    string `json:"total"`
}

type Seats struct {
	Cabin                string            `json:"cabin,omitempty"`
	Number               string            `json:"number,omitempty"`
	CharacteristicsCodes []string          `json:"characteristicsCodes,omitempty"`
	TravelerPricing      []TravelerPricing `json:"travelerPricing,omitempty"`
	Coordinates          Coordinates       `json:"coordinates,omitempty"`
}
