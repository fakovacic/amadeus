package examples

import (
	"fmt"
	"os"
	"testing"

	"github.com/fakovacic/amadeus-golang/amadeus"
	"github.com/fakovacic/amadeus-golang/amadeus/booking"
	"github.com/fakovacic/amadeus-golang/amadeus/shooping"
	"github.com/joho/godotenv"
)

func TestBookingHotelBookings(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestBookingHotelBookings", func(t *testing.T) {

		// create client
		client, err := amadeus.New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		// SEARCH FOR HOTEL OFFERS

		// get request&response
		searchReq, searchResp, err := amadeus.NewRequest(amadeus.ShoopingHotelsOffers)

		// set request params
		searchReq.(*shooping.HotelOffersRequest).SetCityCode("LON")

		// send request
		err = client.Do(searchReq, &searchResp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting destination data", err)
		}

		// get  response
		searchRespData := searchResp.(*shooping.HotelsOffersResponse)

		// check if reponse exist
		if len(searchRespData.Data) == 0 {
			t.Error("return 0 results in offer search request")
		}

		//
		// CHECK HOTEL OFFER
		//

		// get request&response
		pricingReq, pricingResp, err := amadeus.NewRequest(amadeus.ShoopingHotelOffers)

		// set request params
		pricingReq.(*shooping.HotelOffersRequest).SetOfferID(
			searchRespData.GetHotel(3).GetOfferID(0),
		)

		// send request
		err = client.Do(pricingReq, &pricingResp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting pricing data", err)
		}

		// get response
		pricingRespData := pricingResp.(*shooping.HotelOffersResponse)

		// check if reponse exist
		if len(pricingRespData.Data.Offers) == 0 {
			t.Error("return 0 results in offer search request")
		}

		//
		// BOOKING
		//

		// get booking request
		bookingReq, bookingResp, err := amadeus.NewRequest(amadeus.BookingHotelBookings)

		// add offer from flight offers response
		bookingReq.(*booking.HotelBookingsRequest).SetOfferID(
			pricingRespData.Data.GetOfferID(0),
		)

		// Add guest
		bookingReq.(*booking.HotelBookingsRequest).AddGuest(
			bookingReq.(*booking.HotelBookingsRequest).
				NewGuest("MR", "Foo", "Bar", "foo@bar.com", "+33679278416"),
		)

		// Add payment
		bookingReq.(*booking.HotelBookingsRequest).AddPayment(
			bookingReq.(*booking.HotelBookingsRequest).
				NewCard("VI", "4111111111111111", "2023-01"),
		)

		// send request
		err = client.Do(bookingReq, &bookingResp, "POST")
		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get response
		bookingRespData := bookingResp.(*booking.HotelBookingsResponse)

		// check if reponse exist
		if len(bookingRespData.Data) == 0 {
			t.Error("return 0 results in offer booking request")
		}

	})

}
