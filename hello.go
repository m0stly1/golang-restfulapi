package main

import "fmt"
import "github.com/m0stly1/playground/model"

func main() {

	fmt.Println(GetMessage("0"))
	msg_Delete("0");
	fmt.Println(GetMessage("0"))


	m := &model.Message{
		Title : "hej hej",
    	Content : "new body vettu",
	}

	create(m)

	fmt.Println(GetMessage("3"))
}


func create(m *model.Message){

	model.CreateMessage(m);
}


func GetMessage(msg_id string) (*model.Message, error){

	return model.GetMessage(msg_id)
}

func msg_Delete(id string){

	model.DeleteMessage(id)
}
