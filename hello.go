package main


import (
	"github.com/m0stly1/playground1/model"
	"github.com/gorilla/mux"
    "log"
    "net/http"
    "fmt"
    "time"
    )

func main() {
	
	r := mux.NewRouter()

	r.HandleFunc("/api/message", addMessage).Methods("POST")
	r.HandleFunc("/api/message/{id:[0-9]+}", getMessage).Methods("GET")
	r.HandleFunc("/api/message", updateMessage).Methods("PUT")
	r.HandleFunc("/api/message", deleteMessage).Methods("DELETE")


	srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:8000",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())

}


func getMessage(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	id := mux.Vars(r)["id"]

	if !id {

	}



    fmt.Fprintf(w, id)
}

func addMessage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "post", r.URL.Path[1:])
}

func updateMessage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "put", r.URL.Path[1:])
}

func deleteMessage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "delete", r.URL.Path[1:])
}

func create(m *model.Message){

	model.CreateMessage(m);
}


func GetMessage(msg_id string) (*model.Message, error){

	return model.GetMessage(msg_id)
}

func msg_Delete(id string){

	model.DeleteMessage(id)
}
