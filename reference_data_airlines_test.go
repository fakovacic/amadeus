package amadeus

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestReferenceDataAirlines(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestReferenceDataAirlines", func(t *testing.T) {

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
		req, resp, err := amadeus.NewRequest(ReferenceDataAirlines)
		if err != nil {
			t.Fatal("not expected error while getting data", err)
		}

		// set Keyword
		req.(*ReferenceDataAirlinesRequest).AddAirlineCode("KL")

		// send request
		err = amadeus.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting location data", err)
		}

		// get response
		respData := resp.(*ReferenceDataAirlinesResponse)

		// check if reponse exist
		if len(respData.Data) == 0 {
			t.Error("return 0 results in request")
		}

	})

}
