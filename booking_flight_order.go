package amadeus

import "encoding/json"

// FlightCreateOrders

// REQUEST

type FlightCreateOrdersRequest struct {
	Data OrderData `json:"data,omitempty"`
}

// RESPONSE

type FlightCreateOrdersResponse struct {
	Data   OrderData       `json:"data,omitempty"`
	Errors []ErrorResponse `json:"errors,omitempty"`
}

func (a *Amadeus) FlightCreateOrder(request FlightCreateOrdersRequest) (FlightCreateOrdersResponse, error) {

	var response FlightCreateOrdersResponse

	urlStr, err := a.getURL(bookingFlightOrders)
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

func (a *Amadeus) FlightGetOrder(orderID string) (FlightCreateOrdersResponse, error) {

	var response FlightCreateOrdersResponse

	urlStr, err := a.getURL(bookingFlightOrders)
	if err != nil {
		return response, err
	}

	resp, err := a.getRequest(urlStr+"/"+orderID, []string{})
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
