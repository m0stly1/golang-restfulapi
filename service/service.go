package service

import "github.com/m0stly1/playground1/model"



type MessageService interface {
	Get(id string) (*model.Message, error)
	Create(*model.Message) (bool, error)
	Delete(id string) (bool, error)
}


type service struct {}


func NewMessageService() MessageService {
	return &service{}
}


func (*service) Get (id string) (*model.Message, error){
	return model.GetMessage(id)
}


func (*service) Create (m *model.Message) (bool, error){
	return model.CreateMessage(m)
}

func (*service) Delete (id string) (bool, error){
	return model.DeleteMessage(id)
}

