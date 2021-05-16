package storage

import "errors"
import "github.com/m0stly1/playground1/model"
import "strconv"
/*
* - to do
*
* fix returns 
* fix errors
*
*/

var messages = map[string]*model.Message{
	"1": {Id: "1", Title: "first-one", Content: "random first post"},
	"2": {Id: "2", Title: "secound-one", Content: "random secound post"},
}

func GetMessages()(map[string]*model.Message, error){

	if len(messages) < 0 {
		return nil, errors.New("messages not found")
	}


	return messages, nil
}



func GetMessage(msg_id string) (*model.Message, error) {

	msg_exists := Exists(msg_id)

	if msg_exists {
		return messages[msg_id], nil
	}

	return nil, errors.New("message not found")
}

func DeleteMessage(msg_id string) (bool, error) {

	msg_exists := Exists(msg_id)

	if msg_exists {
		delete(messages, msg_id)
		return true, nil
	}

	return false, errors.New("message not found")
}

func CreateMessage(msg *model.Message) (bool,error) {

	err := Validate(msg)

	if err != nil {
		return false, err
	}

	id := LastId()
	msg.Id = id
	messages[id] = msg

	return true, nil
}

func UpdateMessage(msg *model.Message) (bool, error) {

	msg_exists := Exists(msg.Id)

	if msg_exists {
		messages[msg.Id] = msg
		return true, nil
	}

	return false, errors.New("Message does not exist")
}


func Validate(m *model.Message) error {

	if m.Content == "" {
		return errors.New("Message is required")
	}

	return nil
}

func Exists(msg_id string) bool {

	if messages[msg_id] != nil {
		return true
	}

	return false
}

func LastId() string {

	id := len(messages) + 1

	return strconv.Itoa(id)
}
