package amadeus

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
        Operating                []Operating       `json:"operating,omitempty"`
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
	TravelerID             string                 `json:"travelerId,omitempty"`
	FareOption             string                 `json:"fareOption,omitempty"`
	TravelerType           string                 `json:"travelerType,omitempty"`
	SeatAvailabilityStatus string                 `json:"seatAvailabilityStatus,omitempty"`
	Price                  Price                  `json:"price,omitempty"`
	FareDetailsBySegment   []FareDetailsBySegment `json:"fareDetailsBySegment,omitempty"`
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

// Dictionaries
type Dictionaries struct {
	Carriers           map[string]string            `json:"carriers,omitempty"`
	Currencies         map[string]string            `json:"currencies,omitempty"`
	Aircrafts          map[string]string            `json:"aircraft,omitempty"`
	SeatCharacteristic map[string]string            `json:"seatCharacteristic,omitempty"`
	Facility           map[string]string            `json:"facility,omitempty"`
	Locations          map[string]map[string]string `json:"locations,omitempty"`
	// hotel offers
	CurrencyConversionLookupRates map[string]string `json:"currencyConversionLookupRates,omitempty"`
}
