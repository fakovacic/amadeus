package examples

import (
	"fmt"
	"os"
	"testing"

	"github.com/fakovacic/amadeus-golang/amadeus"
	"github.com/joho/godotenv"
)

func TestFlightPricingRequest(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestFlightPricingRequest", func(t *testing.T) {

		// create client
		client, err := amadeus.New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		offerRespData, err := client.Shooping.FlightOffers.Get(
			"currency=EUR",
			"sources=GDS",
			"originLocationCode=LON",
			"destinationLocationCode=BCN",
			"departureDate=2020-10-10",
			"adults=1",
		)

		/*
			offerRespData, err = client.Shooping.FlightOffers.Post(`{
				"currencyCode": "USD",
				"originDestinations": [
				  {
					"id": "1",
					"originLocationCode": "LON",
					"destinationLocationCode": "BCN",
					"departureDateTimeRange": {
					  "date": "2020-10-10"
					}
				  }
				],
				"travelers": [
				  {
					"id": "1",
					"travelerType": "ADULT"
				  }
				],
				"sources": [
				  "GDS"
				]
			  }`)
		*/

		if err != nil {
			t.Fatal("not expected error while getting response", err)
		}

		fmt.Println(offerRespData.GetFlightOffer(1).Body())

		pricingRespData, err := client.Shooping.FlightPricing.Post(
			offerRespData.GetFlightOffer(1).Body(),
			//include=credit-card-fees,bags,other-services,detailed-fare-rules
			//forceClass
		)

		if err != nil {
			t.Fatal("not expected error while getting response", err)
		}

		fmt.Println(pricingRespData)

	})

}
