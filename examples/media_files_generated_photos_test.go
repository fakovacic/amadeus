package examples

import (
	"fmt"
	"os"
	"testing"

	"github.com/fakovacic/amadeus-golang/amadeus"
	"github.com/joho/godotenv"
)

func TestMediaFilesGeneratedPhotos(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestMediaFilesGeneratedPhotos", func(t *testing.T) {

		// create client
		client, err := amadeus.New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		respData, err := client.Media.Files.GeneratedPhotos.Get(
			"category=MOUNTAIN",
		)

		if err != nil {
			t.Fatal("not expected error while getting data", err)
		}

		fmt.Println(respData)

	})

}
