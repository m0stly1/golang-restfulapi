package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/m0stly1/golang-restfulapi/models"
	"github.com/m0stly1/golang-restfulapi/service"
	"github.com/m0stly1/golang-restfulapi/utils"
	"net/http"
)

var (
	s service.MessageService = service.NewMessageService()
)




func GetMessages (w http.ResponseWriter, r *http.Request) {

	messages, err := s.GetAll()


	if err != nil{
		utils.JsonResponse(w, http.StatusNotFound, utils.ServiceError{Message: http.StatusText(404)})
		return
	} 


	utils.JsonResponse(w, http.StatusOK, messages)
}


func GetMessage(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		utils.JsonResponse(w, http.StatusBadRequest, utils.ServiceError{Message: http.StatusText(400)})
		return
	}

	result, err := s.Get(id)

	if err != nil {
		utils.JsonResponse(w, http.StatusNotFound, utils.ServiceError{Message: http.StatusText(404)})
		return
	}

	utils.JsonResponse(w, http.StatusOK, result)
}

func AddMessage(w http.ResponseWriter, r *http.Request) {

	var msg *models.Message

	err := json.NewDecoder(r.Body).Decode(&msg)

	if err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, utils.ServiceError{Message: http.StatusText(400)})
		return
	}

	result, err2 := s.Create(msg)
	if err2 != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, utils.ServiceError{Message: http.StatusText(500)})
		return
	}

	utils.JsonResponse(w, http.StatusCreated, result)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	if id == "" {
		utils.JsonResponse(w, http.StatusBadRequest, utils.ServiceError{Message: http.StatusText(400)})
		return
	}

	result, err := s.Delete(id)

	if err != nil {
		utils.JsonResponse(w, http.StatusNotFound, utils.ServiceError{Message: http.StatusText(404)})
		return
	}

	utils.JsonResponse(w, http.StatusOK, result)

}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {

	var msg *models.Message

	id := mux.Vars(r)["id"]

	if id == "" {
		utils.JsonResponse(w, http.StatusBadRequest, utils.ServiceError{Message: http.StatusText(400)})
	}

	err := json.NewDecoder(r.Body).Decode(&msg)

	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, utils.ServiceError{Message: http.StatusText(500)})
		return
	}

	msg.Id = id
	result, err2 := s.Update(msg)
	if err2 != nil {
		utils.JsonResponse(w, http.StatusNotFound, utils.ServiceError{Message: http.StatusText(404)})
		return
	}

	utils.JsonResponse(w, http.StatusCreated, result)
}
