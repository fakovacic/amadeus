package ereputation

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/structs"
)

const (
	// Hotel Ratings
	// get hotel reputation by hotel ids
	hotelSentimentsURL = "/e-reputation/hotel-sentiments"
)

type HotelSentimentsRequest struct {
	HotelIDs []string
}

// AddHotelIDs add amenities
func (dR *HotelSentimentsRequest) AddHotelIDs(hotelIDs ...string) *HotelSentimentsRequest {

	dR.HotelIDs = hotelIDs

	return dR
}

// ParseParams parse params
func (dR *HotelSentimentsRequest) ParseParams(params []string) *HotelSentimentsRequest {

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
func (dR *HotelSentimentsRequest) SetParam(key, value string) {

	switch key {
	case "hotelIds":

		//
		// split by ,
		//

		dR.AddHotelIDs(value)
		break
	}
}

// GetURL returned key=value format for request on api
func (dR HotelSentimentsRequest) GetURL(reqType string) string {

	// set request url
	url := hotelSentimentsURL

	// add version
	switch reqType {
	case "GET":

		// define query params
		queryParams := []string{}

		url = "/v2" + url

		if len(dR.HotelIDs) != 0 {
			queryParams = append(queryParams, "hotelIds="+strings.Join(dR.HotelIDs, ","))
		}

		return url + "?" + strings.Join(queryParams, "&")
	}

	return ""
}

// GetBody implementation for Request
func (dR HotelSentimentsRequest) GetBody(reqType string) io.Reader {
	return nil
}

type HotelSentimentsResponse struct {
	Meta     structs.Meta       `json:"meta"`
	Warnings []structs.Warnings `json:"warnings"`
	Data     []SentimentsData   `json:"data"`
}

// Decode implement Response interface
func (dR *HotelSentimentsResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type SentimentsData struct {
	HotelID         string     `json:"hotelId"`
	Type            string     `json:"type"`
	OverallRating   int        `json:"overallRating"`
	NumberOfRatings int        `json:"numberOfRatings"`
	NumberOfReviews int        `json:"numberOfReviews"`
	Sentiments      Sentiments `json:"sentiments"`
}

type Sentiments struct {
	SleepQuality     int `json:"sleepQuality"`
	Service          int `json:"service"`
	Facilities       int `json:"facilities"`
	RoomComforts     int `json:"roomComforts"`
	ValueForMoney    int `json:"valueForMoney"`
	Catering         int `json:"catering"`
	SwimmingPool     int `json:"swimmingPool"`
	Location         int `json:"location"`
	Internet         int `json:"internet"`
	PointsOfInterest int `json:"pointsOfInterest"`
	Staff            int `json:"staff"`
}
