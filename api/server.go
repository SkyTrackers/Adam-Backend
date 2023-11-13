package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SkyTrackers/Adam-Backend/api/controller"
	"github.com/SkyTrackers/Adam-Backend/service/flight"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	Server *http.Server
}

func NewApiServer(port int32, flightService *flight.Service) *ApiServer {
	router := mux.NewRouter()

	controller.NewFlightHandler(flightService).RegisterHandlers(router)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	return &ApiServer{
		Server: server,
	}
}

func (api *ApiServer) Start() {
	err := api.Server.ListenAndServe()
	panic(err)
}
