package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)


func TestGetMessage(t *testing.T) { 

	r := mux.NewRouter()
	r.HandleFunc("/message/{id:[0-9]+}", GetMessage).Methods("GET")

	ts := httptest.NewServer(r)

	defer ts.Close()



	tt := []struct {
		name       string
		testCase   int 
		input      string
		statusCode int
	}{
		{
			name:       "Message Exists",
			testCase:	1,
			input:      "1",
			statusCode:   200,
		},
		{
			name:       "Message Exists",
			testCase:	2,
			input:      "2",
			statusCode:   200,
		},
		{
			name:       "Message does not exists",
			testCase:	3,
			input:      "500",
			statusCode:   404,
		},
		{
			name:       "No id",
			testCase:	4,
			input:      "",
			statusCode:   404,
		},
	}


	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			url := ts.URL + "/message/" + tc.input

			resp, err := http.Get(url)
			if err != nil {
				t.Fatal(err)
			}

			status := resp.StatusCode;

			if status != tc.statusCode{
				t.Fatalf("wrong status code: got %d want %d on test case %d", status, tc.statusCode, tc.testCase)
			}
		})
	}
}




func TestGetMessageStatusMethodNotAllowed(t *testing.T) {

	r := mux.NewRouter()
	r.HandleFunc("/message/{id:[0-9]+}", GetMessage).Methods("GET")

	ts := httptest.NewServer(r)

	defer ts.Close()

	url := ts.URL + "/message/1"

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusMethodNotAllowed {
		t.Fatalf("wrong status code: got %d want %d", status, http.StatusMethodNotAllowed)
	}
}

func TestAddMessage(t *testing.T) { 

	r := mux.NewRouter()
	r.HandleFunc("/message/", AddMessage).Methods("POST")

	ts := httptest.NewServer(r)

	defer ts.Close()

	url := ts.URL + "/message/"


	tt := []struct {
		name       string
		testCase   int 
		input      string
		statusCode int
	}{
		{
			name:       "Error unmarshalling data",
			testCase:	1,
			input:      `{"title":"test", "content":,,,"body"}`,
			statusCode:   500,
		},
		{
			name:       "No Content",
			testCase:	2,
			input:      `{"title":"test"}`,
			statusCode:   500,
		},
		{
			name:       "Valid message 1",
			testCase:	3,
			input:      `{"title":"A new Hope (title) ", "content":"A New Content (Hope)"}`,
			statusCode:   201,
		},
		{
			name:       "Valid message 2",
			testCase:	4,
			input:      `{"title":"The title Strikes back ", "content":"The Content Strikes Back"}`,
			statusCode:   201,
		},
	}


	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			resp, err := http.Post(url, "application/json", strings.NewReader(tc.input))
			if err != nil {
				t.Fatal(err)
			}

			status := resp.StatusCode;

			if status != tc.statusCode{
				t.Fatalf("wrong status code: got %d want %d on test case %d", status, tc.statusCode, tc.testCase)
			}
		})
	}
}


func TestUpdateMessage(t *testing.T) { 

	r := mux.NewRouter()
	r.HandleFunc("/message/{id:[0-9]+}", UpdateMessage).Methods("PUT")

	ts := httptest.NewServer(r)

	tt := []struct {
		id		string
		name       string
		testCase   int 
		input      string
		statusCode int
	}{
		{
			id :		"1",
			name:       "Error unmarshalling data",
			testCase:	1,
			input:      `{"title":"test", "content":,,,"body"}`,
			statusCode:   500,
		},
		{
			id :		"1",
			name:       "No Content",
			testCase:	2,
			input:      `{"title":"Mandalorian", "content":"no spoilers"}`,
			statusCode:   201,
		},
		{
			id :		"10",
			name:       "Message does not exists",
			testCase:	3,
			input:      `{"title":"The title Strikes back", "content":"The Content Strikes Back"}`,
			statusCode:   400,
		},
		{
			id :		"2",
			name:       "Valid message 2",
			testCase:	4,
			input:      `{"title":"The title Strikes back ", "content":"The Content Strikes Back"}`,
			statusCode:   201,
		},
	}


	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			url := ts.URL + "/message/" + tc.id

			req, err := http.NewRequest("PUT", url, strings.NewReader(tc.input))

			if err != nil {
				t.Fatal(err)
			}

			client := &http.Client{}
			resp, err := client.Do(req)

			if err != nil {
				t.Fatal(err)
			}

			status := resp.StatusCode;

			if status != tc.statusCode{
				t.Fatalf("wrong status code: got %d want %d on test case %d", status, tc.statusCode, tc.testCase)
			}
		})
	}
}




func TestDeleteMessage(t *testing.T) {

	r := mux.NewRouter()
	r.HandleFunc("/message/{id:[0-9]+}", DeleteMessage).Methods("DELETE")

	ts := httptest.NewServer(r)

	url := ts.URL + "/message/1"

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusOK {
		t.Fatalf("wrong status code: got %d want %d", status, http.StatusOK)
	}
}