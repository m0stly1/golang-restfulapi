package service

import "github.com/m0stly1/playground1/model"



type Accessor interface{
	GetMessage(n string)
}

type MessageService struct {
	a Accessor
}

func NewMessageService (a Accessor) MessageService{
	return MessageService {
		a: a,
	}
}


func (p MessageService) Get (n int) (*Message, error){

	p := ps.a.GetMessage(n)

	if p.First == ""{
		return *Message {}, fmt.Errorf("no Person found")
	}

	return p, nil
}