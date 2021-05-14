package handlers

import (
	"bytes"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMessage(t *testing.T) {

	r := mux.NewRouter()
	r.HandleFunc("/message/{id:[0-9]+}", GetMessage).Methods("GET")

	ts := httptest.NewServer(r)

	defer ts.Close()

	// our hard coded messages
	ids := []string{"1", "2"}
	for _, id := range ids {
		url := ts.URL + "/message/" + id

		resp, err := http.Get(url)
		if err != nil {
			t.Fatal(err)
		}

		if status := resp.StatusCode; status != http.StatusOK {
			t.Fatalf("wrong status code: got %d want %d", status, http.StatusOK)
		}

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

func TestGetMessageStatusNotFound(t *testing.T) {

	r := mux.NewRouter()
	r.HandleFunc("/message/{id:[0-9]+}", GetMessage).Methods("GET")

	ts := httptest.NewServer(r)

	defer ts.Close()

	url := ts.URL + "/message/32"

	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusNotFound {
		t.Fatalf("wrong status code: got %d want %d", status, http.StatusNotFound)
	}
}

func TestGetMessageMissingId(t *testing.T) {

	r := mux.NewRouter()
	r.HandleFunc("/message/{id:[0-9]+}", GetMessage).Methods("GET")

	ts := httptest.NewServer(r)

	defer ts.Close()

	url := ts.URL + "/message/"

	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusNotFound {
		t.Fatalf("wrong status code: got %d want %d", status, http.StatusNotFound)
	}
}

func TestAddMessage(t *testing.T) {

	r := mux.NewRouter()
	r.HandleFunc("/message/", AddMessage).Methods("POST")

	ts := httptest.NewServer(r)

	defer ts.Close()

	url := ts.URL + "/message/"

	message := []byte(`{"title":"test", "content":"body"}`)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(message))

	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusCreated {
		t.Fatalf("wrong status code: got %d want %d", status, http.StatusOK)
	}
}

func TestAddMessageStatusBadRequest(t *testing.T) {

	r := mux.NewRouter()
	r.HandleFunc("/message/", AddMessage).Methods("POST")

	ts := httptest.NewServer(r)

	defer ts.Close()

	url := ts.URL + "/message/"

	// syntax error in json request
	message := []byte(`{"title":"test", "content":,,,"body"}`)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(message))

	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Fatalf("wrong status code: got %d want %d", status, http.StatusOK)
	}
}

func TestAddMessageMissingContent(t *testing.T) {

	r := mux.NewRouter()
	r.HandleFunc("/message/", AddMessage).Methods("POST")

	ts := httptest.NewServer(r)

	defer ts.Close()

	url := ts.URL + "/message/"

	message := []byte(`{"title":"test"}`)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(message))

	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusInternalServerError {
		t.Fatalf("wrong status code: got %d want %d", status, http.StatusInternalServerError)
	}
}

func TestUpdateMessage(t *testing.T) {

	r := mux.NewRouter()
	r.HandleFunc("/message/{id:[0-9]+}", UpdateMessage).Methods("PUT")

	ts := httptest.NewServer(r)

	url := ts.URL + "/message/1"

	message := []byte(`{"title":"test", "content":"body"}`)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(message))

	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusCreated {
		t.Fatalf("wrong status code: got %d want %d", status, http.StatusCreated)
	}
}

func TestUpdateMessageInternal(t *testing.T) {

	r := mux.NewRouter()
	r.HandleFunc("/message/{id:[0-9]+}", UpdateMessage).Methods("PUT")

	ts := httptest.NewServer(r)

	url := ts.URL + "/message/1"

	// syntax error
	message := []byte(`{"title":"test"`)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(message))

	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusInternalServerError {
		t.Fatalf("wrong status code: got %d want %d", status, http.StatusInternalServerError)
	}
}
