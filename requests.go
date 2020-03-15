package amadeusgolang

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func (a *Amadeus) request(reqPayload, url string) ([]byte, error) {

	if a.Token.expired() {
		err := a.getToken()
		if err != nil {
			return nil, err
		}
	}

	payload := strings.NewReader(reqPayload)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", a.Token.getBearer())
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (a *Amadeus) FlightOffers(request FlightOffersSearchRequest) (FlightOffersSearchResponse, error) {

	var response FlightOffersSearchResponse

	urlStr := a.getURL(shoopingFlightOffers)

	reqPayload, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	resp, err := a.request(string(reqPayload), urlStr)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (a *Amadeus) FlightPricing(request FlightOffersPriceRequest) (FlightOffersPriceResponse, error) {

	var response FlightOffersPriceResponse

	urlStr := a.getURL(shoopingFlightOffersPricing)

	reqPayload, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	resp, err := a.request(string(reqPayload), urlStr)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (a *Amadeus) FlightCreateOrder(request FlightCreateOrdersRequest) (FlightCreateOrdersResponse, error) {

	var response FlightCreateOrdersResponse

	urlStr := a.getURL(bookingFlightOrders)

	reqPayload, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	resp, err := a.request(string(reqPayload), urlStr)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
