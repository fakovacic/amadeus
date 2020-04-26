package examples

import (
	"fmt"
	"os"
	"testing"

	"github.com/fakovacic/amadeus-golang/amadeus"
	"github.com/joho/godotenv"
)

func TestNew(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	tt := []struct {
		Name     string
		Key      string
		Secret   string
		Env      string
		Expected string
	}{
		{
			Name:     "Blank",
			Key:      "",
			Secret:   "",
			Env:      "",
			Expected: "error",
		},
		{
			Name:     "Environment",
			Key:      os.Getenv("AMADEUS_CLIENT_ID"),
			Secret:   os.Getenv("AMADEUS_CLIENT_SECRET"),
			Env:      os.Getenv("AMADEUS_ENV"),
			Expected: "success",
		},
		{
			Name:     "Environment error",
			Key:      os.Getenv("AMADEUS_CLIENT_ID"),
			Secret:   os.Getenv("AMADEUS_CLIENT_SECRET"),
			Env:      "tester",
			Expected: "error",
		},
	}

	for _, tc := range tt {

		t.Run(tc.Name, func(t *testing.T) {

			client, err := amadeus.New(tc.Key, tc.Secret)
			if tc.Expected == "error" && err == nil {
				t.Fatal("expected error")
			}

			err = client.SetENV(tc.Env)
			if tc.Expected == "error" && err == nil {
				t.Fatal("expected error")
			}

			if tc.Expected == "success" && err != nil {
				t.Fatal("expected success", err)
			}
		})

	}

}
