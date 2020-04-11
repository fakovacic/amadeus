package amadeus

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestReferenceDataLocationsPois(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestReferenceDataLocationsPois", func(t *testing.T) {

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
		req, resp, err := amadeus.NewRequest(ReferenceDataLocationsPois)
		if err != nil {
			t.Fatal("not expected error while getting data", err)
		}

		// set Geo
		//req.(*ReferenceDataLocationsPoisRequest).SetGeo("41.397158", "2.160873")

		// set square
		req.(*ReferenceDataLocationsPoisRequest).SetSquare("41.397158", "2.160873", "41.394582", "2.177181")

		// send request
		err = amadeus.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting location data", err)
		}

		// get response
		respData := resp.(*ReferenceDataLocationsPoisResponse)

		// check if reponse exist
		if len(respData.Data) == 0 {
			t.Error("return 0 results in request")
		}

	})

}

func TestReferenceDataLocationsPoi(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestReferenceDataLocationsPoi", func(t *testing.T) {

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
		req, resp, err := amadeus.NewRequest(ReferenceDataLocationsPoi)
		if err != nil {
			t.Fatal("not expected error while getting data", err)
		}

		// set square
		req.(*ReferenceDataLocationsPoisRequest).SetPoisID("AB3F122E3E")

		// send request
		err = amadeus.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting location data", err)
		}

		// get response
		respData := resp.(*ReferenceDataLocationsPoiResponse)

		// check if reponse exist
		if respData.Data.ID == "" {
			t.Error("return 0 results in request")
		}

	})

}
