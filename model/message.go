package model

import "errors"

type Message struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Content string `json:content`
}


func Validate(m *Message) (bool, error){

	if m.Content == "" {
		return false, errors.New("Message is required")
	}

	return true, nil
}


func Exists(msg_id string) bool{

	if messages[msg_id] != nil {
		return true
	}

	return false
}


func LastId() int{

	return len(messages) 
}