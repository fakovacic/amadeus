package amadeus

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestFlightOffers(t *testing.T) {

	t.Run("TestFlightOffers", func(t *testing.T) {

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

		resp, err := amadeus.FlightOffers(s)
		if err != nil {
			t.Error("not expected error", err)
		}

		if resp.Meta.Count == 0 {
			t.Error("return 0 results in offer search request")
		}

	})

}

func TestFlightOffersOneway(t *testing.T) {

	t.Run("TestFlightOffersOneway", func(t *testing.T) {

		err := godotenv.Load()
		if err != nil {
			fmt.Println("Not found .env file")
		}

		amadeus, err := New(os.Getenv("AMADEUS_CLIENT_ID"), os.Getenv("AMADEUS_CLIENT_SECRET"), os.Getenv("AMADEUS_ENV"))
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		s := NewSearchRequest("EUR", "GDS")
		s.Oneway("LON", "PAR", time.Now().AddDate(0, 5, 0).Format("2006-01-02"))
		s.AddTravelers(1, 0, 0)

		resp, err := amadeus.FlightOffers(s)
		if err != nil {
			t.Error("not expected error", err)
		}

		if resp.Meta.Count == 0 {
			t.Error("return 0 results in offer search request")
		}

	})

}
