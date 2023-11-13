package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/SkyTrackers/Adam-Backend/service/flight"
	"github.com/gorilla/mux"
)

type FlightHandler struct {
	service *flight.Service
}

func NewFlightHandler(flightService *flight.Service) *FlightHandler {
	return &FlightHandler{flightService}
}

func (h *FlightHandler) RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/flights", h.createFlight).Methods("POST")
	router.HandleFunc("/flights", h.getFlights).Methods("GET")
	router.HandleFunc("/flights/{flightId}", h.getFlight).Methods("GET")
}

func (h *FlightHandler) createFlight(w http.ResponseWriter, r *http.Request) {
	var flight flight.Flight
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	err = json.Unmarshal(body, &flight)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	h.service.AddFlight(flight)
	fmt.Fprintf(w, "Flight created successfuly")
	w.WriteHeader(http.StatusOK)
}

func (h *FlightHandler) getFlights(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.service.GetFlights())
}

func (h *FlightHandler) getFlight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	flightId := vars["flightId"]
	flight := h.service.GetFlight(flightId)
	json.NewEncoder(w).Encode(flight)
}
