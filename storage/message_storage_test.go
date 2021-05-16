package storage

import (
	"testing"
	"github.com/m0stly1/golang-restfulapi/models"
)


func TestCreateMessages(t *testing.T) {

	tt := []struct {
		name       string
		testCase   int
		input      *models.Message
		expected   bool
	}{
		{
			name:       "No Message, Fail",
			testCase:	1,
			input:      &models.Message{},
			expected:   false,
		},
		{
			name:       "No Content, Fail",
			testCase:	2,
			input:      &models.Message{Title: "I do not exist", Content : ""},
			expected:   false,
		},
		{
			name:       "Valid message 1",
			testCase:	3,
			input:      &models.Message{Title: "This is a new message", Content : "And this is it´s content"},
			expected:   true,
		},
		{
			name:       "Valid message 2",
			testCase:	4,
			input:      &models.Message{Title: "I do not exist, create me", Content : "And more content"},
			expected:   true,
		},
	}


	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			actual, _ := CreateMessage(tc.input)

			if actual != tc.expected{
				t.Error("message is not valid")
			}

		})
	}

}



func TestUpdateMessageError(t *testing.T){
	
	tt := []struct {
		name       string
		testCase   int 
		input      *models.Message
		expected   bool
	}{
		{
			name:       "no message",
			testCase:	1,
			input:      &models.Message{},
			expected:   false,
		},
		{
			name:       "none existing id",
			testCase:	2,
			input:      &models.Message{Id:"43", Title: "I do not exist"},
			expected:   false,
		},
	}


	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			actual, _ := UpdateMessage(tc.input)

			if actual != tc.expected{
				t.Error("message is not valid")
			}

		})
	}

}



func TestUpdateMessageSuccess(t *testing.T){
	
	tt := []struct {
		name       string
		testCase   int 
		input      *models.Message
		expected   bool
	}{
		{
			name:       "Updating message 1",
			testCase:	1,
			input:      &models.Message{Id:"2", Title: "I do exist, update me", Content: "Random Content"},
			expected:   true,
		},
		{
			name:       "Updating message 2",
			testCase:	2,
			input:      &models.Message{Id:"2", Title: "I do also exist, update me", Content: "Random Content"},
			expected:   true,
		},
	}


	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			actual, _ := UpdateMessage(tc.input)

			if actual != tc.expected{
				
				t.Error("Test Case failed", tc.testCase)
			}

		})
	}

}



func TestGetMessage(t *testing.T) {

	msg, err := GetMessage("32")

	if msg != nil {
		t.Error("Message should not exist")
	}

	if err == nil {
		t.Error("should have error msg")
	}

	msg_success, err := GetMessage("1")

	if msg_success == nil {
		t.Error("Message should not exist")
	}

	if err != nil {
		t.Error("should not have error here")
	}

}

func TestDeleteMessage(t *testing.T) {

	size_before := LastId()
	removed, _ := DeleteMessage("1")
	size_after := LastId()

	if !removed {

		msg_success, err := GetMessage("1")

		if msg_success == nil {
			t.Error("Message should not exist")
		}

		if err != nil {
			t.Error("should not have error here")
		}

		if size_before == size_after {
			t.Error("message not deleted, size is the same")
		}

		t.Error("msg should be removed")
	}

}