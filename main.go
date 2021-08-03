package main

import (
	"log"
	"net/http"

	"olympic-medals/actions"
	"olympic-medals/api"
	"olympic-medals/data_base"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	data_base.Config()
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api/countries", api.GetCountries).Methods("GET")
	r.HandleFunc("/api/countries/{id}", api.GetCountry).Methods("GET")
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	r.HandleFunc("/", actions.FileServer)

	server := &http.Server{
		Addr:           ":3000",
		Handler:        r,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening...")
	server.ListenAndServe()

}
