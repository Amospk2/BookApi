package main

import (
	"api/domain/author"
	"api/infra/controllers"
	"api/test/mock"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func TestAuthorListWhenIsOk(t *testing.T) {

	ctl := controllers.NewAuthorController(
		mock.NewAuthorRepositoryMock([]author.Author{
			{
				Id:   "123",
				Name: "Kentaro Miura",
			},
		}),
	)

	req, err := http.NewRequest("GET", "/author", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctl.GetAuthors())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"authors":[{"id":"123","name":"Kentaro Miura"}]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAuthorListWhenHasQuery(t *testing.T) {

	ctl := controllers.NewAuthorController(
		mock.NewAuthorRepositoryMock([]author.Author{
			{
				Id:   "123",
				Name: "Kentaro Miura",
			},
		}),
	)

	req, err := http.NewRequest("GET", "/author?test=11", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctl.GetAuthors())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"authors":[{"id":"123","name":"Kentaro Miura"}]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAuthorGetById(t *testing.T) {

	ctl := controllers.NewAuthorController(
		mock.NewAuthorRepositoryMock([]author.Author{
			{
				Id:   "123",
				Name: "Kentaro Miura",
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/author/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/author/{id}", ctl.GetAuthorById())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"123","name":"Kentaro Miura"}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}


func TestAuthorGetByIdNotFound(t *testing.T) {

	ctl := controllers.NewAuthorController(
		mock.NewAuthorRepositoryMock([]author.Author{
			{
				Id:   "123",
				Name: "Kentaro Miura",
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/author/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/author/{id}", ctl.GetAuthorById())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := ``

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAuthorCreate(t *testing.T) {

	ctl := controllers.NewAuthorController(
		mock.NewAuthorRepositoryMock([]author.Author{}),
	)

	rr := httptest.NewRecorder()

	content := author.Author{
		Id:   "123",
		Name: "Kentaro Miura",
	}

	jsonData, err := json.Marshal(content)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/author", bytes.NewReader([]byte(jsonData)))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", jwt.MapClaims{
		"user": "123",
		"exp":  time.Now().Add(time.Duration(time.Hour * 1)).Unix(),
	})

	req = req.WithContext(ctx)

	router := mux.NewRouter()
	router.HandleFunc("/author", ctl.CreateAuthor())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	author := author.Author{}
	json.Unmarshal(rr.Body.Bytes(), &author)

	if !(strings.Compare(author.Name, "Kentaro Miura") == 0) {
		t.Errorf("handler returned unexpected body")
	}
}

func TestAuthorUpdate(t *testing.T) {

	ctl := controllers.NewAuthorController(
		mock.NewAuthorRepositoryMock([]author.Author{
			{
				Id:   "123",
				Name: "Kentaro Miura",
			},
		}),
	)

	rr := httptest.NewRecorder()

	content := author.Author{
		Id:   "123",
		Name: "Araki",
	}

	jsonData, err := json.Marshal(content)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/author/123", bytes.NewReader([]byte(jsonData)))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", jwt.MapClaims{
		"user": "123",
		"exp":  time.Now().Add(time.Duration(time.Hour * 1)).Unix(),
	})

	req = req.WithContext(ctx)

	router := mux.NewRouter()
	router.HandleFunc("/author/{id}", ctl.UpdateAuthor())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	author := author.Author{}
	json.Unmarshal(rr.Body.Bytes(), &author)

	if !(strings.Compare(author.Name, "Araki") == 0) {
		t.Errorf("handler returned unexpected body")
	}
}

func TestAuthorDelete(t *testing.T) {

	ctl := controllers.NewAuthorController(
		mock.NewAuthorRepositoryMock([]author.Author{
			{
				Id:   "123",
				Name: "Kentaro Miura",
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("DELETE", "/author/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/author/{id}", ctl.Delete())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := ``

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
