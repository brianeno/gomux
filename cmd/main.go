package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/brianeno/gomux/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// First create a new router
	r := mux.NewRouter()
	// Here we Attach a path with handler
	r.HandleFunc("/chargesessions", handlers.GetAllChargeSessions).Methods(http.MethodGet)
	r.HandleFunc("/chargesessions/{id}", handlers.GetChargeSessions).Methods(http.MethodGet)
	r.HandleFunc("/chargesessions", handlers.AddChargeSessions).Methods(http.MethodPost)
	r.HandleFunc("/chargesessions/{id}", handlers.UpdateChargeSessions).Methods(http.MethodPut)
	r.HandleFunc("/chargesessions/{id}", handlers.DeleteChargeSessions).Methods(http.MethodDelete)

	addr := "127.0.0.1:8000"
	srv := &http.Server{
		Handler: r,
		Addr:    addr,
		// enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Application started on", addr)
	log.Fatal(srv.ListenAndServe())
}
