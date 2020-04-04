package amadeus

import "encoding/json"

// FlightOffersPrice

// REQUEST

type FlightOffersPriceRequest struct {
	Data PricingData `json:"data,omitempty"`
}

// RESPONSE

type FlightOffersPriceResponse struct {
	Data   PricingData     `json:"data,omitempty"`
	Errors []ErrorResponse `json:"errors,omitempty"`
}

type PricingData struct {
	Type         string        `json:"type,omitempty"`
	FlightOffers []FlightOffer `json:"flightOffers,omitempty"`
}

func (a *Amadeus) FlightPricing(request FlightOffersPriceRequest) (FlightOffersPriceResponse, error) {

	var response FlightOffersPriceResponse

	urlStr, err := a.getURL(shoopingFlightOffersPricing)
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
