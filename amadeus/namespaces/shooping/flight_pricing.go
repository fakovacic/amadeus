package shooping

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/structs"
)

const (

	///////////
	// 	AIR	 //
	///////////

	// Shooping Flight offers pricing
	// check certain offer if is still active, response with additional data for offer
	flightOffersPricingURL = "/shopping/flight-offers/pricing"
)

// FlightPricingRequest

// REQUEST

type FlightPricingRequest struct {
	Data PricingData `json:"data,omitempty"`
}

// AddOffer add flight offer to request
func (sR *FlightPricingRequest) AddOffer(offer structs.FlightOffer) *FlightPricingRequest {

	sR.Data.Type = "flight-offers-pricing"
	sR.Data.FlightOffers = append(sR.Data.FlightOffers, offer)

	return sR
}

// SetParam set params
func (dR *FlightPricingRequest) SetParam(key, value string) {
	return
}

// ParseBody parse JSON body to request
func (sR *FlightPricingRequest) ParseBody(body string) *FlightPricingRequest {

	var flightOffers []structs.FlightOffer

	err := json.Unmarshal([]byte(body), &flightOffers)
	if err != nil {

		var flightOffer structs.FlightOffer

		err := json.Unmarshal([]byte(body), &flightOffer)
		if err != nil {
			return nil
		}

		sR.AddOffer(flightOffer)
		return sR
	}

	if len(flightOffers) != 0 {
		for _, off := range flightOffers {

			fmt.Println(off)

			sR.AddOffer(off)
		}

	}

	return sR
}

// GetURL returned key=value format for request on api
func (sR FlightPricingRequest) GetURL(reqType string) string {

	// set request url
	url := flightOffersPricingURL

	// add version
	switch reqType {
	case "POST":

		return "/v1" + url
	}

	return ""
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (sR FlightPricingRequest) GetBody(reqType string) io.Reader {

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

type FlightPricingResponse struct {
	Data         PricingData          `json:"data,omitempty"`
	Dictionaries structs.Dictionaries `json:"dictionaries,omitempty"`
}

type PricingData struct {
	Type         string                `json:"type,omitempty"`
	FlightOffers []structs.FlightOffer `json:"flightOffers,omitempty"`
}

// Decode implement Response interface
func (dR *FlightPricingResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

// GetOffer return offer from list
func (dR FlightPricingResponse) GetOffers() []structs.FlightOffer {
	return dR.Data.FlightOffers
}
