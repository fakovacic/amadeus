package amadeus

import (
	"encoding/json"
	"io"
	"strings"
        // "fmt"
)

// ShoppingFlightPricingRequest

// REQUEST

type ShoppingFlightPricingRequest struct {
	Data PricingData `json:"data,omitempty"`
}

// AddOffer add flight offer to request
func (sR *ShoppingFlightPricingRequest) AddOffer(offer FlightOffer) *ShoppingFlightPricingRequest {

	sR.Data.Type = "flight-offers-pricing"
	sR.Data.FlightOffers = append(sR.Data.FlightOffers, offer)

	return sR
}

// GetURL returned key=value format for request on api
func (sR ShoppingFlightPricingRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := shoopingFlightOffersPricingURL

	// add version
	switch reqType {
	case "POST":

		return baseURL + "/v1" + url
	}

	return ""
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (sR ShoppingFlightPricingRequest) GetBody(reqType string) io.Reader {

	switch reqType {
	case "POST":
		reqPayload, err := json.Marshal(sR)
                // fmt.Println(strings.NewReader(string(reqPayload)))
		if err != nil {
			return nil
		}

		return strings.NewReader(string(reqPayload))
	}

	return nil
}

// RESPONSE

type ShoppingFlightPricingResponse struct {
	Data         PricingData     `json:"data,omitempty"`
	Dictionaries Dictionaries    `json:"dictionaries,omitempty"`
	Errors       []ErrorResponse `json:"errors,omitempty"`
}

type PricingData struct {
	Type         string        `json:"type,omitempty"`
	FlightOffers []FlightOffer `json:"flightOffers,omitempty"`
}

// Decode implement Response interface
func (dR *ShoppingFlightPricingResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

// GetOffer return offer from list
func (dR ShoppingFlightPricingResponse) GetOffers() []FlightOffer {
	return dR.Data.FlightOffers
}
