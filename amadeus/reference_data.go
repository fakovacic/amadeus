package amadeus

import (
	"github.com/fakovacic/amadeus-golang/amadeus/client"
	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/referencedata"
	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/referencedata/locations"
	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/referencedata/urls"
)

type ReferenceData struct {
	Airlines
	Urls
	Locations
}

type Airlines struct {
	Request  client.Request
	Response client.Response
}

// Get implement GET request
func (r *Airlines) Get(params ...string) (client.Response, error) {

	// get request&response
	req, resp, err := client.NewRequest(client.ReferenceDataAirlines)
	if err != nil {
		return nil, err
	}

	// set params
	req.(*referencedata.AirlinesRequest).ParseParams(params)

	// send request
	err = client.Do(a, req, &resp, "GET")
	if err != nil {
		return nil, err
	}

	// get response
	return resp.(*referencedata.AirlinesResponse), nil
}

// Post implement POST request
func (r *Airlines) Post(body string) (client.Response, error) {
	return nil, nil
}

// Delete implement DELETE request
func (r *Airlines) Delete(params ...string) (client.Response, error) {
	return nil, nil
}

type Urls struct {
	CheckInLinks
}

type CheckInLinks struct {
	Request  client.Request
	Response client.Response
}

// Get implement GET request
func (r *CheckInLinks) Get(params ...string) (client.Response, error) {

	// get request&response
	req, resp, err := client.NewRequest(client.ReferenceDataUrlsCheckInLinks)
	if err != nil {
		return nil, err
	}

	// set Keyword
	req.(*urls.CheckInLinksRequest).ParseParams(params)

	// send request
	err = client.Do(a, req, &resp, "GET")
	if err != nil {
		return nil, err
	}

	// get response
	return resp.(*urls.CheckInLinksResponse), nil
}

// Post implement POST request
func (r *CheckInLinks) Post(body string) (client.Response, error) {
	return nil, nil
}

// Delete implement DELETE request
func (r *CheckInLinks) Delete(params ...string) (client.Response, error) {
	return nil, nil
}

type Locations struct {
	Request  client.Request
	Response client.Response
	Pois
}

// Get implement GET request
func (r *Locations) Get(params ...string) (client.Response, error) {

	// get request&response
	req, resp, err := client.NewRequest(client.ReferenceDataLocations)
	if err != nil {
		return nil, err
	}

	// set Keyword
	req.(*referencedata.LocationsRequest).ParseParams(params)

	// check if request contain LocationID
	// response is only one Data location, than response is different object
	if req.(*referencedata.LocationsRequest).LocationID != "" {
		// get request&response
		_, resp, err = client.NewRequest(client.ReferenceDataLocation)
		if err != nil {
			return nil, err
		}

		// send request
		err = client.Do(a, req, &resp, "GET")
		if err != nil {
			return nil, err
		}

		// get response
		return resp.(*referencedata.LocationResponse), nil

	}

	// send request
	err = client.Do(a, req, &resp, "GET")
	if err != nil {
		return nil, err
	}

	// get response
	return resp.(*referencedata.LocationsResponse), nil
}

// Post implement POST request
func (r *Locations) Post(body string) (client.Response, error) {
	return nil, nil
}

// Delete implement DELETE request
func (r *Locations) Delete(params ...string) (client.Response, error) {
	return nil, nil
}

type Pois struct {
	Request  client.Request
	Response client.Response
}

// Get implement GET request
func (r *Pois) Get(params ...string) (client.Response, error) {

	// get request&response
	req, resp, err := client.NewRequest(client.ReferenceDataLocationsPois)
	if err != nil {
		return nil, err
	}

	req.(*locations.PoisRequest).ParseParams(params)

	if req.(*locations.PoisRequest).PoisID != "" {

		// get request&response
		_, resp, err := client.NewRequest(client.ReferenceDataLocationsPoi)
		if err != nil {
			return nil, err
		}

		// send request
		err = client.Do(a, req, &resp, "GET")
		if err != nil {
			return nil, err
		}

		return resp.(*locations.PoiResponse), nil

	}

	// send request
	err = client.Do(a, req, &resp, "GET")
	if err != nil {
		return nil, err
	}

	return resp.(*locations.PoisResponse), nil
}

// Post implement POST request
func (r *Pois) Post(body string) (client.Response, error) {
	return nil, nil
}

// Delete implement DELETE request
func (r *Pois) Delete(params ...string) (client.Response, error) {
	return nil, nil
}
