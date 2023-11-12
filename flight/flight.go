package flight

import "time"

type Flight struct {
	Passengers         int32     `json:"passengers"`
	AirplaneId         int32     `json:"airplane_id"`
	DepartureAirportId int32     `json:"departure_airport_id"`
	ArrivalAirportId   int32     `json:"arrival_airport_id"`
	DepartureTime      time.Time `json:"departure_time"`
	ArrivalTime        time.Time `json:"arrival_time"`
}

func NewFlight(
	passengers int32,
	airplaneId int32,
	departureAirportId int32,
	arrivalAirportId int32,
	departureTime time.Time,
	arrivalTime time.Time,
) *Flight {
	return &Flight{
		Passengers:         passengers,
		AirplaneId:         airplaneId,
		DepartureAirportId: departureAirportId,
		ArrivalAirportId:   arrivalAirportId,
		DepartureTime:      departureTime,
		ArrivalTime:        arrivalTime,
	}
}
