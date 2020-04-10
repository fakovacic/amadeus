package amadeus

import (
	"encoding/json"
	"io"
	"strings"
)

type BookingHotelBookingsRequest struct {
	Data HotelBookingsData `json:"data,omitempty"`
}

type HotelBookingsData struct {
	OfferID  string    `json:"offerId,omitempty"`
	Guests   []Guest   `json:"guests,omitempty"`
	Payments []Payment `json:"payments,omitempty"`
}

type Guest struct {
	Name    GuestName    `json:"name,omitempty"`
	Contact GuestContact `json:"contact,omitempty"`
}

type GuestName struct {
	Title     string `json:"title,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}
type GuestContact struct {
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
}

type Payment struct {
	Method string `json:"method,omitempty"`
	Card   Card   `json:"card,omitempty"`
}

type Card struct {
	VendorCode string `json:"vendorCode,omitempty"`
	CardNumber string `json:"cardNumber,omitempty"`
	ExpiryDate string `json:"expiryDate,omitempty"`
}

// SetOfferID add offer id
func (sR *BookingHotelBookingsRequest) SetOfferID(offerID string) *BookingHotelBookingsRequest {

	sR.Data.OfferID = offerID

	return sR
}

// NewGuest add contact to request
func (sR BookingHotelBookingsRequest) NewGuest(title, firstName, lastName, email, phone string) *Guest {

	// check title
	// check first name
	// check last name
	// check email
	// check phone

	return &Guest{
		Name: GuestName{
			Title:     title,
			FirstName: firstName,
			LastName:  lastName,
		},
		Contact: GuestContact{
			Email: email,
			Phone: phone,
		},
	}
}

// AddGuest add guest to request
func (sR *BookingHotelBookingsRequest) AddGuest(guest *Guest) *BookingHotelBookingsRequest {

	sR.Data.Guests = append(sR.Data.Guests, *guest)

	return sR
}

// NewCard create card struct
func (sR BookingHotelBookingsRequest) NewCard(vendorCode, cardNumber, expiryDate string) *Payment {

	// check vendorCode
	// check cardNumber
	// check expiryDate

	return &Payment{
		Method: "creditCard",
		Card: Card{
			VendorCode: vendorCode,
			CardNumber: cardNumber,
			ExpiryDate: expiryDate,
		},
	}
}

// AddPayment add payment to request
func (sR *BookingHotelBookingsRequest) AddPayment(payment *Payment) *BookingHotelBookingsRequest {

	sR.Data.Payments = append(sR.Data.Payments, *payment)

	return sR
}

// GetURL returned key=value format for request on api
func (sR BookingHotelBookingsRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := bookingHotelBookingsURL

	// add version
	switch reqType {
	case "POST":
		return baseURL + "/v1" + url
	}

	return ""
}

// GetBody prepare struct to request
func (sR BookingHotelBookingsRequest) GetBody(reqType string) io.Reader {

	switch reqType {
	case "POST":
		reqPayload, err := json.Marshal(sR)
		if err != nil {
			return nil
		}

		return strings.NewReader(string(reqPayload))
	}

	return nil
}

type BookingHotelBookingsResponse struct {
	Data     []HotelBookingsResponseData `json:"data,omitempty"`
	Warnings []Warnings                  `json:"warnings,omitempty"`
	Errors   []ErrorResponse             `json:"errors,omitempty"`
}

// Decode implement Response interface
func (dR *BookingHotelBookingsResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type HotelBookingsResponseData struct {
	Type                   string              `json:"type,omitempty"`
	ID                     string              `json:"id,omitempty"`
	ProviderConfirmationID string              `json:"providerConfirmationId,omitempty"`
	AssociatedRecords      []AssociatedRecords `json:"associatedRecords,omitempty"`
	Self                   string              `json:"self,omitempty"`
}

type AssociatedRecords struct {
	Reference        string `json:"reference,omitempty"`
	OriginSystemCode string `json:"originSystemCode,omitempty"`
}
