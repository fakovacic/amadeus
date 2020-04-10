package amadeus

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestAirportPredictionsOnTime(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestAirportPredictionsOnTime", func(t *testing.T) {

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
		req, resp, err := amadeus.NewRequest(AirportPredictionsOnTime)
		if err != nil {
			t.Fatal("not expected error while getting data", err)
		}

		// set Keyword
		req.(*AirportPredictionsOnTimeRequest).SetAirportCode("JKF").SetDate("2020-08-01")

		// send request
		err = amadeus.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting response data", err)
		}

		// get response
		respData := resp.(*AirportPredictionsOnTimeResponse)

		// check if reponse exist
		if respData.Data.ID == "" {
			t.Error("no results in response")
		}

	})

}
