# Amadeus API Golang SDK

> 

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