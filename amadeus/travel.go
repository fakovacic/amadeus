package amadeus

import (
	"github.com/fakovacic/amadeus-golang/amadeus/client"
	analitycs "github.com/fakovacic/amadeus-golang/amadeus/namespaces/travel/analytics"
	predictions "github.com/fakovacic/amadeus-golang/amadeus/namespaces/travel/predictions"
)

type Travel struct {
	Analytics
	Predictions TravelPredictions
}

type Analytics struct {
	AirTraffic
}

type AirTraffic struct {
	Traveled
	Booked
	Busiest
}

type Traveled struct {
	Request  client.Request
	Response client.Response
}

// Get implement GET request
func (r *Traveled) Get(params ...string) (client.Response, error) {

	// get request&response
	req, resp, err := client.NewRequest(client.TravelAnalyticsAirTraffic)
	if err != nil {
		return nil, err
	}

	req.(*analitycs.AirTrafficRequest).SetType("TRAVELED").ParseParams(params)

	// send request
	err = client.Do(a, req, &resp, "GET")
	if err != nil {
		return nil, err
	}

	return resp.(*analitycs.AirTrafficResponse), nil
}

// Post implement POST request
func (r *Traveled) Post(body string) (client.Response, error) {
	return nil, nil
}

// Delete implement DELETE request
func (r *Traveled) Delete(params ...string) (client.Response, error) {
	return nil, nil
}

type Booked struct {
	Request  client.Request
	Response client.Response
}

// Get implement GET request
func (r *Booked) Get(params ...string) (client.Response, error) {

	// get request&response
	req, resp, err := client.NewRequest(client.TravelAnalyticsAirTraffic)
	if err != nil {
		return nil, err
	}

	req.(*analitycs.AirTrafficRequest).SetType("BOOKED").ParseParams(params)

	// send request
	err = client.Do(a, req, &resp, "GET")
	if err != nil {
		return nil, err
	}

	return resp.(*analitycs.AirTrafficResponse), nil
}

// Post implement POST request
func (r *Booked) Post(body string) (client.Response, error) {
	return nil, nil
}

// Delete implement DELETE request
func (r *Booked) Delete(params ...string) (client.Response, error) {
	return nil, nil
}

type Busiest struct {
	Request  client.Request
	Response client.Response
}

// Get implement GET request
func (r *Busiest) Get(params ...string) (client.Response, error) {

	// get request&response
	req, resp, err := client.NewRequest(client.TravelAnalyticsAirTraffic)
	if err != nil {
		return nil, err
	}

	req.(*analitycs.AirTrafficRequest).SetType("BUSIEST").ParseParams(params)

	// send request
	err = client.Do(a, req, &resp, "GET")
	if err != nil {
		return nil, err
	}

	return resp.(*analitycs.AirTrafficResponse), nil
}

// Post implement POST request
func (r *Busiest) Post(body string) (client.Response, error) {
	return nil, nil
}

// Delete implement DELETE request
func (r *Busiest) Delete(params ...string) (client.Response, error) {
	return nil, nil
}

type TravelPredictions struct {
	TripPurpose
}

type TripPurpose struct {
	Request  client.Request
	Response client.Response
}

// Get implement GET request
func (r *TripPurpose) Get(params ...string) (client.Response, error) {

	// get request&response
	req, resp, err := client.NewRequest(client.TravelPredictionTripPurpose)
	if err != nil {
		return nil, err
	}

	req.(*predictions.TripPurposeRequest).ParseParams(params)

	// send request
	err = client.Do(a, req, &resp, "GET")
	if err != nil {
		return nil, err
	}

	return resp.(*predictions.TripPurposeResponse), nil
}

// Post implement POST request
func (r *TripPurpose) Post(body string) (client.Response, error) {
	return nil, nil
}

// Delete implement DELETE request
func (r *TripPurpose) Delete(params ...string) (client.Response, error) {
	return nil, nil
}
