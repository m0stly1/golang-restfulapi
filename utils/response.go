package utils

import (
	"encoding/json"
	"net/http"
	"log"
)

func JsonResponse(w http.ResponseWriter, status int, i interface{}){
	
	response, err := json.Marshal(i)

	if err != nil{
		log.Fatal(err)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(response)

	if err != nil{
		log.Fatal(err)
		return
	}

}