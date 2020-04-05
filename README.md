# Go client for Amadeus REST API

## Coverage

### Authentification

 * POST /v1/security/oauth2/token
 * GET  /v1/security/oauth2/token/{token}

### AIR

#### Search & Shooping

 * POST /v2/shopping/flight-offers
 * POST /v1/shopping/flight-offers/pricing
 * GET  /v1/shopping/flight-destinations

#### Booking

 * POST /v1/booking/flight-orders
 * GET  /v1/booking/flight-orders/{orderID}

## Missing endpoints

### AIR

#### Search & Shooping

 * GET  /v1/shopping/flight-offers
 * GET  /v2/shopping/flight-offers
 * GET  /v1/shopping/flight-dates
 * POST /v1/shopping/seatmaps

#### Travel Insights

 * GET /v1/travel/analytics/air-traffic/traveled
 * GET /v1/travel/analytics/air-traffic/booked
 * GET /v1/travel/analytics/air-traffic/busiest-period

#### Utilities

 * GET /v1/reference-data/locations/airports
 * GET /v1/reference-data/locations
 * GET /v1/reference-data/locations/{id}
 * GET /v2/reference-data/urls/checkin-links
 * GET /v1/reference-data/airlines

#### Artificial Inteligence

 * POST /v1/shopping/flight-offers/prediction
 * GET  /v1/travel/predictions/flight-delay
 * GET  /v1/airport/predictions/on-time


### Hotel

#### Search & Shooping

 * GET /v2/shopping/hotel-offers
 * GET /v2/shopping/hotel-offers/by-hotel
 * GET /v2/shopping/hotel-offers/{id}

#### Travel Insights

 * GET /v2/e-reputation/hotel-sentiments

#### Booking

 * GET /v1/booking/hotel-bookings


## Usage

- get package 

```
go get https://github.com/fakovacic/amadeus-golang
```

### Flight order creation

- init client

```

amadeus, err := New(os.Getenv("AMADEUS_CLIENT_ID"), os.Getenv("AMADEUS_CLIENT_SECRET"), os.Getenv("AMADEUS_ENV"))

```

- send search offers request to retrive flight offers

```

s := FlightOffersSearchRequest{
    CurrencyCode: "EUR",
    OriginDestinations: []OriginDestination{
        OriginDestination{
            ID:              "1",
            OriginCode:      "LON",
            DestinationCode: "PAR",
            DepartureDateTimeRange: TimeRange{
                Date: time.Now().AddDate(0, 5, 0).Format("2006-01-02"),
            },
        },
    },
    Travelers: []Travelers{
        Travelers{
            ID:           "1",
            TravelerType: "ADULT",
        },
    },
    Sources: []string{
        "GDS",
    },
}

offersResp, err := amadeus.FlightOffers(s)
```

- construct flight search request

```
s := NewSearchRequest("EUR", "GDS")
s.Oneway("LON", "PAR", time.Now().AddDate(0, 5, 0).Format("2006-01-02"))
s.AddTravelers(1, 0, 0)

offersResp, err := amadeus.FlightOffers(s)
```

- check for pricing for first option in offers response


```
p := FlightOffersPriceRequest{
    Data: PricingData{
        Type: "flight-offers-pricing",
        FlightOffers: []FlightOffer{
            offersResp.Data[0],
        },
    },
}

pricingResp, err := amadeus.FlightPricing(p)
```

- create booking/order for priced option

```
o := FlightCreateOrdersRequest{
    Data: OrderData{
        Type: "flight-order",
        FlightOffers: []FlightOffer{
            pricingResp.Data.FlightOffers[0],
        },
        Travelers: []Traveler{
            Traveler{
                ID:          "1",
                DateOfBirth: "1980-02-15",
                Name: Name{
                    FirstName: "Foo",
                    LastName:  "Bar",
                },
                Gender: "MALE",
                Contact: TravelerContact{
                    EmailAddress: "foo@bar.com",
                    Phones: []Phone{
                        Phone{
                            DeviceType:         "MOBILE",
                            CountryCallingCode: "33",
                            Number:             "480080072",
                        },
                    },
                },
            },
        },
        TicketingAgreement: TicketingAgreement{
            Option: "DELAY_TO_CANCEL",
            Delay:  "6D",
        },
        Contacts: []Contact{
            Contact{
                AddresseeName: AddresseeName{
                    FirstName: "Foo",
                    LastName:  "Bar",
                },
                CompanyName: "TESTING",
                Purpose:     "STANDARD",
                Phones: []Phone{
                    Phone{
                        DeviceType:         "MOBILE",
                        CountryCallingCode: "33",
                        Number:             "480080072",
                    },
                },
                EmailAddress: "foo@bar.com",
                Address: Address{
                    Lines: []string{
                        "Street 25",
                    },
                    PostalCode:  "45453",
                    CityName:    "Madrid",
                    CountryCode: "ES",
                },
            },
        },
    },
}

orderResp, err := amadeus.FlightCreateOrder(o)

```



