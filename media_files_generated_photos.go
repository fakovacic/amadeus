package amadeus

import (
	"encoding/json"
	"io"
	"strings"
)

type PhotosCategory string

func (l PhotosCategory) String() string {
	return string(l)
}

const (
	MOUNTAIN PhotosCategory = "MOUNTAIN"
	BEACH                   = "BEACH"
)

type MediaFilesGeneratedPhotosRequest struct {
	Category PhotosCategory `json:"category"`
}

// SetCategory set category
func (dR *MediaFilesGeneratedPhotosRequest) SetCategory(categ PhotosCategory) *MediaFilesGeneratedPhotosRequest {

	dR.Category = categ

	return dR
}

// GetURL returned key=value format for request on api
func (dR MediaFilesGeneratedPhotosRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := mediaFilesGeneratedPhotosURL

	// add version
	switch reqType {
	case "GET":

		// define query params
		queryParams := []string{}

		url = baseURL + "/v2" + url

		if dR.Category != "" {
			queryParams = append(queryParams, "category="+dR.Category.String())
		}

		return url + "?" + strings.Join(queryParams, "&")
	}

	return ""
}

// GetBody implementation for Request
func (dR MediaFilesGeneratedPhotosRequest) GetBody(reqType string) io.Reader {
	return nil
}

type MediaFilesGeneratedPhotosResponse struct {
	Data   PhotosData      `json:"data,omitempty"`
	Errors []ErrorResponse `json:"errors,omitempty"`
}

// Decode implement Response interface
func (dR *MediaFilesGeneratedPhotosResponse) Decode(rsp []byte) error {

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
