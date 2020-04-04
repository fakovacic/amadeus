package amadeus

import (
	"os"
	"testing"
	"time"
)

func TestFlightPricing(t *testing.T) {

	t.Run("TestFlightPricing", func(t *testing.T) {

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
			t.Error("not expected error", err)
		}

		if len(pricingResp.Data.FlightOffers) == 0 {
			t.Error("return 0 results in price request")
		}

	})

}
