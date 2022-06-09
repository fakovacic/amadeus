package main

import (
	"fmt"
	"os"
	"time"
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
        // .SetCarrer
	offerReq.(*amadeus.ShoppingFlightOffersRequest).SetCurrency("USD").SetSources("GDS").Return(
		"AUS",
		"DFW",
		time.Now().AddDate(0, 1, 0).Format("2006-01-02"),
		time.Now().AddDate(0, 2, 0).Format("2006-01-02"),
	).AddTravelers(1, 0, 0)

	// send request
	err = client.Do(offerReq, &offerResp, "GET")

	// get response
	offerRespData := offerResp.(*amadeus.ShoppingFlightOffersResponse)

        // get pricing request&response
        pricingReq, pricingResp, err := client.NewRequest(amadeus.ShoppingFlightPricing)

        // add offer from flight offers response
        // this should really be a parsed as json
        // and offered as a random id 
        pricingReq.(*amadeus.ShoppingFlightPricingRequest).AddOffer(
             // use id 250 to avoid "No fare applicable" error
             offerRespData.GetOffer(3),
        )

        err = client.Do(pricingReq, &pricingResp, "POST")

        // get response
        pricingRespData := pricingResp.(*amadeus.ShoppingFlightPricingResponse)

        // get booking request
        bookingReq, bookingResp, err := client.NewRequest(amadeus.BookingFlightOrder)

        // add offer from flight pricing response
        bookingReq.(*amadeus.BookingFlightOrderRequest).AddOffers(
          pricingRespData.GetOffers(),
        ).AddTicketingAgreement("DELAY_TO_CANCEL", "6D")

        // add payment
        bookingReq.(*amadeus.BookingFlightOrderRequest).AddPayment(
          bookingReq.(*amadeus.BookingFlightOrderRequest).
            NewCard("VI", "4111111111111111", "2023-01"),
        )

        // add traveler
        bookingReq.(*amadeus.BookingFlightOrderRequest).AddTraveler(
              bookingReq.(*amadeus.BookingFlightOrderRequest).
                 NewTraveler(
                      "Foo", "Bar", "MALE", "1990-02-15",
                 ).
                  AddEmail("foo@bar.com").
                  AddMobile("33", "5550800072"),
        )

        // send request
        err = client.Do(bookingReq, &bookingResp, "POST")

        // get flight booking response
        bookingRespData := bookingResp.(*amadeus.BookingFlightOrderResponse)
        fmt.Println(bookingRespData)
}
