package amadeus

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMediaFilesGeneratedPhotos(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestMediaFilesGeneratedPhotos", func(t *testing.T) {

		// create client
		amadeus, err := New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
			os.Getenv("AMADEUS_ENV"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		// get request&response
		req, resp, err := amadeus.NewRequest(MediaFilesGeneratedPhotos)
		if err != nil {
			t.Fatal("not expected error while getting data", err)
		}

		// set Keyword
		req.(*MediaFilesGeneratedPhotosRequest).SetCategory(MOUNTAIN)

		// send request
		err = amadeus.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting location data", err)
		}

		// get response
		respData := resp.(*MediaFilesGeneratedPhotosResponse)

		// check if reponse exist
		if respData.Data.AttachmentURI == "" {
			t.Error("return 0 results in request")
		}

	})

}
