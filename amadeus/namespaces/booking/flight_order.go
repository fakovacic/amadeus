package booking

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"

	"github.com/fakovacic/amadeus-golang/amadeus/namespaces/structs"
)

// FlightOrder

// REQUEST

type FlightOrderRequest struct {
	Data OrderData `json:"data,omitempty"`
}

// AddOffer add flight offer to request
func (sR *FlightOrderRequest) AddOffer(offer structs.FlightOffer) *FlightOrderRequest {

	sR.Data.Type = "flight-order"
	sR.Data.FlightOffers = append(sR.Data.FlightOffers, offer)

	return sR
}

// AddOffers add flight offer slice to request
func (sR *FlightOrderRequest) AddOffers(offers []structs.FlightOffer) *FlightOrderRequest {

	sR.Data.Type = "flight-order"
	sR.Data.FlightOffers = offers

	return sR
}

// AddTicketingAgreement add ticketing agreement
func (sR *FlightOrderRequest) AddTicketingAgreement(option, delay string) *FlightOrderRequest {

	sR.Data.TicketingAgreement.Option = option
	sR.Data.TicketingAgreement.Delay = delay

	return sR
}

// AddTraveler add traveler to request
func (sR FlightOrderRequest) NewTraveler(firstName, lastName, gender, dateOfBirth string) *Traveler {

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
func (sR *FlightOrderRequest) AddTraveler(traveler *Traveler) *FlightOrderRequest {

	travelersCount := len(sR.Data.Travelers)
	travelersCount++

	traveler.ID = strconv.Itoa(travelersCount)

	sR.Data.Travelers = append(sR.Data.Travelers, *traveler)

	return sR
}

// NewContact add contact to request
func (sR FlightOrderRequest) NewContact(firstName, lastName, company, purpose string) *Contact {

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
func (sR *FlightOrderRequest) AddContact(contact *Contact) *FlightOrderRequest {

	sR.Data.Contacts = append(sR.Data.Contacts, *contact)

	return sR
}

// SetParam set params
func (dR *FlightOrderRequest) SetParam(key, value string) {
	return
}

// GetURL returned key=value format for request on api
func (sR FlightOrderRequest) GetURL(reqType string) string {

	// set request url
	url := flightOrdersURL

	// add version
	switch reqType {
	case "GET":
		return "/v1" + url + "/" + sR.Data.ID
	case "DELETE":
		return "/v1" + url + "/" + sR.Data.ID
	case "POST":
		return "/v1" + url
	}

	return ""
}

// GetBody prepare struct values to slice
// returned key=value format for request on api
func (sR FlightOrderRequest) GetBody(reqType string) io.Reader {

	switch reqType {
	case "GET":
		return nil
	case "DELETE":
		return nil
	case "POST":
		reqPayload, err := json.Marshal(sR)
		if err != nil {
			return nil
		}

		return strings.NewReader(string(reqPayload))
	}

	return nil
}

// RESPONSE

type FlightOrderResponse struct {
	Data         OrderData            `json:"data,omitempty"`
	Dictionaries structs.Dictionaries `json:"dictionaries,omitempty"`
}

// Decode implement Response interface
func (dR *FlightOrderResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type OrderData struct {
	ID                 string                `json:"id,omitempty"`
	Type               string                `json:"type,omitempty"`
	AssociatedRecords  []AssociatedRecord    `json:"associatedRecords,omitempty"`
	FlightOffers       []structs.FlightOffer `json:"flightOffers,omitempty"`
	Travelers          []Traveler            `json:"travelers,omitempty"`
	TicketingAgreement TicketingAgreement    `json:"ticketingAgreement,omitempty"`
	Contacts           []Contact             `json:"contacts,omitempty"`
	Remarks            Remarks               `json:"remarks,omitempty"`
	FormOfPayments     []FormOfPayments      `json:"formOfPayments,omitempty"`
	AutomatedProcess   []AutomatedProcess    `json:"automatedProcess,omitempty"`
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
