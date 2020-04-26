package analitycs

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/structs"
)

const (

	// Flight Busiest Traveling Period
	// Flight Most Booked Destinations
	// Flight Most Traveled Destinations
	airTrafficURL = "/travel/analytics/air-traffic"
)

/*
type AirTrafficType string

const (
	TRAVELED AirTrafficType = "TRAVELED"
	BOOKED                  = "BOOKED"
	BUSIEST                 = "BUSIEST"
)

type AirDirection string

const (
	ARRIVING  AirDirection = "ARRIVING"
	DEPARTING              = "DEPARTING"
)
*/

type AirTrafficRequest struct {
	Type      string `json:"type"`
	Origin    string `json:"origin"`
	Period    string `json:"period"`
	Max       string `json:"max"`
	Sort      string `json:"sort"`
	Direction string `json:"direction"` // use for Booked destinations

	// TODO
	//Fields []string `json:"fields"`
}

// SetType set type
func (dR *AirTrafficRequest) SetType(reqType string) *AirTrafficRequest {

	dR.Type = reqType

	return dR
}

// SetOrigin set origin
func (dR *AirTrafficRequest) SetOrigin(origin string) *AirTrafficRequest {

	dR.Origin = origin

	return dR
}

// SetPeriod set period
func (dR *AirTrafficRequest) SetPeriod(period string) *AirTrafficRequest {

	dR.Period = period

	return dR
}

// SetMax set max
func (dR *AirTrafficRequest) SetMax(max string) *AirTrafficRequest {

	dR.Max = max

	return dR
}

// SetSort set max
func (dR *AirTrafficRequest) SetSort(sort string) *AirTrafficRequest {

	dR.Sort = sort

	return dR
}

// SetDirection set max
func (dR *AirTrafficRequest) SetDirection(direction string) *AirTrafficRequest {

	dR.Direction = direction

	return dR
}

// ParseParams parse params
func (dR *AirTrafficRequest) ParseParams(params []string) *AirTrafficRequest {

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
func (dR *AirTrafficRequest) SetParam(key, value string) {

	switch key {
	case "origin":
		dR.SetOrigin(value)
		break
	case "period":
		dR.SetPeriod(value)
		break
	case "max":
		dR.SetMax(value)
		break
	case "sort":
		dR.SetSort(value)
		break
	case "direction":
		dR.SetDirection(value)
		break
	}
}

// GetURL returned key=value format for request on api
func (dR AirTrafficRequest) GetURL(reqType string) string {

	// set request url
	url := airTrafficURL

	// add version
	switch reqType {
	case "GET":

		// define query params
		queryParams := []string{}

		url = "/v1" + url

		switch dR.Type {
		case "TRAVELED":

			url = url + "/traveled"

			queryParams = append(queryParams, "originCityCode="+dR.Origin)
			queryParams = append(queryParams, "period="+dR.Period)

			if dR.Sort != "" {
				queryParams = append(queryParams, "sort="+dR.Sort)
			}

			if dR.Max != "" {
				queryParams = append(queryParams, "max="+dR.Max)
			}

			break
		case "BOOKED":

			url = url + "/booked"

			queryParams = append(queryParams, "originCityCode="+dR.Origin)
			queryParams = append(queryParams, "period="+dR.Period)

			if dR.Sort != "" {
				queryParams = append(queryParams, "sort="+dR.Sort)
			}

			if dR.Max != "" {
				queryParams = append(queryParams, "max="+dR.Max)
			}

			break
		case "BUSIEST":

			url = url + "/busiest-period"

			queryParams = append(queryParams, "cityCode="+dR.Origin)
			queryParams = append(queryParams, "period="+dR.Period)

			if dR.Direction != "" {
				queryParams = append(queryParams, "direction="+fmt.Sprintf("%v", dR.Direction))
			}

			break
		}

		url = url + "?" + strings.Join(queryParams, "&")

		break
	}

	return url
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (dR AirTrafficRequest) GetBody(reqType string) io.Reader {
	return nil
}

type AirTrafficResponse struct {
	Meta     structs.Meta       `json:"meta,omitempty"`
	Data     []AnayticsData     `json:"data,omitempty"`
	Warnings []structs.Warnings `json:"warnings,omitempty"`
}

// Decode implement Response interface
func (dR *AirTrafficResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type AnayticsData struct {
	Type        string    `json:"type,omitempty"`
	SubType     string    `json:"subType,omitempty"`
	Destination string    `json:"destination,omitempty"`
	Analytics   Analytics `json:"analytics,omitempty"`
}

type Analytics struct {
	Flights   AnalyticsFlights   `json:"flights,omitempty"`
	Travelers AnalyticsTravelers `json:"travelers,omitempty"`
}

type AnalyticsFlights struct {
	Score int `json:"score,omitempty"`
}
type AnalyticsTravelers struct {
	Score int `json:"score,omitempty"`
}
