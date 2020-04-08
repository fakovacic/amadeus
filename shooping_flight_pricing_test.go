package amadeus

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestShoppingFlightPricingRequest(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestShoppingFlightPricingRequest", func(t *testing.T) {

		// create client
		amadeus, err := New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
			os.Getenv("AMADEUS_ENV"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		// get offer flights request
		offerReq, offerResp, err := amadeus.NewRequest(ShoppingFlightOffers)

		// set offer flights params
		offerReq.(*ShoppingFlightOffersRequest).
			SetCurrency("EUR").
			SetSources("GDS").
			Oneway(
				"LON",
				"PAR",
				time.Now().AddDate(0, 5, 0).Format("2006-01-02"),
			).
			AddTravelers(1, 0, 0)

		// send request flight offers
		err = amadeus.Do(offerReq, &offerResp, "POST")
		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get flight offers response
		offerRespData := offerResp.(*ShoppingFlightOffersResponse)

		// check if data is returned
		if len(offerRespData.Data) == 0 {
			t.Fatal("return 0 results in offer search request")
		}

		// get pricing request
		pricingReq, pricingResp, err := amadeus.NewRequest(ShoppingFlightPricing)

		// add offer from flight offers response
		pricingReq.(*ShoppingFlightPricingRequest).AddOffer(
			offerRespData.GetOffer(1),
		)

		// send request for flight pricing
		err = amadeus.Do(pricingReq, &pricingResp, "POST")
		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get flight pricing response
		pricingRespData := pricingResp.(*ShoppingFlightPricingResponse)

		// check if reponse exist
		if len(pricingRespData.Data.FlightOffers) == 0 {
			t.Error("return 0 results in offer search request")
		}

	})

}
