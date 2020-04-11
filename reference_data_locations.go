package amadeus

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type LocationSubType string

func (l LocationSubType) String() string {
	return string(l)
}

const (
	AIRPORT LocationSubType = "AIRPORT"
	CITY                    = "CITY"
)

type LocationViewType string

const (
	FULL  LocationViewType = "FULL"
	LIGHT                  = "LIGHT"
)

type ReferenceDataLocationsRequest struct {
	LocationID  string           `json:"locationId"`
	Latitude    string           `json:"latitude"`
	Longitude   string           `json:"longitude"`
	Radius      string           `json:"radius"`
	SubType     []string         `json:"subType"`
	Keyword     string           `json:"keyword"`
	CountryCode string           `json:"countryCode"`
	View        LocationViewType `json:"view"`
	Sort        string           `json:"sort"`

	// TODO
	//page[offset]
	//page[limit]
}

// SetLocationID set location id
func (dR *ReferenceDataLocationsRequest) SetLocationID(id string) *ReferenceDataLocationsRequest {

	dR.LocationID = id

	return dR
}

// SetLatitude set latitude
func (dR *ReferenceDataLocationsRequest) SetLatitude(latitude string) *ReferenceDataLocationsRequest {

	dR.Latitude = latitude

	return dR
}

// SetLongitude set longitude
func (dR *ReferenceDataLocationsRequest) SetLongitude(longitude string) *ReferenceDataLocationsRequest {

	dR.Longitude = longitude

	return dR
}

// SetRadius set radius
func (dR *ReferenceDataLocationsRequest) SetRadius(radius string) *ReferenceDataLocationsRequest {

	dR.Radius = radius

	return dR
}

// SetSort set sort
func (dR *ReferenceDataLocationsRequest) SetSort(sort string) *ReferenceDataLocationsRequest {

	dR.Sort = sort

	return dR
}

// SetView set sort
func (dR *ReferenceDataLocationsRequest) SetView(view LocationViewType) *ReferenceDataLocationsRequest {

	dR.View = view

	return dR
}

// AddSubType add view
func (dR *ReferenceDataLocationsRequest) AddSubType(view LocationSubType) *ReferenceDataLocationsRequest {

	dR.SubType = append(dR.SubType, view.String())

	return dR
}

// SetKeyword set keyword
func (dR *ReferenceDataLocationsRequest) SetKeyword(keyword string) *ReferenceDataLocationsRequest {

	dR.Keyword = keyword

	return dR
}

// SetCountryCode set countryCode
func (dR *ReferenceDataLocationsRequest) SetCountryCode(countryCode string) *ReferenceDataLocationsRequest {

	dR.CountryCode = countryCode

	return dR
}

// Helper functions

// SetGeo set latitude & longitude
func (dR *ReferenceDataLocationsRequest) SetGeo(latitude, longitude string) *ReferenceDataLocationsRequest {

	dR.SetLatitude(latitude)
	dR.SetLongitude(longitude)

	return dR
}

// GetByKeyword set
func (dR *ReferenceDataLocationsRequest) GetByKeyword(keyword string, subTypes ...LocationSubType) *ReferenceDataLocationsRequest {

	dR.SetKeyword(keyword)

	if len(subTypes) != 0 {
		for _, sT := range subTypes {
			dR.AddSubType(sT)
		}

	}

	return dR
}

// GetURL returned key=value format for request on api
func (dR ReferenceDataLocationsRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := referenceDataLocationsURL

	// add version
	switch reqType {
	case "GET":

		// define query params
		queryParams := []string{}

		url = baseURL + "/v1" + url

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
				queryParams = append(queryParams, "view="+fmt.Sprintf("%v", dR.View))
			}

			// add sort

			return url + "?" + strings.Join(queryParams, "&")
		}

		break
	}

	return ""
}

// GetBody implementation for Request
func (dR ReferenceDataLocationsRequest) GetBody(reqType string) io.Reader {
	return nil
}

type ReferenceDataLocationsResponse struct {
	Meta Meta           `json:"meta"`
	Data []LocationData `json:"data"`
}

// Decode implement Response interface
func (dR *ReferenceDataLocationsResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type ReferenceDataLocationResponse struct {
	Meta Meta         `json:"meta"`
	Data LocationData `json:"data"`
}

// Decode implement Response interface
func (dR *ReferenceDataLocationResponse) Decode(rsp []byte) error {

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
	GeoCode        GeoCode           `json:"geoCode"`
	Address        LocationAddress   `json:"address"`
	Distance       Distance          `json:"distance"`
	Analytics      LocationAnalytics `json:"analytics"`
	Relevance      float64           `json:"relevance"`
}

type GeoCode struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
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
