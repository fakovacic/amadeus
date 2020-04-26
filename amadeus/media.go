package amadeus

import (
	"github.com/fakovacic/amadeus-golang/amadeus/client"
	files "github.com/fakovacic/amadeus-golang/amadeus/namespaces/media/files"
)

type Media struct {
	Files
}

type Files struct {
	GeneratedPhotos
}

type GeneratedPhotos struct {
	Request  client.Request
	Response client.Response
}

// Get implement GET request
func (r *GeneratedPhotos) Get(params ...string) (client.Response, error) {

	// get request&response
	req, resp, err := client.NewRequest(client.MediaFilesGeneratedPhotos)
	if err != nil {
		return nil, err
	}

	// parse params
	req.(*files.GeneratedPhotosRequest).ParseParams(params)

	// send request
	err = client.Do(a, req, &resp, "GET")
	if err != nil {
		return nil, err
	}

	return resp.(*files.GeneratedPhotosResponse), nil
}

// Post implement POST request
func (r *GeneratedPhotos) Post(body string) (client.Response, error) {
	return nil, nil
}

// Delete implement DELETE request
func (r *GeneratedPhotos) Delete(params ...string) (client.Response, error) {
	return nil, nil
}
