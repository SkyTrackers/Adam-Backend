package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/SkyTrackers/Adam-Backend/flight"
	"github.com/gorilla/mux"
)

type Flights []*flight.Flight

var flights Flights

func main() {
	loadDefaultFlights()
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", index)
	router.HandleFunc("/flights", flightsIndex).Methods("GET")
	router.HandleFunc("/flights", flightsCreate).Methods("POST")
	router.HandleFunc("/flights/{flightId}", flightIdInfo)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func loadDefaultFlights() Flights {
	defaultFlight1 := flight.NewFlight(100, 1, 1, 2, time.Now(), time.Now())
	flights = append(flights, defaultFlight1)
	return flights
}

func getAllFlights() Flights {
	return flights
}

func flightsCreate(w http.ResponseWriter, r *http.Request) {
	// create new flight
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

	flights = append(flights, &flight)
	fmt.Fprintf(w, "Flight successfuly created")
	w.WriteHeader(http.StatusOK)
}

func flightsIndex(w http.ResponseWriter, r *http.Request) {
	// show all flights
	json.NewEncoder(w).Encode(getAllFlights())
}

func flightIdInfo(w http.ResponseWriter, r *http.Request) {
	// show specific flight
	vars := mux.Vars(r)
	flightId := vars["flightId"]
	fmt.Fprint(w, "Information from flight ", flightId)
}
