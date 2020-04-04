package amadeus

// FlightOffer

// RESPONSE

type FlightOffer struct {
	Type                     string            `json:"type,omitempty"`
	ID                       string            `json:"id,omitempty"`
	Source                   string            `json:"source,omitempty"`
	InstantTicketingRequired bool              `json:"instantTicketingRequired,omitempty"`
	NonHomogeneous           bool              `json:"nonHomogeneous,omitempty"`
	OneWay                   bool              `json:"oneWay,omitempty"`
	LastTicketingDate        string            `json:"lastTicketingDate,omitempty"`
	NumberOfBookableSeats    int               `json:"numberOfBookableSeats,omitempty"`
	Itineraries              []Itinerarie      `json:"itineraries,omitempty"`
	Price                    Price             `json:"price,omitempty"`
	PricingOptions           PricingOption     `json:"pricingOptions,omitempty"`
	ValidatingAirlineCodes   []string          `json:"validatingAirlineCodes,omitempty"`
	TravelerPricings         []TravelerPricing `json:"travelerPricings,omitempty"`
	PaymentCardRequired      bool              `json:"paymentCardRequired,omitempty"`
}

type Itinerarie struct {
	Duration string    `json:"duration,omitempty"`
	Segments []Segment `json:"segments,omitempty"`
}

type Segment struct {
	ID              string        `json:"id,omitempty"`
	Departure       Destination   `json:"departure,omitempty"`
	Arrival         Destination   `json:"arrival,omitempty"`
	CarrierCode     string        `json:"carrierCode,omitempty"`
	Number          string        `json:"number,omitempty"`
	Aircraft        Aircraft      `json:"aircraft,omitempty"`
	Operating       Operating     `json:"operating,omitempty"`
	Duration        string        `json:"duration,omitempty"`
	Co2Emissions    []Co2Emission `json:"co2Emissions,omitempty"`
	NumberOfStops   int           `json:"numberOfStops,omitempty"`
	BlacklistedInEU bool          `json:"blacklistedInEU,omitempty"`
}

type Destination struct {
	IataCode string `json:"iataCode,omitempty"`
	Terminal string `json:"terminal,omitempty"`
	At       string `json:"at,omitempty"`
}

type Aircraft struct {
	Code string `json:"code,omitempty"`
}
type Operating struct {
	CarrierCode string `json:"carrierCode,omitempty"`
}

type Co2Emission struct {
	Weight     string `json:"weight,omitempty"`
	WeightUnit string `json:"weightUnit,omitempty"`
	Cabin      string `json:"cabin,omitempty"`
}

type Price struct {
	Currency        string  `json:"currency,omitempty"`
	Total           string  `json:"total,omitempty"`
	Base            string  `json:"base,omitempty"`
	Fees            []Fees  `json:"fees,omitempty"`
	Taxes           []Taxes `json:"taxes,omitempty"`
	GrandTotal      string  `json:"grandTotal,omitempty"`
	BillingCurrency string  `json:"billingCurrency,omitempty"`
}

type Fees struct {
	Amount string `json:"amount,omitempty"`
	Type   string `json:"type,omitempty"`
}

type Taxes struct {
	Amount string `json:"amount,omitempty"`
	Code   string `json:"code,omitempty"`
}

type PricingOption struct {
	FareType                []string `json:"fareType,omitempty"`
	IncludedCheckedBagsOnly bool     `json:"includedCheckedBagsOnly,omitempty"`
}

type TravelerPricing struct {
	TravelerID           string                 `json:"travelerId,omitempty"`
	FareOption           string                 `json:"fareOption,omitempty"`
	TravelerType         string                 `json:"travelerType,omitempty"`
	Price                Price                  `json:"price,omitempty"`
	FareDetailsBySegment []FareDetailsBySegment `json:"fareDetailsBySegment,omitempty"`
}

type FareDetailsBySegment struct {
	SegmentID           string              `json:"segmentId,omitempty"`
	Cabin               string              `json:"cabin,omitempty"`
	FareBasis           string              `json:"fareBasis,omitempty"`
	BrandedFare         string              `json:"brandedFare,omitempty"`
	Class               string              `json:"class,omitempty"`
	IncludedCheckedBags IncludedCheckedBags `json:"includedCheckedBags,omitempty"`
}

type IncludedCheckedBags struct {
	Quantity int `json:"quantity,omitempty"`
}

// CreateOrder specific

type OrderData struct {
	ID                 string             `json:"id,omitempty"`
	Type               string             `json:"type,omitempty"`
	AssociatedRecords  []AssociatedRecord `json:"associatedRecords,omitempty"`
	FlightOffers       []FlightOffer      `json:"flightOffers,omitempty"`
	Travelers          []Traveler         `json:"travelers,omitempty"`
	TicketingAgreement TicketingAgreement `json:"ticketingAgreement,omitempty"`
	Contacts           []Contact          `json:"contacts,omitempty"`
	Remarks            Remarks            `json:"remarks,omitempty"`
	FormOfPayments     []FormOfPayments   `json:"formOfPayments,omitempty"`
	AutomatedProcess   []AutomatedProcess `json:"automatedProcess,omitempty"`
}

type AssociatedRecord struct {
	Reference        string `json:"reference,omitempty"`
	CreationDate     string `json:"creationDate,omitempty"`
	OriginSystemCode string `json:"originSystemCode,omitempty"`
	FlightOfferID    string `json:"flightOfferId,omitempty"`
}

type Traveler struct {
	ID          string          `json:"id,omitempty"`
	DateOfBirth string          `json:"dateOfBirth,omitempty"`
	Name        Name            `json:"name,omitempty"`
	Gender      string          `json:"gender,omitempty"`
	Contact     TravelerContact `json:"contact,omitempty"`
	Documents   []Document      `json:"documents,omitempty,omitempty"`
}

type Name struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

type TravelerContact struct {
	EmailAddress string  `json:"emailAddress,omitempty,omitempty"`
	Phones       []Phone `json:"phones,omitempty,omitempty"`
}

type Contact struct {
	AddresseeName AddresseeName `json:"addresseeName,omitempty"`
	CompanyName   string        `json:"companyName,omitempty"`
	Purpose       string        `json:"purpose,omitempty"`
	Phones        []Phone       `json:"phones,omitempty,omitempty"`
	EmailAddress  string        `json:"emailAddress,omitempty,omitempty"`
	Address       Address       `json:"address,omitempty,omitempty"`
}

type AddresseeName struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}
type Address struct {
	Lines       []string `json:"lines,omitempty"`
	PostalCode  string   `json:"postalCode,omitempty"`
	CountryCode string   `json:"countryCode,omitempty"`
	CityName    string   `json:"cityName,omitempty"`
	StateName   string   `json:"stateName,omitempty"`
	PostalBox   string   `json:"postalBox,omitempty"`
}

type Phone struct {
	DeviceType         string `json:"deviceType,omitempty"`
	CountryCallingCode string `json:"countryCallingCode,omitempty"`
	Number             string `json:"number,omitempty"`
}

type Document struct {
	DocumentType     string `json:"documentType,omitempty"`
	BirthPlace       string `json:"birthPlace,omitempty"`
	IssuanceLocation string `json:"issuanceLocation,omitempty"`
	IssuanceDate     string `json:"issuanceDate,omitempty"`
	Number           string `json:"number,omitempty"`
	ExpiryDate       string `json:"expiryDate,omitempty"`
	IssuanceCountry  string `json:"issuanceCountry,omitempty"`
	ValidityCountry  string `json:"validityCountry,omitempty"`
	Nationality      string `json:"nationality,omitempty"`
	Holder           bool   `json:"holder,omitempty"`
}

type FormOfPayments struct {
	Other Other `json:"other,omitempty"`
}

type Other struct {
	Method         string   `json:"method,omitempty"`
	FlightOfferIds []string `json:"flightOfferIds,omitempty"`
}

type Remarks struct {
	General []Remark `json:"general,omitempty"`
}

type Remark struct {
	SubType string `json:"subType,omitempty"`
	Text    string `json:"text,omitempty"`
}

type TicketingAgreement struct {
	Option string `json:"option,omitempty"`
	Delay  string `json:"delay,omitempty"`
}

type AutomatedProcess struct {
	Code     string `json:"code,omitempty"`
	Queue    Queue  `json:"queue,omitempty"`
	OfficeID string `json:"officeId,omitempty"`
}

type Queue struct {
	Number   string `json:"number,omitempty"`
	Category string `json:"category,omitempty"`
}
