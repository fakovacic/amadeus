package amadeus

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestShoppingFlightOffersRequestOneway(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestShoppingFlightOffersRequestOneway", func(t *testing.T) {

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
		req, resp, err := amadeus.NewRequest(ShoppingFlightOffers)

		req.(*ShoppingFlightOffersRequest).
			SetCurrency("EUR").
			SetSources("GDS").
			Oneway(
				"LON",
				"PAR",
				time.Now().AddDate(0, 5, 0).Format("2006-01-02"),
			).
			AddTravelers(1, 0, 0)

		// send GET request
		err = amadeus.Do(req, &resp, "GET")

		// send POST request
		//err = amadeus.Do(req, &resp, "GET")

		if err != nil {
			t.Fatal("not expected error while geting flight offers data", err)
		}

		// get response data
		respData := resp.(*ShoppingFlightOffersResponse)

		if len(respData.Data) == 0 {
			t.Error("return 0 results in offer search request")
		}
	})

}
