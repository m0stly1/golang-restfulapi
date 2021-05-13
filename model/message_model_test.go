package model



import (
	"testing"
)



func TestGetMessage(t *testing.T){

	msg, err := GetMessage("3")

	if msg != nil{
		t.Error("Message should not exist")
	}

	if err == nil {
		t.Error("should have error msg")
	}


	msg_success, err := GetMessage("0")

	if msg_success == nil {
		t.Error("Message should not exist")
	}

	if err != nil {
		t.Error("should not have error here")
	}

}



func TestDeleteMessage(t *testing.T){

	size_before := LastId()
	removed := DeleteMessage("0")
	size_after := LastId()

	if !removed {

		msg_success, err := GetMessage("0")

		if msg_success == nil {
		t.Error("Message should not exist")
		}

		if err != nil {
			t.Error("should not have error here")
		}

		if size_before == size_after{
			t.Error("message not deleted, size is the same")
		}

		t.Error("msg should be removed")
	}

}


func TestCreateMessage(t *testing.T){

	m := &Message{
		Title : "hej hej",
		Content : "",
	}

	iscreated, err := CreateMessage(m)

	if iscreated{
		t.Error("Msg should not be created. Body is empty")
	}

	if err == nil{
		t.Error("message is not valid")
	}


}