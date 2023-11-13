package flight

import (
	"strconv"
)

type Service struct {
	Flights []Flight
}

func NewService() *Service {
	return &Service{[]Flight{}}
}

func (s *Service) GetFlights() []Flight {
	return s.Flights
}

func (s *Service) GetFlight(flightId string) *Flight {
	flights := s.GetFlights()
	for _, flight := range flights {
		if strconv.Itoa(int(flight.Id)) == flightId {
			return &flight
		}
	}
	return nil
}

func (s *Service) AddFlight(flight Flight) {
	s.Flights = append(s.Flights, flight)
}
