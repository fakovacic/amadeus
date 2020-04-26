package examples

import (
	"fmt"
	"os"
	"testing"

	"github.com/fakovacic/amadeus-golang/amadeus"
	"github.com/joho/godotenv"
)

func TestLocationsPois(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestLocationsPois", func(t *testing.T) {

		// create client
		client, err := amadeus.New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		/*
				respData, err := client.ReferenceData.Locations.Pois.Get(
					"latitude=41.397158",
					"longitude=2.160873",
				)


			respData, err := client.ReferenceData.Locations.Pois.Get(
				"north=41.397158",
				"west=2.160873",
				"south=41.394582",
				"east=2.177181",
			)
		*/

		respData, err := client.ReferenceData.Locations.Pois.Get(
			"poisId=AB3F122E3E",
		)

		if err != nil {
			t.Fatal("not expected error while getting response", err)
		}

		fmt.Println(respData)

	})

}
