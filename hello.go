package main


import (
	"github.com/m0stly1/playground1/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/message", handlers.AddMessage).Methods("POST")
	r.HandleFunc("/message/{id:[0-9]+}", handlers.GetMessage).Methods("GET")
	r.HandleFunc("/message/{id:[0-9]+}", handlers.UpdateMessage).Methods("PUT")
	r.HandleFunc("/message/{id:[0-9]+}", handlers.DeleteMessage).Methods("DELETE")
	r.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")

	srv := &http.Server{
		Handler:	r,
		Addr:		"127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}