package booking

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/structs"
)

type HotelBookingsRequest struct {
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
func (sR *HotelBookingsRequest) SetOfferID(offerID string) *HotelBookingsRequest {

	sR.Data.OfferID = offerID

	return sR
}

// NewGuest add contact to request
func (sR HotelBookingsRequest) NewGuest(title, firstName, lastName, email, phone string) *Guest {

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
func (sR *HotelBookingsRequest) AddGuest(guest *Guest) *HotelBookingsRequest {

	sR.Data.Guests = append(sR.Data.Guests, *guest)

	return sR
}

// NewCard create card struct
func (sR HotelBookingsRequest) NewCard(vendorCode, cardNumber, expiryDate string) *Payment {

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
func (sR *HotelBookingsRequest) AddPayment(payment *Payment) *HotelBookingsRequest {

	sR.Data.Payments = append(sR.Data.Payments, *payment)

	return sR
}

// SetParam set params
func (dR *HotelBookingsRequest) SetParam(key, value string) {
	return
}

// GetURL returned key=value format for request on api
func (sR HotelBookingsRequest) GetURL(reqType string) string {

	// set request url
	url := hotelBookingsURL

	// add version
	switch reqType {
	case "POST":
		return "/v1" + url
	}

	return ""
}

// GetBody prepare struct to request
func (sR HotelBookingsRequest) GetBody(reqType string) io.Reader {

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

type HotelBookingsResponse struct {
	Data     []HotelBookingsResponseData `json:"data,omitempty"`
	Warnings []structs.Warnings          `json:"warnings,omitempty"`
}

// Decode implement Response interface
func (dR *HotelBookingsResponse) Decode(rsp []byte) error {

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
