package amadeus

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestShoppingFlightDestination(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestShoppingFlightDestination", func(t *testing.T) {

		// create client
		amadeus, err := New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
			os.Getenv("AMADEUS_ENV"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		// get flights destinations
		req, resp, err := amadeus.NewRequest(ShoppingFlightDestination)

		// set flights destination request params
		req.(*ShoppingFlightDestinationRequest).SetOrigin("LON")

		// send request flight destinations
		err = amadeus.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting destination data", err)
		}

		// get flight destination response
		respData := resp.(*ShoppingFlightDestinationResponse)

		// check if reponse exist
		if len(respData.Data) == 0 {
			t.Error("return 0 results in offer search request")
		}

	})

}
