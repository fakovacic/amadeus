package amadeus

import (
	"github.com/fakovacic/amadeus-golang/amadeus/client"
	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/airport/predictions"
)

type Airport struct {
	Predictions AirportPredictions
}

type AirportPredictions struct {
	OnTime
}

type OnTime struct {
	Request  client.Request
	Response client.Response
}

// Get implement GET request
func (r *OnTime) Get(params ...string) (client.Response, error) {

	// get request&response
	req, resp, err := client.NewRequest(client.AirportPredictionsOnTime)
	if err != nil {
		return nil, err
	}

	req.(*predictions.OnTimeRequest).ParseParams(params)

	// send request
	err = client.Do(a, req, &resp, "GET")
	if err != nil {
		return nil, err
	}

	// get response
	return resp.(*predictions.OnTimeResponse), nil
}

// Post implement POST request
func (r *OnTime) Post(body string) (client.Response, error) {
	return nil, nil
}

// Delete implement DELETE request
func (r *OnTime) Delete(params ...string) (client.Response, error) {
	return nil, nil
}
