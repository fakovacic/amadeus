package examples

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/fakovacic/amadeus-golang/amadeus"
	"github.com/fakovacic/amadeus-golang/amadeus/booking"
	"github.com/fakovacic/amadeus-golang/amadeus/shooping"
	"github.com/joho/godotenv"
)

func TestSeatmapsRequestFlightOffers(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestSeatmapsRequestFlightOffers", func(t *testing.T) {

		// create client
		client, err := amadeus.New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		// get request & response
		offerReq, offerResp, err := amadeus.NewRequest(amadeus.ShoppingFlightOffers)

		offerReq.(*shooping.FlightOffersRequest).
			SetCurrency("EUR").
			SetSources("GDS").
			Oneway(
				"LON",
				"PAR",
				time.Now().AddDate(0, 5, 0).Format("2006-01-02"),
			).
			AddTravelers(1, 0, 0)

		// send POST request
		err = client.Do(offerReq, &offerResp, "POST")

		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get response data
		offerRespData := offerResp.(*shooping.FlightOffersResponse)

		if len(offerRespData.Data) == 0 {
			t.Error("return 0 results in offer search request")
		}

		// get request & response
		seatMapsReq, seatMapsResp, err := amadeus.NewRequest(amadeus.ShoppingSeatmaps)

		seatMapsReq.(*shooping.SeatmapsRequest).AddOffer(
			offerRespData.GetOffer(10),
		)

		// send POST request
		err = client.Do(seatMapsReq, &seatMapsResp, "POST")

		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get response data
		seatMapsRespData := seatMapsResp.(*shooping.SeatmapsResponse)

		if len(seatMapsRespData.Data) == 0 {
			t.Error("return 0 results in seatMaps request")
		}

		fmt.Println("------------------")
		fmt.Println(seatMapsRespData)
		fmt.Println("------------------")

	})

}

func TestSeatmapsRequestFlightOrder(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestSeatmapsRequestFlightOrder", func(t *testing.T) {

		// create client
		client, err := amadeus.New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		// get offer flights request
		offerReq, offerResp, err := amadeus.NewRequest(amadeus.ShoppingFlightOffers)

		// set offer flights params
		offerReq.(*shooping.FlightOffersRequest).
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
		err = client.Do(offerReq, &offerResp, "POST")
		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get flight offers response
		offerRespData := offerResp.(*shooping.FlightOffersResponse)

		// check if data is returned
		if len(offerRespData.Data) == 0 {
			t.Error("return 0 results in offer search request")
		}

		// get pricing request
		pricingReq, pricingResp, err := amadeus.NewRequest(amadeus.ShoppingFlightPricing)

		// add offer from flight offers response
		pricingReq.(*shooping.FlightPricingRequest).AddOffer(
			offerRespData.GetOffer(1),
		)

		fmt.Println("------------------")
		fmt.Println(pricingReq)
		fmt.Println("------------------")

		// send request for flight pricing
		err = client.Do(pricingReq, &pricingResp, "POST")
		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get flight pricing response
		pricingRespData := pricingResp.(*shooping.FlightPricingResponse)

		// check if reponse exist
		if len(pricingRespData.Data.FlightOffers) == 0 {
			t.Error("return 0 results in offer search request")
		}

		// get booking request
		bookingReq, bookingResp, err := amadeus.NewRequest(amadeus.BookingFlightOrder)

		// add offer from flight offers response
		bookingReq.(*booking.FlightOrderRequest).AddOffers(
			pricingRespData.GetOffers(),
		).AddTicketingAgreement("DELAY_TO_CANCEL", "6D")

		// Add traveler
		bookingReq.(*booking.FlightOrderRequest).AddTraveler(
			bookingReq.(*booking.FlightOrderRequest).
				NewTraveler(
					"Foo", "Bar", "MALE", "1990-02-15",
				).
				AddEmail("foo@bar.com").
				AddMobile("33", "480080072"),
		)

		// Add contact
		bookingReq.(*booking.FlightOrderRequest).AddContact(
			bookingReq.(*booking.FlightOrderRequest).
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
		err = client.Do(bookingReq, &bookingResp, "POST")
		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get flight pricing response
		bookingRespData := bookingResp.(*booking.FlightOrderResponse)

		fmt.Println("------------------")
		fmt.Println(bookingRespData)
		fmt.Println("------------------")

		// check if reponse exist
		if len(bookingRespData.Data.FlightOffers) == 0 {
			t.Error("return 0 results in offer booking request")
		}

		// get request & response
		seatMapsReq, seatMapsResp, err := amadeus.NewRequest(amadeus.ShoppingSeatmaps)

		seatMapsReq.(*shooping.SeatmapsRequest).SetFlightOrderID(
			bookingRespData.Data.ID,
		)

		// send POST request
		err = client.Do(seatMapsReq, &seatMapsResp, "GET")

		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get response data
		seatMapsRespData := seatMapsResp.(*shooping.SeatmapsResponse)

		if len(seatMapsRespData.Data) == 0 {
			t.Error("return 0 results in seatMaps request")
		}

		fmt.Println("------------------")
		fmt.Println(seatMapsRespData)
		fmt.Println("------------------")

	})

}
