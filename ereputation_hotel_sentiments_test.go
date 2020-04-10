package amadeus

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestEReputationHotelSentiments(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestEReputationHotelSentiments", func(t *testing.T) {

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
		req, resp, err := amadeus.NewRequest(EReputationHotelSentiments)
		if err != nil {
			t.Fatal("not expected error while getting data", err)
		}

		// set Keyword
		req.(*EReputationHotelSentimentsRequest).AddHotelIDs("ELONMFS", "ADNYCCTB", "XXXYYY01")

		// send request
		err = amadeus.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting location data", err)
		}

		// get response
		respData := resp.(*EReputationHotelSentimentsResponse)

		// check if reponse exist
		if len(respData.Data) == 0 {
			t.Error("return 0 results in request")
		}

	})

}
