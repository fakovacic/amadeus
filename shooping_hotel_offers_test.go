package amadeus

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestShoopingHotelsOffers(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestShoopingHotelsOffers", func(t *testing.T) {

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
		req, resp, err := amadeus.NewRequest(ShoopingHotelsOffers)

		// set flights destination request params
		req.(*ShoopingHotelOffersRequest).SetCityCode("LON")

		// send request flight destinations
		err = amadeus.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting destination data", err)
		}

		// get flight destination response
		respData := resp.(*ShoopingHotelsOffersResponse)

		// check if reponse exist
		if len(respData.Data) == 0 {
			t.Error("return 0 results in offer search request")
		}

	})

}

func TestShoopingHotelOffers(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestShoopingHotelOffers", func(t *testing.T) {

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
		req, resp, err := amadeus.NewRequest(ShoopingHotelOffers)

		// set flights destination request params
		req.(*ShoopingHotelOffersRequest).SetHotelID("HILONBE3")

		// send request flight destinations
		err = amadeus.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting destination data", err)
		}

		// get flight destination response
		respData := resp.(*ShoopingHotelOffersResponse)

		// check if reponse exist
		if len(respData.Data.Offers) == 0 {
			t.Error("return 0 results in offer search request")
		}

	})

}

func TestShoopingHotelOffer(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestShoopingHotelOffer", func(t *testing.T) {

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
		req, resp, err := amadeus.NewRequest(ShoopingHotelOffers)

		// set flights destination request params
		req.(*ShoopingHotelOffersRequest).SetOfferID("3442F83BF3BF9482A7B058D67959FE807FB69B9344F85A4B3F1893DD903E1791")

		// send request flight destinations
		err = amadeus.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting destination data", err)
		}

		// get flight destination response
		respData := resp.(*ShoopingHotelOffersResponse)

		// check if reponse exist
		if len(respData.Data.Offers) == 0 {
			t.Error("return 0 results in offer search request")
		}

	})

}
