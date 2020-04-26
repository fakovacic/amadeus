package examples

import (
	"fmt"
	"os"
	"testing"

	"github.com/fakovacic/amadeus-golang/amadeus"
	"github.com/joho/godotenv"
)

func TestAirportPredictionOnTime(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestAirportPredictionOnTime", func(t *testing.T) {

		// create client
		client, err := amadeus.New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		respData, err := client.Airport.Predictions.OnTime.Get(
			"airportCode=JFK",
			"date=2020-10-01",
		)
		if err != nil {
			t.Fatal("not expected error while getting response", err)
		}

		fmt.Println(respData)

	})

}
