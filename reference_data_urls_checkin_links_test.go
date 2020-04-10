package amadeus

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestReferenceDataUrlsCheckinLinks(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestReferenceDataUrlsCheckinLinks", func(t *testing.T) {

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
		req, resp, err := amadeus.NewRequest(ReferenceDataUrlsCheckinLinks)
		if err != nil {
			t.Fatal("not expected error while getting data", err)
		}

		// set Keyword
		req.(*ReferenceDataUrlsCheckinLinksRequest).SetAirlineCode("KL")

		// send request
		err = amadeus.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting location data", err)
		}

		// get response
		respData := resp.(*ReferenceDataUrlsCheckinLinksResponse)

		// check if reponse exist
		if len(respData.Data) == 0 {
			t.Error("return 0 results in request")
		}

	})

}
