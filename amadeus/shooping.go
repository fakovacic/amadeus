package amadeus

import (
	"github.com/fakovacic/amadeus-golang/amadeus/client"
	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/shooping"
)

type Shooping struct {
	FlightDestination
	FlightDates
	FlightOffers
	FlightPricing
}

type FlightDestination struct {
	Request  client.Request
	Response client.Response
}

// Get implement GET request
func (r *FlightDestination) Get(params ...string) (client.Response, error) {

	// get request&response
	req, resp, err := client.NewRequest(client.ShoppingFlightDestination)
	if err != nil {
		return nil, err
	}

	// parse params
	req.(*shooping.FlightDestinationRequest).ParseParams(params)

	// send request
	err = client.Do(a, req, &resp, "GET")
	if err != nil {
		return nil, err
	}

	return resp.(*shooping.FlightDestinationResponse), nil
}

// Post implement POST request
func (r *FlightDestination) Post(body string) (client.Response, error) {
	return nil, nil
}

// Delete implement DELETE request
func (r *FlightDestination) Delete(params ...string) (client.Response, error) {
	return nil, nil
}

type FlightDates struct {
	Request  client.Request
	Response client.Response
}

// Get implement GET request
func (r *FlightDates) Get(params ...string) (client.Response, error) {

	// get request&response
	req, resp, err := client.NewRequest(client.ShoppingFlightDates)
	if err != nil {
		return nil, err
	}

	// parse params
	req.(*shooping.FlightDatesRequest).ParseParams(params)

	// send request
	err = client.Do(a, req, &resp, "GET")
	if err != nil {
		return nil, err
	}

	return resp.(*shooping.FlightDatesResponse), nil
}

// Post implement POST request
func (r *FlightDates) Post(body string) (client.Response, error) {
	return nil, nil
}

// Delete implement DELETE request
func (r *FlightDates) Delete(params ...string) (client.Response, error) {
	return nil, nil
}

type FlightOffers struct {
	Request  client.Request
	Response client.Response
}

// Get implement GET request
func (r *FlightOffers) Get(params ...string) (*shooping.FlightOffersResponse, error) {

	// get request&response
	req, resp, err := client.NewRequest(client.ShoppingFlightOffers)
	if err != nil {
		return nil, err
	}

	// parse params
	req.(*shooping.FlightOffersRequest).ParseParams(params)

	// send request
	err = client.Do(a, req, &resp, "GET")
	if err != nil {
		return nil, err
	}

	return resp.(*shooping.FlightOffersResponse), nil
}

// Post implement POST request
func (r *FlightOffers) Post(body string) (*shooping.FlightOffersResponse, error) {

	// get request&response
	req, resp, err := client.NewRequest(client.ShoppingFlightOffers)
	if err != nil {
		return nil, err
	}

	// parse params
	req.(*shooping.FlightOffersRequest).ParseBody(body)

	// send request
	err = client.Do(a, req, &resp, "POST")
	if err != nil {
		return nil, err
	}

	return resp.(*shooping.FlightOffersResponse), nil
}

// Delete implement DELETE request
func (r *FlightOffers) Delete(params ...string) (client.Response, error) {
	return nil, nil
}

type FlightPricing struct {
	Request  client.Request
	Response client.Response
}

// Get implement GET request
func (r *FlightPricing) Get(params ...string) (*shooping.FlightOffersResponse, error) {
	return nil, nil
}

// Post implement POST request
func (r *FlightPricing) Post(body string) (*shooping.FlightPricingResponse, error) {

	// get request&response
	req, resp, err := client.NewRequest(client.ShoppingFlightPricing)
	if err != nil {
		return nil, err
	}

	// parse params
	req.(*shooping.FlightPricingRequest).ParseBody(body)

	// send request
	err = client.Do(a, req, &resp, "POST")
	if err != nil {
		return nil, err
	}

	return resp.(*shooping.FlightPricingResponse), nil
}

// Delete implement DELETE request
func (r *FlightPricing) Delete(params ...string) (client.Response, error) {
	return nil, nil
}
