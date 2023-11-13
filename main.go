package main

import (
	"github.com/SkyTrackers/Adam-Backend/api"
	"github.com/SkyTrackers/Adam-Backend/service/flight"
)

func main() {
	flightService := flight.NewService()

	apiServer := api.NewApiServer(8080, flightService)
	apiServer.Start()
}
