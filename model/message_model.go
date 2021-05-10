package model

import "errors"

var messages = map[string]*Message{
	"0" : {Id: "1", Title : "first-one", Content: "random first post"},
	"1" : {Id: "2", Title : "secound-one", Content: "random secound post"},
}


func GetMessage (msg_id string) (*Message, error){

	msg_exists := Exists(msg_id)

	if (msg_exists){
		return messages[msg_id], nil
	}

	return nil, errors.New("message do not exist")
}


func DeleteMessage(msg_id string) bool{

	msg_exists := Exists(msg_id)

	if (msg_exists){
		delete(messages, msg_id)
		return true
	}

	return false
}


func CreateMessage(msg *Message) (bool, error){

	isvalid, err := Validate(msg)

	if !isvalid{
		return false, err
	}

	messages["3"] = msg

	return true, nil
}
