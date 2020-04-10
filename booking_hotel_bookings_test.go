package amadeus

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestBookingHotelBookings(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestBookingHotelBookings", func(t *testing.T) {

		// create client
		amadeus, err := New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
			os.Getenv("AMADEUS_ENV"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		// SEARCH FOR HOTEL OFFERS

		// get request&response
		searchReq, searchResp, err := amadeus.NewRequest(ShoopingHotelsOffers)

		// set request params
		searchReq.(*ShoopingHotelOffersRequest).SetCityCode("LON")

		// send request
		err = amadeus.Do(searchReq, &searchResp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting destination data", err)
		}

		// get  response
		searchRespData := searchResp.(*ShoopingHotelsOffersResponse)

		// check if reponse exist
		if len(searchRespData.Data) == 0 {
			t.Error("return 0 results in offer search request")
		}

		//
		// CHECK HOTEL OFFER
		//

		// get request&response
		pricingReq, pricingResp, err := amadeus.NewRequest(ShoopingHotelOffers)

		// set request params
		pricingReq.(*ShoopingHotelOffersRequest).SetOfferID(
			searchRespData.GetHotel(3).GetOfferID(0),
		)

		// send request
		err = amadeus.Do(pricingReq, &pricingResp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting pricing data", err)
		}

		// get response
		pricingRespData := pricingResp.(*ShoopingHotelOffersResponse)

		// check if reponse exist
		if len(pricingRespData.Data.Offers) == 0 {
			t.Error("return 0 results in offer search request")
		}

		//
		// BOOKING
		//

		// get booking request
		bookingReq, bookingResp, err := amadeus.NewRequest(BookingHotelBookings)

		// add offer from flight offers response
		bookingReq.(*BookingHotelBookingsRequest).SetOfferID(
			pricingRespData.Data.GetOfferID(0),
		)

		// Add guest
		bookingReq.(*BookingHotelBookingsRequest).AddGuest(
			bookingReq.(*BookingHotelBookingsRequest).
				NewGuest("MR", "Foo", "Bar", "foo@bar.com", "+33679278416"),
		)

		// Add payment
		bookingReq.(*BookingHotelBookingsRequest).AddPayment(
			bookingReq.(*BookingHotelBookingsRequest).
				NewCard("VI", "4111111111111111", "2023-01"),
		)

		// send request
		err = amadeus.Do(bookingReq, &bookingResp, "POST")
		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get response
		bookingRespData := bookingResp.(*BookingHotelBookingsResponse)

		// check if reponse exist
		if len(bookingRespData.Data) == 0 {
			t.Error("return 0 results in offer booking request")
		}

	})

}
