package amadeus

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestFlightOrder(t *testing.T) {

	t.Run("TestFlightOrder", func(t *testing.T) {

		err := godotenv.Load()
		if err != nil {
			fmt.Println("Not found .env file")
		}

		amadeus, err := New(os.Getenv("AMADEUS_CLIENT_ID"), os.Getenv("AMADEUS_CLIENT_SECRET"), os.Getenv("AMADEUS_ENV"))
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		s := FlightOffersSearchRequest{
			CurrencyCode: "EUR",
			OriginDestinations: []OriginDestination{
				OriginDestination{
					ID:              "1",
					OriginCode:      "LON",
					DestinationCode: "PAR",
					DepartureDateTimeRange: TimeRange{
						Date: time.Now().AddDate(0, 5, 0).Format("2006-01-02"),
					},
				},
			},
			Travelers: []Travelers{
				Travelers{
					ID:           "1",
					TravelerType: "ADULT",
				},
			},
			Sources: []string{
				"GDS",
			},
		}

		offersResp, err := amadeus.FlightOffers(s)
		if err != nil {
			t.Fatal("not expected error", err)
		}

		// check error field
		if offersResp.Meta.Count == 0 {
			t.Fatal("return 0 results in offer search request")
		}

		p := FlightOffersPriceRequest{
			Data: PricingData{
				Type: "flight-offers-pricing",
				FlightOffers: []FlightOffer{
					offersResp.Data[0],
				},
			},
		}

		pricingResp, err := amadeus.FlightPricing(p)
		if err != nil {
			t.Fatal("not expected error", err)
		}

		// check error field
		if len(pricingResp.Data.FlightOffers) == 0 {
			t.Fatal("return 0 results in price request")
		}

		o := FlightCreateOrdersRequest{
			Data: OrderData{
				Type: "flight-order",
				FlightOffers: []FlightOffer{
					pricingResp.Data.FlightOffers[0],
				},
				Travelers: []Traveler{
					Traveler{
						ID:          "1",
						DateOfBirth: "1980-02-15",
						Name: Name{
							FirstName: "Foo",
							LastName:  "Bar",
						},
						Gender: "MALE",
						Contact: TravelerContact{
							EmailAddress: "foo@bar.com",
							Phones: []Phone{
								Phone{
									DeviceType:         "MOBILE",
									CountryCallingCode: "33",
									Number:             "480080072",
								},
							},
						},
					},
				},
				TicketingAgreement: TicketingAgreement{
					Option: "DELAY_TO_CANCEL",
					Delay:  "6D",
				},
				Contacts: []Contact{
					Contact{
						AddresseeName: AddresseeName{
							FirstName: "Foo",
							LastName:  "Bar",
						},
						CompanyName: "TESTING",
						Purpose:     "STANDARD",
						Phones: []Phone{
							Phone{
								DeviceType:         "MOBILE",
								CountryCallingCode: "33",
								Number:             "480080072",
							},
						},
						EmailAddress: "foo@bar.com",
						Address: Address{
							Lines: []string{
								"Street 25",
							},
							PostalCode:  "45453",
							CityName:    "Madrid",
							CountryCode: "ES",
						},
					},
				},
			},
		}

		orderResp, err := amadeus.FlightCreateOrder(o)
		if err != nil {
			t.Error("not expected error", err)
		}

		if len(orderResp.Data.AssociatedRecords) == 0 {
			t.Error("return 0 results in price request")
		}

		fmt.Println(orderResp)

	})

}

func TestFlightGetOrder(t *testing.T) {
	t.Fatal("TODO")
}
