package flight

import "time"

type Flight struct {
	Id                 int64     `json:"id"`
	Passengers         int64     `json:"passengers"`
	AirplaneId         int64     `json:"airplane_id"`
	DepartureAirportId int64     `json:"departure_airport_id"`
	ArrivalAirportId   int64     `json:"arrival_airport_id"`
	DepartureTime      time.Time `json:"departure_time"`
	ArrivalTime        time.Time `json:"arrival_time"`
}

func NewFlight(
	id int64,
	passengers int64,
	airplaneId int64,
	departureAirportId int64,
	arrivalAirportId int64,
	departureTime time.Time,
	arrivalTime time.Time,
) *Flight {
	return &Flight{
		Id:                 id,
		Passengers:         passengers,
		AirplaneId:         airplaneId,
		DepartureAirportId: departureAirportId,
		ArrivalAirportId:   arrivalAirportId,
		DepartureTime:      departureTime,
		ArrivalTime:        arrivalTime,
	}
}
