package handlers

import(
	"github.com/m0stly1/playground1/errors"
	"github.com/m0stly1/playground1/service"
	"github.com/m0stly1/playground1/model"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)


var (
	storage service.MessageService = service.NewMessageService()
)



func GetMessage(w http.ResponseWriter, r *http.Request) {


	id := mux.Vars(r)["id"]

	if id == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	msg, err := storage.Get(id)


	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error finding post"})
		return
	}


	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)

}


func AddMessage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var msg *model.Message

	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	result, err2 := storage.Create(msg)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)


}