package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/m0stly1/playground1/utils"
	"github.com/m0stly1/playground1/model"
	"github.com/m0stly1/playground1/service"
	"net/http"
)

var (
	s service.MessageService = service.NewMessageService()
)

func GetMessage(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	w.Header().Add("content-type", "application/json")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ServiceError{Message: "id is required"})
		return
	}

	result, err := s.Get(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.ServiceError{Message: "Message not found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func AddMessage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var msg *model.Message

	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	result, err2 := s.Create(msg)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.ServiceError{Message: "Error saving the message"})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.ServiceError{Message: "Id is required"})
		return
	}

	result, err := s.Delete(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.ServiceError{Message: "Error removing the message"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var msg *model.Message

	id := mux.Vars(r)["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ServiceError{Message: "id is required"})
		return
	}

	err := json.NewDecoder(r.Body).Decode(&msg)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	msg.Id = id
	result, err2 := s.Update(msg)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ServiceError{Message: "Error saving the message"})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}

