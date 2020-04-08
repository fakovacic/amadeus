package amadeus

// Generic structs

type ErrorResponse struct {
	Code   int    `json:"code,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
	Source struct {
		Pointer string `json:"pointer,omitempty"`
		Example string `json:"example,omitempty"`
	} `json:"source,omitempty"`
	Status int `json:"status,omitempty"`
}

type Data struct {
	Type          string `json:"type,omitempty"`
	Origin        string `json:"origin,omitempty"`
	Destination   string `json:"destination,omitempty"`
	DepartureDate string `json:"departureDate,omitempty"`
	ReturnDate    string `json:"returnDate,omitempty"`
	Price         Price  `json:"price,omitempty"`
	Links         Links  `json:"links,omitempty"`
}

type Links struct {
	FlightDates  string `json:"flightDates,omitempty"`
	FlightOffers string `json:"flightOffers,omitempty"`
	Self         string `json:"self,omitempty"`
}

// FlightOffer

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
	Weight     int    `json:"weight,omitempty"`
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
