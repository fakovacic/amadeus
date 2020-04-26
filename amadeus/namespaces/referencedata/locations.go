package referencedata

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/structs"
)

const (

	// Airport & City Search
	// Airport Nearest Relevant
	locationsURL = "/reference-data/locations"
)

/*
const (
	AIRPORT  = string "AIRPORT"
	CITY           = "CITY"
)

const (
	FULL   = string "FULL"
	LIGHT        = "LIGHT"
)
*/
type LocationsRequest struct {
	LocationID  string   `json:"locationId"`
	Latitude    string   `json:"latitude"`
	Longitude   string   `json:"longitude"`
	Radius      string   `json:"radius"`
	SubType     []string `json:"subType"`
	Keyword     string   `json:"keyword"`
	CountryCode string   `json:"countryCode"`
	View        string   `json:"view"`
	Sort        string   `json:"sort"`

	// TODO
	//page[offset]
	//page[limit]
}

// SetLocationID set location id
func (dR *LocationsRequest) SetLocationID(id string) *LocationsRequest {

	dR.LocationID = id

	return dR
}

// SetLatitude set latitude
func (dR *LocationsRequest) SetLatitude(latitude string) *LocationsRequest {

	dR.Latitude = latitude

	return dR
}

// SetLongitude set longitude
func (dR *LocationsRequest) SetLongitude(longitude string) *LocationsRequest {

	dR.Longitude = longitude

	return dR
}

// SetRadius set radius
func (dR *LocationsRequest) SetRadius(radius string) *LocationsRequest {

	dR.Radius = radius

	return dR
}

// SetSort set sort
func (dR *LocationsRequest) SetSort(sort string) *LocationsRequest {

	dR.Sort = sort

	return dR
}

// SetView set sort
func (dR *LocationsRequest) SetView(view string) *LocationsRequest {

	dR.View = view

	return dR
}

// AddSubType add view
func (dR *LocationsRequest) AddSubType(view string) *LocationsRequest {

	dR.SubType = append(dR.SubType, view)

	return dR
}

// SetKeyword set keyword
func (dR *LocationsRequest) SetKeyword(keyword string) *LocationsRequest {

	dR.Keyword = keyword

	return dR
}

// SetCountryCode set countryCode
func (dR *LocationsRequest) SetCountryCode(countryCode string) *LocationsRequest {

	dR.CountryCode = countryCode

	return dR
}

// Helper functions

// SetGeo set latitude & longitude
func (dR *LocationsRequest) SetGeo(latitude, longitude string) *LocationsRequest {

	dR.SetLatitude(latitude)
	dR.SetLongitude(longitude)

	return dR
}

// GetByKeyword set
func (dR *LocationsRequest) GetByKeyword(keyword string, subTypes ...string) *LocationsRequest {

	dR.SetKeyword(keyword)

	if len(subTypes) != 0 {
		for _, sT := range subTypes {
			dR.AddSubType(sT)
		}

	}

	return dR
}

// ParseParams parse params
func (dR *LocationsRequest) ParseParams(params []string) *LocationsRequest {

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
func (dR *LocationsRequest) SetParam(key, value string) {

	switch key {
	case "locationId":
		dR.SetLocationID(value)
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
	case "subType":

		//
		// parse by ,
		//

		dR.AddSubType(value)
		break
	case "keyword":
		dR.SetKeyword(value)
		break
	case "countryCode":
		dR.SetCountryCode(value)
		break
	case "view":
		dR.SetView(value)
		break
	case "sort":
		dR.SetSort(value)
		break
	}
}

// GetURL returned key=value format for request on api
func (dR LocationsRequest) GetURL(reqType string) string {

	// set request url
	url := locationsURL

	// add version
	switch reqType {
	case "GET":

		// define query params
		queryParams := []string{}

		url = "/v1" + url

		if dR.LocationID != "" {

			return url + "/" + dR.LocationID

		}

		if dR.Latitude != "" && dR.Longitude != "" {

			queryParams = append(queryParams, "latitude="+dR.Latitude)
			queryParams = append(queryParams, "longitude="+dR.Longitude)

			if dR.Radius != "" {
				queryParams = append(queryParams, "radius="+dR.Radius)
			}

			// add sort

			return url + "/airports?" + strings.Join(queryParams, "&")

		}

		if dR.Keyword != "" && len(dR.SubType) != 0 {

			queryParams = append(queryParams, "keyword="+dR.Keyword)
			queryParams = append(queryParams, "subType="+strings.Join(dR.SubType, ","))

			if dR.CountryCode != "" {
				queryParams = append(queryParams, "countryCode="+dR.CountryCode)
			}

			if dR.View != "" {
				queryParams = append(queryParams, "view="+dR.View)
			}

			// add sort

			return url + "?" + strings.Join(queryParams, "&")
		}

		break
	}

	return ""
}

// GetBody implementation for Request
func (dR LocationsRequest) GetBody(reqType string) io.Reader {
	return nil
}

type LocationsResponse struct {
	Meta structs.Meta   `json:"meta"`
	Data []LocationData `json:"data"`
}

// Decode implement Response interface
func (dR *LocationsResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type LocationResponse struct {
	Meta structs.Meta `json:"meta"`
	Data LocationData `json:"data"`
}

// Decode implement Response interface
func (dR *LocationResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type LocationData struct {
	Type           string            `json:"type"`
	SubType        string            `json:"subType"`
	Name           string            `json:"name"`
	DetailedName   string            `json:"detailedName"`
	TimeZoneOffset string            `json:"timeZoneOffset"`
	IataCode       string            `json:"iataCode"`
	GeoCode        structs.GeoCode   `json:"geoCode"`
	Address        LocationAddress   `json:"address"`
	Distance       Distance          `json:"distance"`
	Analytics      LocationAnalytics `json:"analytics"`
	Relevance      float64           `json:"relevance"`
}

type LocationAddress struct {
	CityName    string `json:"cityName"`
	CityCode    string `json:"cityCode"`
	CountryName string `json:"countryName"`
	CountryCode string `json:"countryCode"`
	StateCode   string `json:"stateCode"`
	RegionCode  string `json:"regionCode"`
}

type Distance struct {
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

type LocationAnalytics struct {
	Flights   LocationFlights   `json:"flights"`
	Travelers LocationTravelers `json:"travelers"`
}

type LocationFlights struct {
	Score int `json:"score"`
}
type LocationTravelers struct {
	Score int `json:"score"`
}
