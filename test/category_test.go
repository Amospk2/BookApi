package main

import (
	"api/domain/category"
	"api/infra/controllers"
	"api/test/mock"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func TestCategoryListWhenIsOk(t *testing.T) {

	ctl := controllers.NewCategoryController(
		mock.NewCategoryRepositoryMock([]category.Category{
			{
				Id:   "123",
				Name: "Doces",
			},
		}),
	)

	req, err := http.NewRequest("GET", "/categorys", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctl.GetCategorys())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"categorys":[{"id":"123","name":"Doces"}]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCategoryListWhenHasQuery(t *testing.T) {

	ctl := controllers.NewCategoryController(
		mock.NewCategoryRepositoryMock([]category.Category{
			{
				Id:   "123",
				Name: "Doces",
			},
		}),
	)

	req, err := http.NewRequest("GET", "/categorys?test=11", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctl.GetCategorys())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"categorys":[{"id":"123","name":"Doces"}]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCategoryGetById(t *testing.T) {

	ctl := controllers.NewCategoryController(
		mock.NewCategoryRepositoryMock([]category.Category{
			{
				Id:   "123",
				Name: "Doces",
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/categorys/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/categorys/{id}", ctl.GetCategoryById())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"123","name":"Doces"}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCategoryGetByIdNotFound(t *testing.T) {

	ctl := controllers.NewCategoryController(
		mock.NewCategoryRepositoryMock([]category.Category{
			{
				Id:   "123",
				Name: "Doces",
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/categorys/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/categorys/{id}", ctl.GetCategoryById())
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

func TestCategoryCreate(t *testing.T) {

	ctl := controllers.NewCategoryController(
		mock.NewCategoryRepositoryMock([]category.Category{}),
	)

	rr := httptest.NewRecorder()

	content := category.Category{
		Id:   "123",
		Name: "Doces",
	}

	jsonData, err := json.Marshal(content)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/categorys", bytes.NewReader([]byte(jsonData)))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", jwt.MapClaims{
		"user": "123",
		"exp":  time.Now().Add(time.Duration(time.Hour * 1)).Unix(),
	})

	req = req.WithContext(ctx)

	router := mux.NewRouter()
	router.HandleFunc("/categorys", ctl.CreateCategory())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	category := category.Category{}
	json.Unmarshal(rr.Body.Bytes(), &category)

	fmt.Println(category)
	if !(strings.Compare(category.Name, "Doces") == 0) {
		t.Errorf("handler returned unexpected body")
	}
}

func TestCategoryUpdate(t *testing.T) {

	ctl := controllers.NewCategoryController(
		mock.NewCategoryRepositoryMock([]category.Category{
			{
				Id:   "123",
				Name: "Doces",
			},
		}),
	)

	rr := httptest.NewRecorder()

	content := category.Category{
		Id:   "123",
		Name: "Salgados",
	}

	jsonData, err := json.Marshal(content)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/categorys/123", bytes.NewReader([]byte(jsonData)))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", jwt.MapClaims{
		"user": "123",
		"exp":  time.Now().Add(time.Duration(time.Hour * 1)).Unix(),
	})

	req = req.WithContext(ctx)

	router := mux.NewRouter()
	router.HandleFunc("/categorys/{id}", ctl.UpdateCategory())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	productEdited := category.Category{}
	json.Unmarshal(rr.Body.Bytes(), &productEdited)

	fmt.Println(productEdited)
	if !(strings.Compare(productEdited.Name, "Salgados") == 0) {
		t.Errorf("handler returned unexpected body")
	}
}

func TestCategoryDelete(t *testing.T) {

	ctl := controllers.NewCategoryController(
		mock.NewCategoryRepositoryMock([]category.Category{
			{
				Id:   "123",
				Name: "Doces",
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("DELETE", "/categorys/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/categorys/{id}", ctl.Delete())
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
