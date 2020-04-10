package amadeus

import (
	"encoding/json"
	"io"
	"strings"
)

type EReputationHotelSentimentsRequest struct {
	HotelIDs []string
}

// AddHotelIDs add amenities
func (dR *EReputationHotelSentimentsRequest) AddHotelIDs(hotelIDs ...string) *EReputationHotelSentimentsRequest {

	dR.HotelIDs = hotelIDs

	return dR
}

// GetURL returned key=value format for request on api
func (dR EReputationHotelSentimentsRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := ereputationHotelSentimentsURL

	// add version
	switch reqType {
	case "GET":

		// define query params
		queryParams := []string{}

		url = baseURL + "/v2" + url

		if len(dR.HotelIDs) != 0 {
			queryParams = append(queryParams, "hotelIds="+strings.Join(dR.HotelIDs, ","))
		}

		return url + "?" + strings.Join(queryParams, "&")
	}

	return ""
}

// GetBody implementation for Request
func (dR EReputationHotelSentimentsRequest) GetBody(reqType string) io.Reader {
	return nil
}

type EReputationHotelSentimentsResponse struct {
	Meta     Meta             `json:"meta"`
	Warnings []Warnings       `json:"warnings"`
	Data     []SentimentsData `json:"data"`
}

// Decode implement Response interface
func (dR *EReputationHotelSentimentsResponse) Decode(rsp []byte) error {

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
