package service

import "github.com/m0stly1/golang-restfulapi/storage"
import "github.com/m0stly1/golang-restfulapi/model"

type MessageService interface {
	Get(id string) (*model.Message, error)
	GetAll()(map[string]*model.Message, error)
	Create(*model.Message) (bool, error)
	Delete(id string) (bool, error)
	Update(*model.Message) (bool, error)
}

type service struct{}
 
func NewMessageService() MessageService {
	return &service{}
}

func (*service) GetAll() (map[string]*model.Message, error) {
	return storage.GetMessages()
}

func (*service) Get(id string) (*model.Message, error) {
	return storage.GetMessage(id)
}

func (*service) Create(m *model.Message) (bool, error) {
	return storage.CreateMessage(m)
}

func (*service) Delete(id string) (bool, error) {
	return storage.DeleteMessage(id)
}

func (*service) Update(m *model.Message) (bool, error) {
	return storage.UpdateMessage(m)
}
