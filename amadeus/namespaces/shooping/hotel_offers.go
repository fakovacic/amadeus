package shooping

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/structs"
)

const (

	/////////////
	// 	HOTEL  //
	/////////////

	// Hotel Search
	// search for hotel offers on destination
	hotelOffersURL = "/shopping/hotel-offers"
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

type HotelOffersRequest struct {
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
func (sR *HotelOffersRequest) SetOfferID(offerID string) *HotelOffersRequest {

	sR.OfferID = offerID

	return sR
}

// SetHotelID set hotel id
func (sR *HotelOffersRequest) SetHotelID(hotelID string) *HotelOffersRequest {

	sR.HotelID = hotelID

	return sR
}

// SetCityCode set city code
func (sR *HotelOffersRequest) SetCityCode(cityCode string) *HotelOffersRequest {

	sR.CityCode = cityCode

	return sR
}

// SetCheckInDate set checkin date
func (sR *HotelOffersRequest) SetCheckInDate(checkInDate string) *HotelOffersRequest {

	// check date

	sR.CheckInDate = checkInDate

	return sR
}

// SetCheckOutDate set checkout date
func (sR *HotelOffersRequest) SetCheckOutDate(checkOutDate string) *HotelOffersRequest {

	// check date

	sR.CheckOutDate = checkOutDate

	return sR
}

// SetAdults set adults
func (sR *HotelOffersRequest) SetAdults(adults int) *HotelOffersRequest {

	// check date

	sR.Adults = adults

	return sR
}

// AddChildAges add childAges
func (sR *HotelOffersRequest) AddChildAges(childAges ...int) *HotelOffersRequest {

	sR.ChildAges = childAges

	return sR
}

// SetHotelName set hotel name
func (sR *HotelOffersRequest) SetHotelName(hotelName string) *HotelOffersRequest {

	sR.HotelName = hotelName

	return sR
}

// SetLatitude set latitude
func (sR *HotelOffersRequest) SetLatitude(latitude float64) *HotelOffersRequest {

	sR.Latitude = latitude

	return sR
}

// SetLongitude set longitude
func (sR *HotelOffersRequest) SetLongitude(longitude float64) *HotelOffersRequest {

	sR.Longitude = longitude

	return sR
}

// SetRadius set radius
func (sR *HotelOffersRequest) SetRadius(radius int) *HotelOffersRequest {

	sR.Radius = radius

	return sR
}

// SetRadiusUnit set radiusUnit | KM or MILE
func (sR *HotelOffersRequest) SetRadiusUnit(radiusUnit RadiusUnit) *HotelOffersRequest {

	sR.RadiusUnit = radiusUnit

	return sR
}

// AddHotelIDs add hotelIDs
func (sR *HotelOffersRequest) AddHotelIDs(hotelIDs ...string) *HotelOffersRequest {

	sR.HotelIDs = hotelIDs

	return sR
}

// AddChains add hotel chains filter
func (sR *HotelOffersRequest) AddChains(chains ...string) *HotelOffersRequest {

	sR.Chains = chains

	return sR
}

// AddRateCodes add rateCodes
func (sR *HotelOffersRequest) AddRateCodes(rateCodes ...string) *HotelOffersRequest {

	sR.RateCodes = rateCodes

	return sR
}

// AddAmenities add amenities
func (sR *HotelOffersRequest) AddAmenities(amenities ...string) *HotelOffersRequest {

	sR.Amenities = amenities

	return sR
}

// AddRatings add ratings
func (sR *HotelOffersRequest) AddRatings(ratings ...int) *HotelOffersRequest {

	sR.Ratings = ratings

	return sR
}

// SetPriceRange set priceRange
func (sR *HotelOffersRequest) SetPriceRange(priceRange string) *HotelOffersRequest {

	sR.PriceRange = priceRange

	return sR
}

// SetCurrency set currency
func (sR *HotelOffersRequest) SetCurrency(currency string) *HotelOffersRequest {

	sR.Currency = currency

	return sR
}

// SetPaymentPolicy set payment policy
func (sR *HotelOffersRequest) SetPaymentPolicy(paymentPolicy HotelPaymentPolicy) *HotelOffersRequest {

	sR.PaymentPolicy = paymentPolicy

	return sR
}

// SetBoardType set board type
func (sR *HotelOffersRequest) SetBoardType(BoardType HotelBoardType) *HotelOffersRequest {

	sR.BoardType = BoardType

	return sR
}

// IsIncludeClosed include closed
func (sR *HotelOffersRequest) IsIncludeClosed(includeClosed bool) *HotelOffersRequest {

	if includeClosed {
		sR.IncludeClosed = true
	}

	return sR
}

// IsBestRateOnly best rate only
func (sR *HotelOffersRequest) IsBestRateOnly(bestRateOnly bool) *HotelOffersRequest {

	if bestRateOnly {
		sR.BestRateOnly = true
	}

	return sR
}

// SetView set view
func (sR *HotelOffersRequest) SetView(view HotelView) *HotelOffersRequest {

	sR.View = view

	return sR
}

// SetSort set sort
func (sR *HotelOffersRequest) SetSort(sort HotelSort) *HotelOffersRequest {

	sR.Sort = sort

	return sR
}

// SetLang set lang
func (sR *HotelOffersRequest) SetLang(lang string) *HotelOffersRequest {

	sR.Lang = lang

	return sR
}

// Helper functions

// SetGeo set latitude & longitude
func (sR *HotelOffersRequest) SetGeo(latitude, longitude float64) *HotelOffersRequest {

	sR.SetLatitude(latitude)
	sR.SetLongitude(longitude)

	return sR
}

// JoinInts join int slice to string
func (sR HotelOffersRequest) JoinInts(intSlice []int) string {

	valuesText := []string{}

	for i := range intSlice {
		number := intSlice[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	return strings.Join(valuesText, ",")
}

// SetParam set params
func (dR *HotelOffersRequest) SetParam(key, value string) {
	return
}

// GetURL returned key=value format for request on api
func (sR HotelOffersRequest) GetURL(reqType string) string {

	// set request url
	url := hotelOffersURL

	// add version
	switch reqType {
	case "GET":

		url = "/v2" + url

		// define query params
		queryParams := []string{}

		// check request values
		if sR.OfferID != "" {

			return url + "/" + sR.OfferID

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

	return url
}

// GetBody prepare request body
func (sR HotelOffersRequest) GetBody(reqType string) io.Reader {
	return nil
}

type HotelOffersResponse struct {
	Data         structs.HotelData    `json:"data,omitempty"`
	Meta         structs.Meta         `json:"meta,omitempty"`
	Dictionaries structs.Dictionaries `json:"dictionaries,omitempty"`
}

// Decode implement Response interface
func (dR *HotelOffersResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type HotelsOffersResponse struct {
	Data         []structs.HotelData  `json:"data,omitempty"`
	Meta         structs.Meta         `json:"meta,omitempty"`
	Dictionaries structs.Dictionaries `json:"dictionaries,omitempty"`
}

// Decode implement Response interface
func (dR *HotelsOffersResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

// GetOffer return offer from list
func (dR HotelsOffersResponse) GetHotel(offerNum int) structs.HotelData {
	return dR.Data[offerNum]
}
