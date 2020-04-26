package examples

import (
	"fmt"
	"os"
	"testing"

	"github.com/fakovacic/amadeus-golang/amadeus"
	"github.com/fakovacic/amadeus-golang/amadeus/shooping"
	"github.com/joho/godotenv"
)

func TestShoopingHotelsOffers(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestShoopingHotelsOffers", func(t *testing.T) {

		// create client
		client, err := amadeus.New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		// get flights destinations
		req, resp, err := amadeus.NewRequest(amadeus.ShoopingHotelsOffers)

		// set flights destination request params
		req.(*shooping.HotelOffersRequest).SetCityCode("LON")

		// send request flight destinations
		err = client.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting destination data", err)
		}

		// get flight destination response
		respData := resp.(*shooping.HotelsOffersResponse)

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
		client, err := amadeus.New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		// get flights destinations
		req, resp, err := amadeus.NewRequest(amadeus.ShoopingHotelOffers)

		// set flights destination request params
		req.(*shooping.HotelOffersRequest).SetHotelID("HILONBE3")

		// send request flight destinations
		err = client.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting destination data", err)
		}

		// get flight destination response
		respData := resp.(*shooping.HotelOffersResponse)

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
		client, err := amadeus.New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		// get flights destinations
		req, resp, err := amadeus.NewRequest(amadeus.ShoopingHotelOffers)

		// set flights destination request params
		req.(*shooping.HotelOffersRequest).SetOfferID("3442F83BF3BF9482A7B058D67959FE807FB69B9344F85A4B3F1893DD903E1791")

		// send request flight destinations
		err = client.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting destination data", err)
		}

		// get flight destination response
		respData := resp.(*shooping.HotelOffersResponse)

		// check if reponse exist
		if len(respData.Data.Offers) == 0 {
			t.Error("return 0 results in offer search request")
		}

	})

}
