package main

import (
	"api/domain/book"
	"api/infra/controllers"
	"api/test/mock"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/fs"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func TestBookListWhenIsOk(t *testing.T) {

	ctl := controllers.NewBookController(
		mock.NewBookRepositoryMock([]book.Book{
			{
				Id:           "123",
				Title:        "title",
				Category:     "category",
				Subtitle:     "subtitle",
				Description:  "description",
				Release_date: "2024-01-01",
				Language:     "language",
				Publisher:    "publisher",
				Author:       "author",
				Page_number:  123,
				Imagem:       "imagem",
				Owner:        "owner",
			},
		}),
	)

	req, err := http.NewRequest("GET", "/book", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctl.GetBooks())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"books":[{"id":"123","title":"title","category":"category","subtitle":"subtitle","description":"description","release_date":"2024-01-01","publisher":"publisher","language":"language","author":"author","page_number":123,"imagem":"imagem","owner":"owner"}]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestBookListWhenHasQuery(t *testing.T) {

	ctl := controllers.NewBookController(
		mock.NewBookRepositoryMock([]book.Book{
			{
				Id:           "123",
				Title:        "title",
				Category:     "category",
				Subtitle:     "subtitle",
				Description:  "description",
				Release_date: "2024-01-01",
				Language:     "language",
				Publisher:    "publisher",
				Author:       "author",
				Page_number:  123,
				Imagem:       "imagem",
				Owner:        "owner",
			},
		}),
	)

	req, err := http.NewRequest("GET", "/book?test=11", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctl.GetBooks())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"books":[{"id":"123","title":"title","category":"category","subtitle":"subtitle","description":"description","release_date":"2024-01-01","publisher":"publisher","language":"language","author":"author","page_number":123,"imagem":"imagem","owner":"owner"}]}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestBookGetById(t *testing.T) {

	ctl := controllers.NewBookController(
		mock.NewBookRepositoryMock([]book.Book{
			{
				Id:           "123",
				Title:        "title",
				Category:     "category",
				Subtitle:     "subtitle",
				Description:  "description",
				Release_date: "2024-01-01",
				Language:     "language",
				Publisher:    "publisher",
				Author:       "author",
				Page_number:  123,
				Imagem:       "imagem",
				Owner:        "owner",
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/book/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/book/{id}", ctl.GetBookById())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"123","title":"title","category":"category","subtitle":"subtitle","description":"description","release_date":"2024-01-01","publisher":"publisher","language":"language","author":"author","page_number":123,"imagem":"imagem","owner":"owner"}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestBookGetByIdNotFound(t *testing.T) {

	ctl := controllers.NewBookController(
		mock.NewBookRepositoryMock([]book.Book{
			{
				Id:           "123",
				Title:        "title",
				Category:     "category",
				Subtitle:     "subtitle",
				Description:  "description",
				Release_date: "2024-01-01",
				Language:     "language",
				Publisher:    "publisher",
				Author:       "author",
				Page_number:  123,
				Imagem:       "imagem",
				Owner:        "owner",
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/book/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/book/{id}", ctl.GetBookById())
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

func TestBookCreateWithImage(t *testing.T) {
	os.MkdirAll("public/book/", fs.ModePerm)

	defer os.RemoveAll("public")

	ctl := controllers.NewBookController(
		mock.NewBookRepositoryMock([]book.Book{}),
	)

	file, _ := os.Open("mock/image.png")

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "1710107242.png")

	io.Copy(part, file)

	writer.WriteField("title", "title")
	writer.WriteField("category", "category")
	writer.WriteField("subtitle", "subtitle")
	writer.WriteField("description", "description")
	writer.WriteField("release_date", "2024-01-01")
	writer.WriteField("language", "language")
	writer.WriteField("publisher", "publisher")
	writer.WriteField("author", "author")
	writer.WriteField("owner", "owner")
	writer.WriteField("page_number", "1")
	writer.WriteField("rate", "1")

	writer.Close()

	req, _ := http.NewRequest("POST", "/book", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	ctx := context.WithValue(req.Context(), "user", jwt.MapClaims{
		"user": "123",
		"exp":  time.Now().Add(time.Duration(time.Hour * 1)).Unix(),
	})

	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/book", ctl.CreateBook())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	book := book.Book{}
	json.Unmarshal(rr.Body.Bytes(), &book)

}

func TestBookCreateWithoutImage(t *testing.T) {
	os.MkdirAll("public/book/", fs.ModePerm)

	defer os.RemoveAll("public")

	ctl := controllers.NewBookController(
		mock.NewBookRepositoryMock([]book.Book{}),
	)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("title", "title")
	writer.WriteField("category", "category")
	writer.WriteField("subtitle", "subtitle")
	writer.WriteField("description", "description")
	writer.WriteField("release_date", "2024-01-01")
	writer.WriteField("language", "language")
	writer.WriteField("publisher", "publisher")
	writer.WriteField("author", "author")
	writer.WriteField("owner", "owner")
	writer.WriteField("page_number", "1")
	writer.WriteField("rate", "1")

	writer.Close()

	req, _ := http.NewRequest("POST", "/book", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	ctx := context.WithValue(req.Context(), "user", jwt.MapClaims{
		"user": "123",
		"exp":  time.Now().Add(time.Duration(time.Hour * 1)).Unix(),
	})

	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/book", ctl.CreateBook())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	book := book.Book{}
	json.Unmarshal(rr.Body.Bytes(), &book)

}

func TestBookUpdate(t *testing.T) {

	ctl := controllers.NewBookController(
		mock.NewBookRepositoryMock([]book.Book{
			{
				Id:           "123",
				Title:        "title",
				Category:     "category",
				Subtitle:     "subtitle",
				Description:  "description",
				Release_date: "2024-01-01",
				Language:     "language",
				Publisher:    "publisher",
				Author:       "author",
				Page_number:  123,
				Imagem:       "imagem",
				Owner:        "owner",
			},
		}),
	)

	rr := httptest.NewRecorder()

	content := book.Book{
		Title:       "Codigo Limpo",
		Page_number: 123,
	}

	jsonData, err := json.Marshal(content)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/book/123", bytes.NewReader([]byte(jsonData)))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", jwt.MapClaims{
		"user": "123",
		"exp":  time.Now().Add(time.Duration(time.Hour * 1)).Unix(),
	})

	req = req.WithContext(ctx)

	router := mux.NewRouter()
	router.HandleFunc("/book/{id}", ctl.UpdateBook())
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	book := book.Book{}
	json.Unmarshal(rr.Body.Bytes(), &book)

	compare := (strings.Compare(book.Title, "Codigo Limpo") == 0 &&
		book.Page_number == 123)

	if !compare {
		t.Errorf("handler returned unexpected body")
	}
}

func TestBookDelete(t *testing.T) {

	ctl := controllers.NewBookController(
		mock.NewBookRepositoryMock([]book.Book{
			{
				Id:           "123",
				Title:        "title",
				Category:     "category",
				Subtitle:     "subtitle",
				Description:  "description",
				Release_date: "2024-01-01",
				Language:     "language",
				Publisher:    "publisher",
				Author:       "author",
				Page_number:  123,
				Imagem:       "imagem",
				Owner:        "owner",
			},
		}),
	)

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("DELETE", "/book/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/book/{id}", ctl.Delete())
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
