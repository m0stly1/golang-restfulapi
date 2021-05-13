package model

import "errors"
import "fmt"

var messages = map[string]*Message{
	"1" : {Id: "1", Title : "first-one", Content: "random first post"},
	"2" : {Id: "2", Title : "secound-one", Content: "random secound post"},
}


func GetMessage (msg_id string) (*Message, error){

	msg_exists := Exists(msg_id)

	if (msg_exists){
		return messages[msg_id], nil
	}

	return nil, errors.New("message do not exist")
}


func DeleteMessage(msg_id string) (bool, error){

	msg_exists := Exists(msg_id)
	fmt.Println("heej")
	if (msg_exists){
		delete(messages, msg_id)
		return true, nil
	}

	return false, errors.New("something very serious")
}


func CreateMessage(msg *Message) (bool, error){

	isvalid, err := Validate(msg)

	if !isvalid{
		return false, err
	}

	id := LastId()
	msg.Id = id
	messages[id] = msg

	return true, nil
}
