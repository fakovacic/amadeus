package amadeus

import (
	"github.com/fakovacic/amadeus-golang/amadeus/client"
	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/ereputation"
)

type EReputation struct {
	HotelSentiments
}

type HotelSentiments struct {
	Request  client.Request
	Response client.Response
}

// Get implement GET request
func (r *HotelSentiments) Get(params ...string) (client.Response, error) {

	// get request&response
	req, resp, err := client.NewRequest(client.EReputationHotelSentiments)
	if err != nil {
		return nil, err
	}

	// parse params
	req.(*ereputation.HotelSentimentsRequest).ParseParams(params)

	// send request
	err = client.Do(a, req, &resp, "GET")
	if err != nil {
		return nil, err
	}

	return resp.(*ereputation.HotelSentimentsResponse), nil
}

// Post implement POST request
func (r *HotelSentiments) Post(body string) (client.Response, error) {
	return nil, nil
}

// Delete implement DELETE request
func (r *HotelSentiments) Delete(params ...string) (client.Response, error) {
	return nil, nil
}
