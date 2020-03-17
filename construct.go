package amadeusgolang

import "strconv"

//FlightOffersSearchRequest

func NewSearchRequest(currency string, sources ...string) FlightOffersSearchRequest {

	var sR FlightOffersSearchRequest

	sR.CurrencyCode = currency

	if len(sources) != 0 {
		sR.Sources = sources
	}

	return sR
}

func (sR *FlightOffersSearchRequest) Oneway(origin, destination, departureDate string) {

	sR.OriginDestinations = append(sR.OriginDestinations, OriginDestination{
		ID:              "1",
		OriginCode:      origin,
		DestinationCode: destination,
		DepartureDateTimeRange: TimeRange{
			Date: departureDate,
		},
	})

}

func (sR *FlightOffersSearchRequest) Return(origin, destination, departureDate, returnDate string) {

	sR.OriginDestinations = append(sR.OriginDestinations, OriginDestination{
		ID:              "1",
		OriginCode:      origin,
		DestinationCode: destination,
		DepartureDateTimeRange: TimeRange{
			Date: departureDate,
		},
	})

	sR.OriginDestinations = append(sR.OriginDestinations, OriginDestination{
		ID:              "2",
		OriginCode:      destination,
		DestinationCode: origin,
		DepartureDateTimeRange: TimeRange{
			Date: returnDate,
		},
	})

}

func (sR *FlightOffersSearchRequest) Multi(origin, destination, departureDate, returnDate string) {
	//TODO
}

func (sR *FlightOffersSearchRequest) AddTravelers(adult, child, infant int) {

	paxCount := 1

	if adult != 0 {

		for i := 0; i <= adult; i++ {

			sR.Travelers = append(sR.Travelers, Travelers{
				ID:           strconv.Itoa(paxCount),
				TravelerType: "ADULT",
			})

			paxCount++
		}

	}

	if child != 0 {

		for i := 0; i <= child; i++ {

			sR.Travelers = append(sR.Travelers, Travelers{
				ID:           strconv.Itoa(paxCount),
				TravelerType: "CHILD",
			})

			paxCount++
		}

	}

	if infant != 0 {

		for i := 0; i <= infant; i++ {

			sR.Travelers = append(sR.Travelers, Travelers{
				ID:           strconv.Itoa(paxCount),
				TravelerType: "INFANT",
			})

			paxCount++
		}

	}

}
