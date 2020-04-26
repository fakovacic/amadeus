package media

import (
	"encoding/json"
	"io"
	"strings"
)

const (

	////////////
	// 	TRIP  //
	////////////

	// AI Generated Photos
	// The AI-Generated Photos API returns a link to download a rendered image of a landscape.
	generatedPhotosURL = "/media/files/generated-photos"
)

type Category string

func (l Category) String() string {
	return string(l)
}

const (
	MOUNTAIN Category = "MOUNTAIN"
	BEACH             = "BEACH"
)

type GeneratedPhotosRequest struct {
	Category string `json:"category"`
}

// SetCategory set category
func (dR *GeneratedPhotosRequest) SetCategory(categ string) *GeneratedPhotosRequest {

	dR.Category = categ

	return dR
}

// ParseParams parse params
func (dR *GeneratedPhotosRequest) ParseParams(params []string) *GeneratedPhotosRequest {

	if len(params) == 0 {
		return dR
	}

	for _, param := range params {
		p := strings.Split(param, "=")

		if len(p) != 2 {
			continue
		}

		dR.SetParam(p[0], p[1])

	}

	return dR
}

// SetParam set param
func (dR *GeneratedPhotosRequest) SetParam(key, value string) {

	switch key {
	case "category":
		dR.SetCategory(value)
		break
	}
}

// GetURL returned key=value format for request on api
func (dR GeneratedPhotosRequest) GetURL(reqType string) string {

	// set request url
	url := generatedPhotosURL

	// add version
	switch reqType {
	case "GET":

		// define query params
		queryParams := []string{}

		url = "/v2" + url

		if dR.Category != "" {
			queryParams = append(queryParams, "category="+dR.Category)
		}

		return url + "?" + strings.Join(queryParams, "&")
	}

	return ""
}

// GetBody implementation for Request
func (dR GeneratedPhotosRequest) GetBody(reqType string) io.Reader {
	return nil
}

type GeneratedPhotosResponse struct {
	Data PhotosData `json:"data,omitempty"`
}

// Decode implement Response interface
func (dR *GeneratedPhotosResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type PhotosData struct {
	Type               string        `json:"type,omitempty"`
	Owner              string        `json:"owner,omitempty"`
	AttachmentURI      string        `json:"attachmentUri,omitempty"`
	Description        string        `json:"description,omitempty"`
	FileKbSize         int           `json:"fileKbSize,omitempty"`
	ExpirationDateTime string        `json:"expirationDateTime,omitempty"`
	MediaMetadata      MediaMetadata `json:"mediaMetadata,omitempty"`
}

type Dimensions struct {
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
	UOM    string `json:"UOM,omitempty"`
}

type MediaMetadata struct {
	Dimensions Dimensions `json:"dimensions,omitempty"`
}
