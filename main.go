package main

import (
	"log"
	"net/http"
	"olympic-medals-table/actions"
	"olympic-medals-table/data_base"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	data_base.Config()
	//db.Find(&AllCountries)
	//fmt.Println(&AllCountries)

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api/countries", data_base.GetCountries).Methods("GET")
	r.HandleFunc("/api/countries/{id}", data_base.GetCountry).Methods("GET")
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
