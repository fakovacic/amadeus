package amadeus

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type TravelAnalyticsAirTrafficType string

const (
	TRAVELED TravelAnalyticsAirTrafficType = "TRAVELED"
	BOOKED                                 = "BOOKED"
	BUSIEST                                = "BUSIEST"
)

type AirDirection string

const (
	ARRIVING  AirDirection = "ARRIVING"
	DEPARTING              = "DEPARTING"
)

type TravelAnalyticsAirTrafficRequest struct {
	Type      TravelAnalyticsAirTrafficType `json:"type"`
	Origin    string                        `json:"origin"`
	Period    string                        `json:"period"`
	Max       string                        `json:"max"`
	Sort      string                        `json:"sort"`
	Direction AirDirection                  `json:"direction"` // use for Booked destinations

	// TODO
	//Fields []string `json:"fields"`
}

// SetType set type
func (dR *TravelAnalyticsAirTrafficRequest) SetType(reqType TravelAnalyticsAirTrafficType) *TravelAnalyticsAirTrafficRequest {

	dR.Type = reqType

	return dR
}

// SetOrigin set origin
func (dR *TravelAnalyticsAirTrafficRequest) SetOrigin(origin string) *TravelAnalyticsAirTrafficRequest {

	dR.Origin = origin

	return dR
}

// SetPeriod set period
func (dR *TravelAnalyticsAirTrafficRequest) SetPeriod(period string) *TravelAnalyticsAirTrafficRequest {

	dR.Period = period

	return dR
}

// SetMax set max
func (dR *TravelAnalyticsAirTrafficRequest) SetMax(max string) *TravelAnalyticsAirTrafficRequest {

	dR.Max = max

	return dR
}

// SetSort set max
func (dR *TravelAnalyticsAirTrafficRequest) SetSort(sort string) *TravelAnalyticsAirTrafficRequest {

	dR.Sort = sort

	return dR
}

// SetDirection set max
func (dR *TravelAnalyticsAirTrafficRequest) SetDirection(direction AirDirection) *TravelAnalyticsAirTrafficRequest {

	dR.Direction = direction

	return dR
}

// GetURL returned key=value format for request on api
func (dR TravelAnalyticsAirTrafficRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := travelAnalyticsAirTrafficURL

	// add version
	switch reqType {
	case "GET":

		// define query params
		queryParams := []string{}

		url = "/v1" + url

		switch dR.Type {
		case TRAVELED:

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
		case BOOKED:

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
		case BUSIEST:

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

	return baseURL + url
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (dR TravelAnalyticsAirTrafficRequest) GetBody(reqType string) io.Reader {
	return nil
}

type TravelAnalyticsAirTrafficResponse struct {
	Warnings []Warnings      `json:"warnings,omitempty"`
	Meta     Meta            `json:"meta,omitempty"`
	Data     []AnayticsData  `json:"data,omitempty"`
	Errors   []ErrorResponse `json:"errors,omitempty"`
}

// Decode implement Response interface
func (dR *TravelAnalyticsAirTrafficResponse) Decode(rsp []byte) error {

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
