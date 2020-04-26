package structs

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

type Warnings struct {
	Status int    `json:"status"`
	Code   int    `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Source Source `json:"source"`
}

type Source struct {
	Pointer   string `json:"pointer"`
	Parameter string `json:"parameter"`
	Example   string `json:"example"`
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

type Meta struct {
	Count    int      `json:"count,omitempty"`
	Currency string   `json:"currency,omitempty"`
	Links    Links    `json:"links,omitempty"`
	Defaults Defaults `json:"defaults,omitempty"`
}

type Defaults struct {
	DepartureDate string `json:"departureDate,omitempty"`
	OneWay        bool   `json:"oneWay,omitempty"`
	Duration      string `json:"duration,omitempty"`
	NonStop       bool   `json:"nonStop,omitempty"`
	ViewBy        string `json:"viewBy,omitempty"`
}
