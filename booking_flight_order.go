package amadeus

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
        "fmt"
)

// BookingFlightOrder

// REQUEST

type BookingFlightOrderRequest struct {
	Data OrderData `json:"data,omitempty"`
}

// AddOffer add flight offer to request
func (sR *BookingFlightOrderRequest) AddOffer(offer FlightOffer) *BookingFlightOrderRequest {

	sR.Data.Type = "flight-order"
	sR.Data.FlightOffers = append(sR.Data.FlightOffers, offer)

	return sR
}

// AddOffers add flight offer slice to request
func (sR *BookingFlightOrderRequest) AddOffers(offers []FlightOffer) *BookingFlightOrderRequest {

	sR.Data.Type = "flight-order"
	sR.Data.FlightOffers = offers

	return sR
}

// AddTicketingAgreement add ticketing agreement
func (sR *BookingFlightOrderRequest) AddTicketingAgreement(option, delay string) *BookingFlightOrderRequest {

	sR.Data.TicketingAgreement.Option = option
	sR.Data.TicketingAgreement.Delay = delay

	return sR
}

// NewCard create card struct
func (sR BookingFlightOrderRequest) NewCard(vendorCode, cardNumber, expiryDate string) *Payment {

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
func (sR *BookingFlightOrderRequest) AddPayment(payment *Payment) *BookingFlightOrderRequest {

        sR.Data.Payments = append(sR.Data.Payments, *payment)

        return sR
}

// AddTraveler add traveler to request
func (sR BookingFlightOrderRequest) NewTraveler(firstName, lastName, gender, dateOfBirth string) *Traveler {

	// check first name
	// check last name
	// check gender
	// check date of birth

	return &Traveler{
		DateOfBirth: dateOfBirth,
		Name: Name{
			FirstName: firstName,
			LastName:  lastName,
		},
		Gender: gender,
	}
}

// AddTraveler add traveler to request
func (sR *BookingFlightOrderRequest) AddTraveler(traveler *Traveler) *BookingFlightOrderRequest {

	travelersCount := len(sR.Data.Travelers)
	travelersCount++

	traveler.ID = strconv.Itoa(travelersCount)

	sR.Data.Travelers = append(sR.Data.Travelers, *traveler)

	return sR
}


// Add carrier code to request
func (sR *BookingFlightOrderRequest) SetCarrier(carrier string) *BookingFlightOrderRequest {

        sR.Data.CarrierCodes.CarrierCode = carrier

        return sR
}

// NewContact add contact to request
func (sR BookingFlightOrderRequest) NewContact(firstName, lastName, company, purpose string) *Contact {

	// check first name
	// check last name
	// check company name
	// check purpose

	return &Contact{
		AddresseeName: AddresseeName{
			FirstName: firstName,
			LastName:  lastName,
		},
		CompanyName: company,
		Purpose:     purpose,
	}
}

// AddContact add contact to request
func (sR *BookingFlightOrderRequest) AddContact(contact *Contact) *BookingFlightOrderRequest {

	sR.Data.Contacts = append(sR.Data.Contacts, *contact)

	return sR
}

// GetURL returned key=value format for request on api
func (sR BookingFlightOrderRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := bookingFlightOrdersURL

	// add version
	switch reqType {
	case "GET":
		return baseURL + "/v1" + url + "/" + sR.Data.ID
	case "DELETE":
		return baseURL + "/v1" + url + "/" + sR.Data.ID
	case "POST":
		return baseURL + "/v1" + url
	}

	return ""
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (sR BookingFlightOrderRequest) GetBody(reqType string) io.Reader {

	switch reqType {
	case "GET":
		return nil
	case "DELETE":
		return nil
	case "POST":
		reqPayload, err := json.Marshal(sR)
                fmt.Println(strings.NewReader(string(reqPayload)))
		if err != nil {
			return nil
		}

		return strings.NewReader(string(reqPayload))
	}

	return nil
}

// RESPONSE

type BookingFlightOrderResponse struct {
	Data         OrderData       `json:"data,omitempty"`
	Dictionaries Dictionaries    `json:"dictionaries,omitempty"`
	Errors       []ErrorResponse `json:"errors,omitempty"`
}

// Decode implement Response interface
func (dR *BookingFlightOrderResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

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
        Payments           []Payment          `json:"payments,omitempty"`
        CarrierCodes       CarrierCode        `json:"operating,omitempty"`
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
        Payments    []Payment       `json:"payments,omitempty"`
}

// AddEmail add email address to traveler
func (t *Traveler) AddEmail(email string) *Traveler {

	// check email address

	t.Contact.EmailAddress = email

	return t
}

// AddMobile add contact to request
func (t *Traveler) AddMobile(countryCode, number string) *Traveler {

	// check phone number
	t.Contact.Phones = append(t.Contact.Phones, Phone{
		DeviceType:         "MOBILE",
		CountryCallingCode: countryCode,
		Number:             number,
	})

	return t
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

// AddEmail add email address to contact
func (c *Contact) AddEmail(email string) *Contact {

	// check email address

	c.EmailAddress = email

	return c
}

// AddMobile add mobile to contact
func (c *Contact) AddMobile(countryCode, number string) *Contact {

	// check phone number
	c.Phones = append(c.Phones, Phone{
		DeviceType:         "MOBILE",
		CountryCallingCode: countryCode,
		Number:             number,
	})

	return c
}

// AddAddress add address to contact
func (c *Contact) AddAddress(countryCode, cityName, postalCode string, lines ...string) *Contact {

	// check phone number
	c.Address = Address{
		Lines:       lines,
		PostalCode:  postalCode,
		CityName:    cityName,
		CountryCode: countryCode,
	}

	return c
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

type CarrierCodes struct {
	General []CarrierCode `json:"carrierCode,omitempty"`
}

type CarrierCode struct {
	SubType     string `json:"subType,omitempty"`
	CarrierCode string `json:"carrierCode,omitempty"`
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
