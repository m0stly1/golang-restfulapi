package storage

import "errors"
import "github.com/m0stly1/golang-restfulapi/models"
import "strconv"

var messages = map[string]*models.Message{
	"1": {Id: "1", Title: "Rogue One: A Star Wars Story", Content: "More Star Wars"},
	"2": {Id: "2", Title: "Solo: A Star Wars Story", Content: "Almost forgot this one"},
}

func GetMessages() (map[string]*models.Message, error) {

	if len(messages) < 0 {
		return nil, errors.New("messages not found")
	}

	return messages, nil
}

func GetMessage(msg_id string) (*models.Message, error) {

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

func CreateMessage(msg *models.Message) (bool, error) {

	err := Validate(msg)

	if err != nil {
		return false, err
	}

	id := LastId()
	msg.Id = id
	messages[id] = msg

	return true, nil
}

func UpdateMessage(msg *models.Message) (bool, error) {

	msg_exists := Exists(msg.Id)

	if msg_exists {
		messages[msg.Id] = msg
		return true, nil
	}

	return false, errors.New("Message does not exist")
}

func Validate(m *models.Message) error {

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
