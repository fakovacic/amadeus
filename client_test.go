package amadeus

import (
	"fmt"
	"os"
	"testing"

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
			Env:      "TESTER",
			Expected: "error",
		},
	}

	for _, tc := range tt {

		t.Run(tc.Name, func(t *testing.T) {

			_, err := New(tc.Key, tc.Secret, tc.Env)
			if tc.Expected == "error" && err == nil {
				t.Fatal("expected error")
			}

			if tc.Expected == "success" && err != nil {
				t.Fatal("expected success")
			}

		})

	}

}

func TestToken(t *testing.T) {

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
			Name:     "Wrong credentials",
			Key:      "This is key",
			Secret:   "My secret",
			Env:      "TEST",
			Expected: "error",
		},
		{
			Name:     "Environment",
			Key:      os.Getenv("AMADEUS_CLIENT_ID"),
			Secret:   os.Getenv("AMADEUS_CLIENT_SECRET"),
			Env:      os.Getenv("AMADEUS_ENV"),
			Expected: "success",
		},
	}

	for _, tc := range tt {

		t.Run(tc.Name, func(t *testing.T) {

			amadeus, err := New(tc.Key, tc.Secret, tc.Env)
			if err != nil {
				t.Fatal("not expected error", err)
			}

			err = amadeus.GetToken()
			if tc.Expected == "error" && err == nil {
				t.Fatal("expected error")
			}

			if tc.Expected == "success" && err != nil {
				t.Fatal("expected success", err)
			}

		})

	}

}

func TestCheckToken(t *testing.T) {
	t.Fatal("TODO")
}

func TestTokenExipry(t *testing.T) {
	t.Fatal("TODO")
}

func TestTokenGetAuthorization(t *testing.T) {
	t.Fatal("TODO")
}
