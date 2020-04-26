package examples

import (
	"fmt"
	"os"
	"testing"

	"github.com/fakovacic/amadeus-golang/amadeus"
	"github.com/joho/godotenv"
)

func TestLocations(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestLocations", func(t *testing.T) {

		// create client
		client, err := amadeus.New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		respData, err := client.ReferenceData.Locations.Get(
			"latitude=49.0000",
			"longitude=2.55",
		)

		respData, err = client.ReferenceData.Locations.Get(
			"keyword=MUC",
			"subType=AIRPORT,CITY",
		)

		respData, err = client.ReferenceData.Locations.Get(
			"locationId=CMUC",
		)

		if err != nil {
			t.Fatal("not expected error while getting response", err)
		}

		fmt.Println(respData)

	})

}
