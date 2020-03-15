package amadeusgolang

// FlightOffersSearchRequest

// REQUEST

type FlightOffersSearchRequest struct {
	CurrencyCode       string               `json:"currencyCode"`
	OriginDestinations []OriginDestinations `json:"originDestinations"`
	Travelers          []Travelers          `json:"travelers"`
	Sources            []string             `json:"sources"`
	SearchCriteria     SearchCriteria       `json:"searchCriteria"`
}
type DepartureDateTimeRange struct {
	Date string `json:"date"`
	Time string `json:"time"`
}
type OriginDestinations struct {
	ID                      string                 `json:"id"`
	OriginLocationCode      string                 `json:"originLocationCode"`
	DestinationLocationCode string                 `json:"destinationLocationCode"`
	DepartureDateTimeRange  DepartureDateTimeRange `json:"departureDateTimeRange"`
}
type Travelers struct {
	ID           string `json:"id"`
	TravelerType string `json:"travelerType"`
}
type CabinRestrictions struct {
	Cabin                string   `json:"cabin"`
	Coverage             string   `json:"coverage"`
	OriginDestinationIds []string `json:"originDestinationIds"`
}
type CarrierRestrictions struct {
	ExcludedCarrierCodes []string `json:"excludedCarrierCodes"`
}
type FlightFilters struct {
	CabinRestrictions   []CabinRestrictions `json:"cabinRestrictions"`
	CarrierRestrictions CarrierRestrictions `json:"carrierRestrictions"`
}
type SearchCriteria struct {
	MaxFlightOffers int           `json:"maxFlightOffers"`
	FlightFilters   FlightFilters `json:"flightFilters"`
}

// RESPONSE

type FlightOffersSearchResponse struct {
	Meta Meta          `json:"meta"`
	Data []FlightOffer `json:"data"`
}

type Meta struct {
	Count int `json:"count"`
}

// FlightOffersPrice

// REQUEST

type FlightOffersPriceRequest struct {
	Data FlightOffer `json:"data"`
}

// RESPONSE

type FlightOffersPriceResponse struct {
	Data PricingData `json:"data"`
}

type PricingData struct {
	Type         string        `json:"type"`
	FlightOffers []FlightOffer `json:"flightOffers"`
}

// FlightCreateOrders

// REQUEST

type FlightCreateOrdersRequest struct {
	Data OrderData `json:"data"`
}

// RESPONSE

type FlightCreateOrdersResponse struct {
	Data OrderData `json:"data"`
}

type FlightOffer struct {
	Type                     string            `json:"type"`
	ID                       string            `json:"id"`
	Source                   string            `json:"source"`
	InstantTicketingRequired bool              `json:"instantTicketingRequired"`
	NonHomogeneous           bool              `json:"nonHomogeneous"`
	OneWay                   bool              `json:"oneWay"`
	LastTicketingDate        string            `json:"lastTicketingDate"`
	NumberOfBookableSeats    int               `json:"numberOfBookableSeats"`
	Itineraries              []Itinerarie      `json:"itineraries"`
	Price                    Price             `json:"price"`
	PricingOptions           PricingOption     `json:"pricingOptions"`
	ValidatingAirlineCodes   []string          `json:"validatingAirlineCodes"`
	TravelerPricings         []TravelerPricing `json:"travelerPricings"`
	PaymentCardRequired      bool              `json:"paymentCardRequired"`
}

type Itinerarie struct {
	Duration string    `json:"duration"`
	Segments []Segment `json:"segments"`
}

type Segment struct {
	ID              string        `json:"id"`
	Departure       Destination   `json:"departure"`
	Arrival         Destination   `json:"arrival"`
	CarrierCode     string        `json:"carrierCode"`
	Number          string        `json:"number"`
	Aircraft        Aircraft      `json:"aircraft"`
	Operating       Operating     `json:"operating"`
	Duration        string        `json:"duration"`
	Co2Emissions    []Co2Emission `json:"co2Emissions"`
	NumberOfStops   int           `json:"numberOfStops"`
	BlacklistedInEU bool          `json:"blacklistedInEU"`
}

type Destination struct {
	IataCode string `json:"iataCode"`
	Terminal string `json:"terminal"`
	At       string `json:"at"`
}

type Aircraft struct {
	Code string `json:"code"`
}
type Operating struct {
	CarrierCode string `json:"carrierCode"`
}

type Co2Emission struct {
	Weight     string `json:"weight"`
	WeightUnit string `json:"weightUnit"`
	Cabin      string `json:"cabin"`
}

type Price struct {
	Currency        string  `json:"currency"`
	Total           string  `json:"total"`
	Base            string  `json:"base"`
	Fees            []Fees  `json:"fees"`
	Taxes           []Taxes `json:"taxes"`
	GrandTotal      string  `json:"grandTotal"`
	BillingCurrency string  `json:"billingCurrency"`
}

type Fees struct {
	Amount string `json:"amount"`
	Type   string `json:"type"`
}

type Taxes struct {
	Amount string `json:"amount"`
	Code   string `json:"code"`
}

type PricingOption struct {
	FareType                []string `json:"fareType"`
	IncludedCheckedBagsOnly bool     `json:"includedCheckedBagsOnly"`
}

type TravelerPricing struct {
	TravelerID           string                 `json:"travelerId"`
	FareOption           string                 `json:"fareOption"`
	TravelerType         string                 `json:"travelerType"`
	Price                Price                  `json:"price"`
	FareDetailsBySegment []FareDetailsBySegment `json:"fareDetailsBySegment"`
}

type FareDetailsBySegment struct {
	SegmentID           string              `json:"segmentId"`
	Cabin               string              `json:"cabin"`
	FareBasis           string              `json:"fareBasis"`
	BrandedFare         string              `json:"brandedFare"`
	Class               string              `json:"class"`
	IncludedCheckedBags IncludedCheckedBags `json:"includedCheckedBags"`
}

type IncludedCheckedBags struct {
	Quantity int `json:"quantity"`
}

// CreateOrder specific

type OrderData struct {
	ID                 string             `json:"id"`
	Type               string             `json:"type"`
	AssociatedRecords  []AssociatedRecord `json:"associatedRecords"`
	FlightOffers       []FlightOffer      `json:"flightOffers"`
	Travelers          []Traveler         `json:"travelers"`
	Contacts           []Contact          `json:"contacts"`
	Remarks            Remarks            `json:"remarks"`
	FormOfPayments     []FormOfPayments   `json:"formOfPayments"`
	TicketingAgreement TicketingAgreement `json:"ticketingAgreement"`
	AutomatedProcess   []AutomatedProcess `json:"automatedProcess"`
}

type AssociatedRecord struct {
	Reference        string `json:"reference"`
	CreationDate     string `json:"creationDate"`
	OriginSystemCode string `json:"originSystemCode"`
	FlightOfferID    string `json:"flightOfferId"`
}

type Traveler struct {
	ID          string     `json:"id"`
	DateOfBirth string     `json:"dateOfBirth"`
	Name        Name       `json:"name"`
	Gender      string     `json:"gender"`
	Contact     Contact    `json:"contact"`
	Documents   []Document `json:"documents,omitempty"`
}

type Name struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Contact struct {
	AddresseeName AddresseeName `json:"addresseeName"`
	Address       Address       `json:"address,omitempty"`
	Purpose       string        `json:"purpose"`
	CompanyName   string        `json:"companyName"`
	Phones        []Phone       `json:"phones,omitempty"`
	EmailAddress  string        `json:"emailAddress,omitempty"`
}

type AddresseeName struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type Address struct {
	Lines       []string `json:"lines"`
	PostalCode  string   `json:"postalCode"`
	CountryCode string   `json:"countryCode"`
	CityName    string   `json:"cityName"`
	StateName   string   `json:"stateName"`
	PostalBox   string   `json:"postalBox"`
}

type Phone struct {
	DeviceType         string `json:"deviceType"`
	CountryCallingCode string `json:"countryCallingCode"`
	Number             string `json:"number"`
}

type Document struct {
	DocumentType     string `json:"documentType"`
	BirthPlace       string `json:"birthPlace"`
	IssuanceLocation string `json:"issuanceLocation"`
	IssuanceDate     string `json:"issuanceDate"`
	Number           string `json:"number"`
	ExpiryDate       string `json:"expiryDate"`
	IssuanceCountry  string `json:"issuanceCountry"`
	ValidityCountry  string `json:"validityCountry"`
	Nationality      string `json:"nationality"`
	Holder           bool   `json:"holder"`
}

type FormOfPayments struct {
	Other Other `json:"other"`
}

type Other struct {
	Method         string   `json:"method"`
	FlightOfferIds []string `json:"flightOfferIds"`
}

type Remarks struct {
	General []Remark `json:"general"`
}

type Remark struct {
	SubType string `json:"subType"`
	Text    string `json:"text"`
}

type TicketingAgreement struct {
	Option string `json:"option"`
	Delay  string `json:"delay"`
}

type AutomatedProcess struct {
	Code     string `json:"code"`
	Queue    Queue  `json:"queue"`
	OfficeID string `json:"officeId"`
}

type Queue struct {
	Number   string `json:"number"`
	Category string `json:"category"`
}
