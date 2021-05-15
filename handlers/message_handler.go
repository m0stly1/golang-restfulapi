package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/m0stly1/playground1/model"
	"github.com/m0stly1/playground1/service"
	"github.com/m0stly1/playground1/utils"
	"net/http"
)

var (
	s service.MessageService = service.NewMessageService()
)

func GetMessage(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		utils.JsonResponse(w, http.StatusBadRequest, utils.ServiceError{Message: "id is required"})
		return
	}

	result, err := s.Get(id)

	if err != nil {
		utils.JsonResponse(w, http.StatusNotFound, utils.ServiceError{Message: "Message not found"})
		return
	}

	utils.JsonResponse(w, http.StatusOK, result)
}

func AddMessage(w http.ResponseWriter, r *http.Request) {

	var msg *model.Message

	err := json.NewDecoder(r.Body).Decode(&msg)

	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, utils.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	result, err2 := s.Create(msg)
	if err2 != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, utils.ServiceError{Message: "Error saving the message"})
		return
	}

	utils.JsonResponse(w, http.StatusCreated, result)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		utils.JsonResponse(w, http.StatusBadRequest, utils.ServiceError{Message: "Id is required"})
		return
	}

	result, err := s.Delete(id)

	if err != nil {
		utils.JsonResponse(w, http.StatusNotFound, utils.ServiceError{Message: "Error removing the message"})
		return
	}

<<<<<<< HEAD
	utils.JsonResponse(w, http.StatusOK, result)
=======
	utils.JsonResponse(w, http.StatusCreated, result)
>>>>>>> master
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {

	var msg *model.Message

	id := mux.Vars(r)["id"]

	if id == "" {
		utils.JsonResponse(w, http.StatusBadRequest, utils.ServiceError{Message: "Missing id"})
	}

	err := json.NewDecoder(r.Body).Decode(&msg)

	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, utils.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	msg.Id = id
	result, err2 := s.Update(msg)
	if err2 != nil {
		utils.JsonResponse(w, http.StatusBadRequest, utils.ServiceError{Message: "Error saving the message"})
		return
	}

	utils.JsonResponse(w, http.StatusCreated, result)
}
