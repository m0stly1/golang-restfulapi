package main


import (
	"github.com/m0stly1/playground1/model"
	"github.com/m0stly1/playground1/service"
	"github.com/m0stly1/playground1/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
	"time"
	"encoding/json"

	)


var (
	storage service.MessageService = service.NewMessageService()
)



func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/message", handlers.AddMessage).Methods("POST")
	r.HandleFunc("/api/message/{id:[0-9]+}", handlers.GetMessage).Methods("GET")
	r.HandleFunc("/api/message", updateMessage).Methods("PUT")
	r.HandleFunc("/api/message", deleteMessage).Methods("DELETE")


	srv := &http.Server{
		Handler:	r,
		Addr:		"127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}



func addMessage(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	var msg *model.Message

	err := json.NewDecoder(r.Body).Decode(&msg)

	storage.Create(msg)

	if err != nil{

	}

}

func updateMessage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "put", r.URL.Path[1:])
}

func deleteMessage(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	id := mux.Vars(r)["id"]

	if id == "" {
		w.WriteHeader(500)
		fmt.Fprintf(w, "id is required")
	}

	storage.Delete(id)

}
