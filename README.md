# Go client for Amadeus REST API

- still work in progress

## Coverage

### Authentification

 * POST /v1/security/oauth2/token
 * GET  /v1/security/oauth2/token/{token}

### AIR

#### Search & Shooping

 * GET  /v2/shopping/flight-offers
 * POST /v2/shopping/flight-offers

 * GET  /v1/shopping/flight-destinations
 * GET  /v1/shopping/flight-dates

 * POST /v1/shopping/flight-offers/pricing

#### Booking

 * POST /v1/booking/flight-orders
 * GET  /v1/booking/flight-orders/{orderID}
 
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

## Missing endpoints

### AIR

#### Search & Shooping

 * POST /v1/shopping/seatmaps

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
go get https://github.com/fakovacic/amadeus
```

### Flight order creation

- init client

```

    amadeus, err := New(
        os.Getenv("AMADEUS_CLIENT_ID"),
        os.Getenv("AMADEUS_CLIENT_SECRET"),
        os.Getenv("AMADEUS_ENV"),
    )
    if err != nil {
        t.Fatal("not expected error while creating client", err)
    }

```

- send search offers request to retrive flight offers

```

    // get offer flights request&response
    offerReq, offerResp, err := amadeus.NewRequest(ShoppingFlightOffers)

    // set offer flights params
    offerReq.(*ShoppingFlightOffersRequest).
        SetCurrency("EUR").
        SetSources("GDS").
        Oneway(
            "FRA",
            "BER",
            time.Now().AddDate(0, 5, 0).Format("2006-01-02"),
        ).
        AddTravelers(1, 0, 0)

    // send request
    err = amadeus.Do(offerReq, &offerResp, "POST")

    // get response
    offerRespData := offerResp.(*ShoppingFlightOffersResponse)

```

- check pricing for first option in offers response


```
    // get pricing request&response
    pricingReq, pricingResp, err := amadeus.NewRequest(ShoppingFlightPricing)

    // add offer from flight offers response
    pricingReq.(*ShoppingFlightPricingRequest).AddOffer(
        offerRespData.GetOffer(0),
    )

    // send request
    err = amadeus.Do(pricingReq, &pricingResp, "POST")

    // get response
    pricingRespData := pricingResp.(*ShoppingFlightPricingResponse)

```

- create booking order for priced option

```
    // get booking request
    bookingReq, bookingResp, err := amadeus.NewRequest(BookingFlightOrder)

    // add offer from flight pricing response
    bookingReq.(*BookingFlightOrderRequest).AddOffers(
        pricingRespData.GetOffers(),
    ).AddTicketingAgreement("DELAY_TO_CANCEL", "6D")

    // add traveler
    bookingReq.(*BookingFlightOrderRequest).AddTraveler(
        bookingReq.(*BookingFlightOrderRequest).
            NewTraveler(
                "Foo", "Bar", "MALE", "1990-02-15",
            ).
            AddEmail("foo@bar.com").
            AddMobile("33", "480080072"),
    )

    // add contact
    bookingReq.(*BookingFlightOrderRequest).AddContact(
        bookingReq.(*BookingFlightOrderRequest).
            NewContact(
                "Foo", "Bar", "TESTING", "STANDARD",
            ).
            AddEmail("foo@bar.com").
            AddMobile("33", "480080072").
            AddAddress("ES", "Madrid", "45453", "Street 25"),
    )

    // send request
    err = amadeus.Do(bookingReq, &bookingResp, "POST")

    // get flight booking response
    bookingRespData := bookingResp.(*BookingFlightOrderResponse)

```



