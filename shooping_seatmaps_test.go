package amadeus

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestShoppingSeatmapsRequestFlightOffers(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestShoppingSeatmapsRequestFlightOffers", func(t *testing.T) {

		// create client
		amadeus, err := New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
			os.Getenv("AMADEUS_ENV"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		// get request & response
		offerReq, offerResp, err := amadeus.NewRequest(ShoppingFlightOffers)

		offerReq.(*ShoppingFlightOffersRequest).
			SetCurrency("EUR").
			SetSources("GDS").
			Oneway(
				"LON",
				"PAR",
				time.Now().AddDate(0, 5, 0).Format("2006-01-02"),
			).
			AddTravelers(1, 0, 0)

		// send POST request
		err = amadeus.Do(offerReq, &offerResp, "POST")

		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get response data
		offerRespData := offerResp.(*ShoppingFlightOffersResponse)

		if len(offerRespData.Data) == 0 {
			t.Error("return 0 results in offer search request")
		}

		// get request & response
		seatMapsReq, seatMapsResp, err := amadeus.NewRequest(ShoppingSeatmaps)

		seatMapsReq.(*ShoppingSeatmapsRequest).AddOffer(
			offerRespData.GetOffer(10),
		)

		// send POST request
		err = amadeus.Do(seatMapsReq, &seatMapsResp, "POST")

		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get response data
		seatMapsRespData := seatMapsResp.(*ShoppingSeatmapsResponse)

		if len(seatMapsRespData.Data) == 0 {
			t.Error("return 0 results in seatMaps request")
		}

		fmt.Println("------------------")
		fmt.Println(seatMapsRespData)
		fmt.Println("------------------")

	})

}

func TestShoppingSeatmapsRequestFlightOrder(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestShoppingSeatmapsRequestFlightOrder", func(t *testing.T) {

		// create client
		amadeus, err := New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
			os.Getenv("AMADEUS_ENV"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		// get offer flights request
		offerReq, offerResp, err := amadeus.NewRequest(ShoppingFlightOffers)

		// set offer flights params
		offerReq.(*ShoppingFlightOffersRequest).
			SetCurrency("EUR").
			SetSources("GDS").
			Oneway(
				"FRA",
				"BER",
				time.Now().AddDate(0, 5, 0).Format("2006-01-02"),
			).
			AddTravelers(1, 0, 0)

		fmt.Println("------------------")
		fmt.Println(offerReq)
		fmt.Println("------------------")

		// send request flight offers
		err = amadeus.Do(offerReq, &offerResp, "POST")
		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get flight offers response
		offerRespData := offerResp.(*ShoppingFlightOffersResponse)

		// check if data is returned
		if len(offerRespData.Data) == 0 {
			t.Error("return 0 results in offer search request")
		}

		// get pricing request
		pricingReq, pricingResp, err := amadeus.NewRequest(ShoppingFlightPricing)

		// add offer from flight offers response
		pricingReq.(*ShoppingFlightPricingRequest).AddOffer(
			offerRespData.GetOffer(1),
		)

		fmt.Println("------------------")
		fmt.Println(pricingReq)
		fmt.Println("------------------")

		// send request for flight pricing
		err = amadeus.Do(pricingReq, &pricingResp, "POST")
		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get flight pricing response
		pricingRespData := pricingResp.(*ShoppingFlightPricingResponse)

		// check if reponse exist
		if len(pricingRespData.Data.FlightOffers) == 0 {
			t.Error("return 0 results in offer search request")
		}

		// get booking request
		bookingReq, bookingResp, err := amadeus.NewRequest(BookingFlightOrder)

		// add offer from flight offers response
		bookingReq.(*BookingFlightOrderRequest).AddOffers(
			pricingRespData.GetOffers(),
		).AddTicketingAgreement("DELAY_TO_CANCEL", "6D")

		// Add traveler
		bookingReq.(*BookingFlightOrderRequest).AddTraveler(
			bookingReq.(*BookingFlightOrderRequest).
				NewTraveler(
					"Foo", "Bar", "MALE", "1990-02-15",
				).
				AddEmail("foo@bar.com").
				AddMobile("33", "480080072"),
		)

		// Add contact
		bookingReq.(*BookingFlightOrderRequest).AddContact(
			bookingReq.(*BookingFlightOrderRequest).
				NewContact(
					"Foo", "Bar", "TESTING", "STANDARD",
				).
				AddEmail("foo@bar.com").
				AddMobile("33", "480080072").
				AddAddress("ES", "Madrid", "45453", "Street 25"),
		)

		fmt.Println("------------------")
		fmt.Println(bookingReq)
		fmt.Println("------------------")

		// send request for flight pricing
		err = amadeus.Do(bookingReq, &bookingResp, "POST")
		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get flight pricing response
		bookingRespData := bookingResp.(*BookingFlightOrderResponse)

		fmt.Println("------------------")
		fmt.Println(bookingRespData)
		fmt.Println("------------------")

		// check if reponse exist
		if len(bookingRespData.Data.FlightOffers) == 0 {
			t.Error("return 0 results in offer booking request")
		}

		// get request & response
		seatMapsReq, seatMapsResp, err := amadeus.NewRequest(ShoppingSeatmaps)

		seatMapsReq.(*ShoppingSeatmapsRequest).SetFlightOrderID(
			bookingRespData.Data.ID,
		)

		// send POST request
		err = amadeus.Do(seatMapsReq, &seatMapsResp, "GET")

		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get response data
		seatMapsRespData := seatMapsResp.(*ShoppingSeatmapsResponse)

		if len(seatMapsRespData.Data) == 0 {
			t.Error("return 0 results in seatMaps request")
		}

		fmt.Println("------------------")
		fmt.Println(seatMapsRespData)
		fmt.Println("------------------")

	})

}
