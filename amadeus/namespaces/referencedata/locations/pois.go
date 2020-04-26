package locations

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/structs"
)

const (

	// Points of Interest
	poisURL = "/reference-data/locations/pois"
)

/*
type Category string

func (c Category) String() string {
	return string(c)
}

const (
	SIGHTS     Category = "SIGHTS"
	NIGHTLIFE           = "NIGHTLIFE"
	RESTARAUNT          = "RESTARAUNT"
	SHOOPING            = "SHOOPING"
)
*/
type PoisRequest struct {
	PoisID     string   `json:"poisId"`
	Latitude   string   `json:"latitude"`
	Longitude  string   `json:"longitude"`
	Radius     string   `json:"radius"`
	North      string   `json:"north"`
	West       string   `json:"west"`
	South      string   `json:"south"`
	East       string   `json:"east"`
	Categories []string `json:"categories"`
	// TODO
	//page[offset]
	//page[limit]
}

// SetPoisID set location id
func (dR *PoisRequest) SetPoisID(id string) *PoisRequest {

	dR.PoisID = id

	return dR
}

// SetLatitude set latitude
func (dR *PoisRequest) SetLatitude(latitude string) *PoisRequest {

	dR.Latitude = latitude

	return dR
}

// SetLongitude set longitude
func (dR *PoisRequest) SetLongitude(longitude string) *PoisRequest {

	dR.Longitude = longitude

	return dR
}

// SetNorth set north
func (dR *PoisRequest) SetNorth(north string) *PoisRequest {

	dR.North = north

	return dR
}

// SetWest set west
func (dR *PoisRequest) SetWest(west string) *PoisRequest {

	dR.West = west

	return dR
}

// SetSouth set south
func (dR *PoisRequest) SetSouth(south string) *PoisRequest {

	dR.South = south

	return dR
}

// SetEast set east
func (dR *PoisRequest) SetEast(east string) *PoisRequest {

	dR.East = east

	return dR
}

// SetRadius set radius
func (dR *PoisRequest) SetRadius(radius string) *PoisRequest {

	dR.Radius = radius

	return dR
}

// AddCategory add category
func (dR *PoisRequest) AddCategory(categs ...string) *PoisRequest {

	dR.Categories = categs

	return dR
}

// Helper functions

// SetGeo set latitude & longitude
func (dR *PoisRequest) SetGeo(latitude, longitude string) *PoisRequest {

	dR.SetLatitude(latitude)
	dR.SetLongitude(longitude)

	return dR
}

// SetSquare set north, west, east, south
func (dR *PoisRequest) SetSquare(north, west, south, east string) *PoisRequest {

	dR.SetNorth(north)
	dR.SetWest(west)
	dR.SetSouth(south)
	dR.SetEast(east)

	return dR
}

/*
// JoinCategs join category slice to string
func (dR PoisRequest) JoinCategs(categSlice []Category) string {

	valuesText := []string{}

	for i := range categSlice {
		c := categSlice[i]
		valuesText = append(valuesText, c.String())
	}

	return strings.Join(valuesText, ",")
}
*/

// ParseParams parse params
func (dR *PoisRequest) ParseParams(params []string) *PoisRequest {

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
func (dR *PoisRequest) SetParam(key, value string) {

	switch key {
	case "poisId":
		dR.SetPoisID(value)
		break
	case "latitude":
		dR.SetLatitude(value)
		break
	case "longitude":
		dR.SetLongitude(value)
		break
	case "radius":
		dR.SetRadius(value)
		break
	case "north":
		dR.SetNorth(value)
		break
	case "west":
		dR.SetWest(value)
		break
	case "south":
		dR.SetSouth(value)
		break
	case "east":
		dR.SetEast(value)
		break
	case "category":
		dR.AddCategory(value)
		break
	}
}

// GetURL returned key=value format for request on api
func (dR PoisRequest) GetURL(reqType string) string {

	// set request url
	url := poisURL

	// add version
	switch reqType {
	case "GET":

		// define query params
		queryParams := []string{}

		url = "/v1" + url

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
			//queryParams = append(queryParams, "category="+dR.JoinCategs(dR.Categories))
			queryParams = append(queryParams, "category="+strings.Join(dR.Categories, ","))
		}

		return url + "?" + strings.Join(queryParams, "&")
	}

	return ""
}

// GetBody implementation for Request
func (dR PoisRequest) GetBody(reqType string) io.Reader {
	return nil
}

type PoisResponse struct {
	Meta structs.Meta `json:"meta"`
	Data []PoisData   `json:"data"`
}

// Decode implement Response interface
func (dR *PoisResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type PoiResponse struct {
	Meta structs.Meta `json:"meta"`
	Data PoisData     `json:"data"`
}

// Decode implement Response interface
func (dR *PoiResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type PoisData struct {
	Type     string          `json:"type"`
	ID       string          `json:"id"`
	Self     PoisSelf        `json:"self"`
	SubType  string          `json:"subType"`
	Name     string          `json:"name"`
	Rank     int             `json:"rank"`
	GeoCode  structs.GeoCode `json:"geoCode"`
	Category string          `json:"category"`
	Tags     []string        `json:"tags"`
}

type PoisSelf struct {
	Href    string   `json:"href"`
	Methods []string `json:"methods"`
}
