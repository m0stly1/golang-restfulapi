package service

import "github.com/m0stly1/golang-restfulapi/storage"
import "github.com/m0stly1/golang-restfulapi/models"

type MessageService interface {
	Get(id string) (*models.Message, error)
	GetAll()(map[string]*models.Message, error)
	Create(*models.Message) (bool, error)
	Delete(id string) (bool, error)
	Update(*models.Message) (bool, error)
}

type service struct{}
 
func NewMessageService() MessageService {
	return &service{}
}

func (*service) GetAll() (map[string]*models.Message, error) {
	return storage.GetMessages()
}

func (*service) Get(id string) (*models.Message, error) {
	return storage.GetMessage(id)
}

func (*service) Create(m *models.Message) (bool, error) {
	return storage.CreateMessage(m)
}

func (*service) Delete(id string) (bool, error) {
	return storage.DeleteMessage(id)
}

func (*service) Update(m *models.Message) (bool, error) {
	return storage.UpdateMessage(m)
}
