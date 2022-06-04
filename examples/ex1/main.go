package main

import (
	"fmt"
	"os"
	"time"
        "encoding/json"
	"amadeus2"
)

func main() {
	client, err := amadeus.New(
		os.Getenv("AMADEUS_CLIENT_ID"),
		os.Getenv("AMADEUS_CLIENT_SECRET"),
		os.Getenv("AMADEUS_ENV"), // set to "TEST"
	)
	if err != nil {
		fmt.Println("not expected error while creating client", err)
	}

	// get offer flights request&response
	offerReq, offerResp, err := client.NewRequest(amadeus.ShoppingFlightOffers)

	// set offer flights params
	offerReq.(*amadeus.ShoppingFlightOffersRequest).SetCurrency("USD").SetSources("GDS").Return(
		"LAX",
		"NYC",
		time.Now().AddDate(0, 5, 0).Format("2006-01-02"),
		time.Now().AddDate(0, 7, 0).Format("2006-01-02"),
	).AddTravelers(1, 0, 0)

	// send request
	err = client.Do(offerReq, &offerResp, "GET")

	// get response
	offerRespData := offerResp.(*amadeus.ShoppingFlightOffersResponse)
        println(offerRespData)

}
