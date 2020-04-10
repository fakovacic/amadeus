package amadeus

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestTravelPredictionTripPurpose(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestTravelPredictionTripPurpose", func(t *testing.T) {

		// create client
		amadeus, err := New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
			os.Getenv("AMADEUS_ENV"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		// get request&response
		req, resp, err := amadeus.NewRequest(TravelPredictionTripPurpose)

		// set return flight params for prediction
		req.(*TravelPredictionTripPurposeRequest).
			ReturnFlight("BCN", "MUC", "2020-10-10", "2020-10-20")

		// send request
		err = amadeus.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting destination data", err)
		}

		// get response
		respData := resp.(*TravelPredictionTripPurposeResponse)

		// check if reponse exist
		if respData.Data.ID == "" {
			t.Error("return 0 results in offer search request")
		}

	})

}
