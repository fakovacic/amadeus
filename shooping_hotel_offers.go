package amadeus

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type HotelView string

func (p HotelView) String() string {
	return string(p)
}

const (
	//HOTELVIEW_NONE geocoordinates, hotel distance
	HOTELVIEW_NONE HotelView = "NONE"
	//HOTELVIEW_LIGHT: NONE view + city name, phone number, fax, address, postal code, country code, state code, ratings, 1 image
	HOTELVIEW_LIGHT = "LIGHT"
	//HOTELVIEW_FULL LIGHT view + hotel description, amenities and facilities
	HOTELVIEW_FULL = "FULL"
)

type HotelSort string

func (p HotelSort) String() string {
	return string(p)
}

const (
	DISTANCE HotelSort = "DISTANCE"
	PRICE              = "PRICE"
)

type HotelBoardType string

func (p HotelBoardType) String() string {
	return string(p)
}

const (
	ROOM_ONLY     HotelBoardType = "RO"  // Room Only
	BREAKFAST                    = "BB"  // Bed & Breakfast
	HALF_BOARD                   = "DBB" // Diner & Bed & Breakfast (only for Aggregators)
	FULL_BOARD                   = "FB"  // Full Board (only for Aggregators)
	ALL_INCLUSIVE                = "AI"  // All Inclusive (only for Aggregators)
)

type HotelPaymentPolicy string

func (p HotelPaymentPolicy) String() string {
	return string(p)
}

const (
	DEPOSIT   HotelPaymentPolicy = "DEPOSIT"
	GUARANTEE                    = "GUARANTEE"
)

type RadiusUnit string

func (r RadiusUnit) String() string {
	return string(r)
}

const (
	MILE RadiusUnit = "MILE"
	KM              = "KM"
)

type ShoopingHotelOffersRequest struct {
	OfferID       string             `json:"offerId,omitempty"`
	HotelID       string             `json:"hotelId,omitempty"`
	CityCode      string             `json:"cityCode,omitempty"`
	CheckInDate   string             `json:"checkInDate"`
	CheckOutDate  string             `json:"checkOutDate"`
	RoomQuantity  int                `json:"roomQuantity"`
	Adults        int                `json:"adults"`
	ChildAges     []int              `json:"childAges"`
	HotelName     string             `json:"hotelName"`
	Latitude      float64            `json:"latitude"`
	Longitude     float64            `json:"longitude"`
	Radius        int                `json:"radius"`
	RadiusUnit    RadiusUnit         `json:"radiusUnit"`
	HotelIDs      []string           `json:"hotelIds"`
	Chains        []string           `json:"chains"`
	RateCodes     []string           `json:"rateCodes"`
	Amenities     []string           `json:"amenities"`
	Ratings       []int              `json:"ratings"`
	PriceRange    string             `json:"priceRange"`
	Currency      string             `json:"currency"`
	PaymentPolicy HotelPaymentPolicy `json:"paymentPolicy"`
	BoardType     HotelBoardType     `json:"boardType"`
	IncludeClosed bool               `json:"includeClosed"`
	BestRateOnly  bool               `json:"bestRateOnly"`
	View          HotelView          `json:"view"`
	Sort          HotelSort          `json:"sort"`
	Lang          string             `json:"lang"`
	//page[limit]
	//page[offset]
}

// SetOfferID set offer id
func (sR *ShoopingHotelOffersRequest) SetOfferID(offerID string) *ShoopingHotelOffersRequest {

	sR.OfferID = offerID

	return sR
}

// SetHotelID set hotel id
func (sR *ShoopingHotelOffersRequest) SetHotelID(hotelID string) *ShoopingHotelOffersRequest {

	sR.HotelID = hotelID

	return sR
}

// SetCityCode set city code
func (sR *ShoopingHotelOffersRequest) SetCityCode(cityCode string) *ShoopingHotelOffersRequest {

	sR.CityCode = cityCode

	return sR
}

// SetCheckInDate set checkin date
func (sR *ShoopingHotelOffersRequest) SetCheckInDate(checkInDate string) *ShoopingHotelOffersRequest {

	// check date

	sR.CheckInDate = checkInDate

	return sR
}

// SetCheckOutDate set checkout date
func (sR *ShoopingHotelOffersRequest) SetCheckOutDate(checkOutDate string) *ShoopingHotelOffersRequest {

	// check date

	sR.CheckOutDate = checkOutDate

	return sR
}

// SetAdults set adults
func (sR *ShoopingHotelOffersRequest) SetAdults(adults int) *ShoopingHotelOffersRequest {

	// check date

	sR.Adults = adults

	return sR
}

// AddChildAges add childAges
func (sR *ShoopingHotelOffersRequest) AddChildAges(childAges ...int) *ShoopingHotelOffersRequest {

	sR.ChildAges = childAges

	return sR
}

// SetHotelName set hotel name
func (sR *ShoopingHotelOffersRequest) SetHotelName(hotelName string) *ShoopingHotelOffersRequest {

	sR.HotelName = hotelName

	return sR
}

// SetLatitude set latitude
func (sR *ShoopingHotelOffersRequest) SetLatitude(latitude float64) *ShoopingHotelOffersRequest {

	sR.Latitude = latitude

	return sR
}

// SetLongitude set longitude
func (sR *ShoopingHotelOffersRequest) SetLongitude(longitude float64) *ShoopingHotelOffersRequest {

	sR.Longitude = longitude

	return sR
}

// SetRadius set radius
func (sR *ShoopingHotelOffersRequest) SetRadius(radius int) *ShoopingHotelOffersRequest {

	sR.Radius = radius

	return sR
}

// SetRadiusUnit set radiusUnit | KM or MILE
func (sR *ShoopingHotelOffersRequest) SetRadiusUnit(radiusUnit RadiusUnit) *ShoopingHotelOffersRequest {

	sR.RadiusUnit = radiusUnit

	return sR
}

// AddHotelIDs add hotelIDs
func (sR *ShoopingHotelOffersRequest) AddHotelIDs(hotelIDs ...string) *ShoopingHotelOffersRequest {

	sR.HotelIDs = hotelIDs

	return sR
}

// AddChains add hotel chains filter
func (sR *ShoopingHotelOffersRequest) AddChains(chains ...string) *ShoopingHotelOffersRequest {

	sR.Chains = chains

	return sR
}

// AddRateCodes add rateCodes
func (sR *ShoopingHotelOffersRequest) AddRateCodes(rateCodes ...string) *ShoopingHotelOffersRequest {

	sR.RateCodes = rateCodes

	return sR
}

// AddAmenities add amenities
func (sR *ShoopingHotelOffersRequest) AddAmenities(amenities ...string) *ShoopingHotelOffersRequest {

	sR.Amenities = amenities

	return sR
}

// AddRatings add ratings
func (sR *ShoopingHotelOffersRequest) AddRatings(ratings ...int) *ShoopingHotelOffersRequest {

	sR.Ratings = ratings

	return sR
}

// SetPriceRange set priceRange
func (sR *ShoopingHotelOffersRequest) SetPriceRange(priceRange string) *ShoopingHotelOffersRequest {

	sR.PriceRange = priceRange

	return sR
}

// SetCurrency set currency
func (sR *ShoopingHotelOffersRequest) SetCurrency(currency string) *ShoopingHotelOffersRequest {

	sR.Currency = currency

	return sR
}

// SetPaymentPolicy set payment policy
func (sR *ShoopingHotelOffersRequest) SetPaymentPolicy(paymentPolicy HotelPaymentPolicy) *ShoopingHotelOffersRequest {

	sR.PaymentPolicy = paymentPolicy

	return sR
}

// SetBoardType set board type
func (sR *ShoopingHotelOffersRequest) SetBoardType(BoardType HotelBoardType) *ShoopingHotelOffersRequest {

	sR.BoardType = BoardType

	return sR
}

// IsIncludeClosed include closed
func (sR *ShoopingHotelOffersRequest) IsIncludeClosed(includeClosed bool) *ShoopingHotelOffersRequest {

	if includeClosed {
		sR.IncludeClosed = true
	}

	return sR
}

// IsBestRateOnly best rate only
func (sR *ShoopingHotelOffersRequest) IsBestRateOnly(bestRateOnly bool) *ShoopingHotelOffersRequest {

	if bestRateOnly {
		sR.BestRateOnly = true
	}

	return sR
}

// SetView set view
func (sR *ShoopingHotelOffersRequest) SetView(view HotelView) *ShoopingHotelOffersRequest {

	sR.View = view

	return sR
}

// SetSort set sort
func (sR *ShoopingHotelOffersRequest) SetSort(sort HotelSort) *ShoopingHotelOffersRequest {

	sR.Sort = sort

	return sR
}

// SetLang set lang
func (sR *ShoopingHotelOffersRequest) SetLang(lang string) *ShoopingHotelOffersRequest {

	sR.Lang = lang

	return sR
}

// Helper functions

// SetGeo set latitude & longitude
func (sR *ShoopingHotelOffersRequest) SetGeo(latitude, longitude float64) *ShoopingHotelOffersRequest {

	sR.SetLatitude(latitude)
	sR.SetLongitude(longitude)

	return sR
}

// JoinInts join int slice to string
func (sR ShoopingHotelOffersRequest) JoinInts(intSlice []int) string {

	valuesText := []string{}

	for i := range intSlice {
		number := intSlice[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	return strings.Join(valuesText, ",")
}

// GetURL returned key=value format for request on api
func (sR ShoopingHotelOffersRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := shoopingHotelOffersURL

	// add version
	switch reqType {
	case "GET":

		url = "/v2" + url

		// define query params
		queryParams := []string{}

		// check request values
		if sR.OfferID != "" {

			return baseURL + url + "/" + sR.OfferID

		} else if sR.HotelID != "" {

			url = url + "/by-hotel"

			queryParams = append(queryParams, "hotelId="+sR.HotelID)

		} else if sR.CityCode != "" {
			queryParams = append(queryParams, "cityCode="+sR.CityCode)
		} else if sR.Latitude != 0 && sR.Longitude != 0 {
			queryParams = append(queryParams, "latitude="+fmt.Sprintf("%v", sR.Latitude))
			queryParams = append(queryParams, "longitude="+fmt.Sprintf("%v", sR.Longitude))
		}

		if sR.CheckInDate != "" {
			queryParams = append(queryParams, "checkInDate="+sR.CheckInDate)
		}

		if sR.CheckOutDate != "" {
			queryParams = append(queryParams, "checkOutDate="+sR.CheckOutDate)
		}

		if sR.RoomQuantity != 0 {
			queryParams = append(queryParams, "roomQuantity="+fmt.Sprintf("%v", sR.RoomQuantity))
		}

		if sR.Adults != 0 {
			queryParams = append(queryParams, "adults="+fmt.Sprintf("%v", sR.Adults))
		}

		if len(sR.ChildAges) != 0 {
			queryParams = append(queryParams, "childAges="+sR.JoinInts(sR.ChildAges))
		}

		if sR.HotelName != "" {
			queryParams = append(queryParams, "hotelName="+sR.HotelName)
		}

		if len(sR.HotelIDs) != 0 {
			queryParams = append(queryParams, "hotelIds="+strings.Join(sR.HotelIDs, ","))
		}

		if len(sR.Chains) != 0 {
			queryParams = append(queryParams, "chains="+strings.Join(sR.Chains, ","))
		}

		if len(sR.RateCodes) != 0 {
			queryParams = append(queryParams, "rateCodes="+strings.Join(sR.RateCodes, ","))
		}

		if len(sR.Amenities) != 0 {
			queryParams = append(queryParams, "amenities="+strings.Join(sR.Amenities, ","))
		}

		if len(sR.Ratings) != 0 {
			queryParams = append(queryParams, "ratings="+sR.JoinInts(sR.Ratings))
		}

		if sR.Radius != 0 && sR.RadiusUnit != "" {
			queryParams = append(queryParams, "radius="+fmt.Sprintf("%v", sR.Radius))
			queryParams = append(queryParams, "radiusUnit="+sR.RadiusUnit.String())
		}

		if sR.PriceRange != "" && sR.Currency != "" {
			queryParams = append(queryParams, "priceRange="+sR.PriceRange)
			queryParams = append(queryParams, "currency="+sR.Currency)
		}

		if sR.PaymentPolicy != "" {
			queryParams = append(queryParams, "paymentPolicy="+sR.PaymentPolicy.String())
		}

		if sR.BoardType != "" {
			queryParams = append(queryParams, "boardType="+sR.BoardType.String())
		}

		if sR.IncludeClosed {
			queryParams = append(queryParams, "includeClosed=true")
		}

		if sR.BestRateOnly {
			queryParams = append(queryParams, "bestRateOnly=true")
		}

		if sR.View != "" {
			queryParams = append(queryParams, "view="+sR.View.String())
		}

		if sR.Sort != "" {
			queryParams = append(queryParams, "sort="+sR.Sort.String())
		}

		if sR.Lang != "" {
			queryParams = append(queryParams, "lang="+sR.Lang)
		}

		url = url + "?" + strings.Join(queryParams, "&")

		break
	}

	return baseURL + url
}

// GetBody prepare request body
func (sR ShoopingHotelOffersRequest) GetBody(reqType string) io.Reader {
	return nil
}

type ShoopingHotelOffersResponse struct {
	Data         HotelData       `json:"data,omitempty"`
	Meta         Meta            `json:"meta,omitempty"`
	Dictionaries Dictionaries    `json:"dictionaries,omitempty"`
	Errors       []ErrorResponse `json:"errors,omitempty"`
}

// Decode implement Response interface
func (dR *ShoopingHotelOffersResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type ShoopingHotelsOffersResponse struct {
	Data         []HotelData     `json:"data,omitempty"`
	Meta         Meta            `json:"meta,omitempty"`
	Dictionaries Dictionaries    `json:"dictionaries,omitempty"`
	Errors       []ErrorResponse `json:"errors,omitempty"`
}

// Decode implement Response interface
func (dR *ShoopingHotelsOffersResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

// GetOffer return offer from list
func (dR ShoopingHotelsOffersResponse) GetHotel(offerNum int) HotelData {
	return dR.Data[offerNum]
}
