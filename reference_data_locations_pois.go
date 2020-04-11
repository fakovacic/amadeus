package amadeus

import (
	"encoding/json"
	"io"
	"strings"
)

type PoisCategory string

func (c PoisCategory) String() string {
	return string(c)
}

const (
	SIGHTS     PoisCategory = "SIGHTS"
	NIGHTLIFE               = "NIGHTLIFE"
	RESTARAUNT              = "RESTARAUNT"
	SHOOPING                = "SHOOPING"
)

type ReferenceDataLocationsPoisRequest struct {
	PoisID     string         `json:"poisId"`
	Latitude   string         `json:"latitude"`
	Longitude  string         `json:"longitude"`
	Radius     string         `json:"radius"`
	North      string         `json:"north"`
	West       string         `json:"west"`
	South      string         `json:"south"`
	East       string         `json:"east"`
	Categories []PoisCategory `json:"view"`
	// TODO
	//page[offset]
	//page[limit]
}

// SetPoisID set location id
func (dR *ReferenceDataLocationsPoisRequest) SetPoisID(id string) *ReferenceDataLocationsPoisRequest {

	dR.PoisID = id

	return dR
}

// SetLatitude set latitude
func (dR *ReferenceDataLocationsPoisRequest) SetLatitude(latitude string) *ReferenceDataLocationsPoisRequest {

	dR.Latitude = latitude

	return dR
}

// SetLongitude set longitude
func (dR *ReferenceDataLocationsPoisRequest) SetLongitude(longitude string) *ReferenceDataLocationsPoisRequest {

	dR.Longitude = longitude

	return dR
}

// SetNorth set north
func (dR *ReferenceDataLocationsPoisRequest) SetNorth(north string) *ReferenceDataLocationsPoisRequest {

	dR.North = north

	return dR
}

// SetWest set west
func (dR *ReferenceDataLocationsPoisRequest) SetWest(west string) *ReferenceDataLocationsPoisRequest {

	dR.West = west

	return dR
}

// SetSouth set south
func (dR *ReferenceDataLocationsPoisRequest) SetSouth(south string) *ReferenceDataLocationsPoisRequest {

	dR.South = south

	return dR
}

// SetEast set east
func (dR *ReferenceDataLocationsPoisRequest) SetEast(east string) *ReferenceDataLocationsPoisRequest {

	dR.East = east

	return dR
}

// SetRadius set radius
func (dR *ReferenceDataLocationsPoisRequest) SetRadius(radius string) *ReferenceDataLocationsPoisRequest {

	dR.Radius = radius

	return dR
}

// AddCategory add category
func (dR *ReferenceDataLocationsPoisRequest) AddCategory(categs ...PoisCategory) *ReferenceDataLocationsPoisRequest {

	dR.Categories = categs

	return dR
}

// Helper functions

// SetGeo set latitude & longitude
func (dR *ReferenceDataLocationsPoisRequest) SetGeo(latitude, longitude string) *ReferenceDataLocationsPoisRequest {

	dR.SetLatitude(latitude)
	dR.SetLongitude(longitude)

	return dR
}

// SetSquare set north, west, east, south
func (dR *ReferenceDataLocationsPoisRequest) SetSquare(north, west, south, east string) *ReferenceDataLocationsPoisRequest {

	dR.SetNorth(north)
	dR.SetWest(west)
	dR.SetSouth(south)
	dR.SetEast(east)

	return dR
}

// JoinCategs join category slice to string
func (dR ReferenceDataLocationsPoisRequest) JoinCategs(categSlice []PoisCategory) string {

	valuesText := []string{}

	for i := range categSlice {
		c := categSlice[i]
		valuesText = append(valuesText, c.String())
	}

	return strings.Join(valuesText, ",")
}

// GetURL returned key=value format for request on api
func (dR ReferenceDataLocationsPoisRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := referenceDataLocationsPoisURL

	// add version
	switch reqType {
	case "GET":

		// define query params
		queryParams := []string{}

		url = baseURL + "/v1" + url

		if dR.PoisID != "" {

			return url + "/" + dR.PoisID

		}

		if dR.Latitude != "" && dR.Longitude != "" {

			queryParams = append(queryParams, "latitude="+dR.Latitude)
			queryParams = append(queryParams, "longitude="+dR.Longitude)

			if dR.Radius != "" {
				queryParams = append(queryParams, "radius="+dR.Radius)
			}

		}

		if dR.North != "" && dR.South != "" && dR.East != "" && dR.West != "" {

			url = url + "/by-square"

			queryParams = append(queryParams, "north="+dR.North)
			queryParams = append(queryParams, "south="+dR.South)
			queryParams = append(queryParams, "east="+dR.East)
			queryParams = append(queryParams, "west="+dR.West)

		}

		if len(dR.Categories) != 0 {
			queryParams = append(queryParams, "category="+dR.JoinCategs(dR.Categories))
		}

		return url + "?" + strings.Join(queryParams, "&")
	}

	return ""
}

// GetBody implementation for Request
func (dR ReferenceDataLocationsPoisRequest) GetBody(reqType string) io.Reader {
	return nil
}

type ReferenceDataLocationsPoisResponse struct {
	Meta Meta       `json:"meta"`
	Data []PoisData `json:"data"`
}

// Decode implement Response interface
func (dR *ReferenceDataLocationsPoisResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type ReferenceDataLocationsPoiResponse struct {
	Meta Meta     `json:"meta"`
	Data PoisData `json:"data"`
}

// Decode implement Response interface
func (dR *ReferenceDataLocationsPoiResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type PoisData struct {
	Type     string   `json:"type"`
	ID       string   `json:"id"`
	Self     PoisSelf `json:"self"`
	SubType  string   `json:"subType"`
	Name     string   `json:"name"`
	Rank     int      `json:"rank"`
	GeoCode  GeoCode  `json:"geoCode"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

type PoisSelf struct {
	Href    string   `json:"href"`
	Methods []string `json:"methods"`
}
