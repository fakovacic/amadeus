package examples

import (
	"fmt"
	"os"
	"testing"

	"github.com/fakovacic/amadeus-golang/amadeus"
	"github.com/joho/godotenv"
)

func TestFlightOffersRequestOneway(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestFlightOffersRequest", func(t *testing.T) {

		// create client
		client, err := amadeus.New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		respData, err := client.Shooping.FlightOffers.Get(
			"currency=EUR",
			"sources=GDS",
			"originLocationCode=LON",
			"destinationLocationCode=BCN",
			"departureDate=2020-10-10",
			"adults=1",
		)

		respData, err = client.Shooping.FlightOffers.Post(`{
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

		if err != nil {
			t.Fatal("not expected error while getting response", err)
		}

		fmt.Println(respData)

	})

}
