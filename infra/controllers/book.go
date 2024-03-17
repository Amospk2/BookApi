package controllers

import (
	"api/domain/book"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type BookController struct {
	repository book.BookRepository
}

func replaceBookForBookRequest(bookRequest *book.Book, book *book.Book) {

	if len(bookRequest.Title) > 0 {
		book.Title = bookRequest.Title
	}

	if len(bookRequest.Category) > 0 {
		book.Category = bookRequest.Category
	}

	if len(bookRequest.Subtitle) > 0 {
		book.Subtitle = bookRequest.Subtitle
	}

	if len(bookRequest.Description) > 0 {
		book.Description = bookRequest.Description
	}

	if _, err := time.Parse("2006-01-02", bookRequest.Release_date); err != nil {
		book.Release_date = bookRequest.Release_date
	}

	if len(bookRequest.Publisher) > 0 {
		book.Publisher = bookRequest.Publisher
	}

	if len(bookRequest.Language) > 0 {
		book.Language = bookRequest.Language
	}

	if len(bookRequest.Author) > 0 {
		book.Author = bookRequest.Author
	}

	if fmt.Sprintf("%T", bookRequest.Page_number) == "int" && bookRequest.Page_number > 0 {
		book.Page_number = bookRequest.Page_number
	}

	if fmt.Sprintf("%T", bookRequest.Rate) == "float" && bookRequest.Rate > 0 {
		book.Rate = bookRequest.Rate
	}

	if len(bookRequest.Owner) > 0 {
		book.Owner = bookRequest.Owner
	}
}

func uploadImage(r *http.Request, book *book.Book) error {
	file, handler, err := r.FormFile("imagem")

	if err != nil {
		return nil
	}

	defer file.Close()

	imgSplited := strings.Split(handler.Filename, ".")

	book.Imagem = "public/book/" + fmt.Sprint(time.Now().Unix()) +
		"." + imgSplited[len(imgSplited)-1]

	f, err := os.OpenFile(
		book.Imagem,
		os.O_WRONLY|os.O_CREATE,
		0666)

	if err != nil {
		return err
	}

	defer f.Close()
	_, _ = io.Copy(f, file)

	return nil
}

func (c *BookController) GetBooks() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			books, err := c.repository.Get()
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{"books": books})
		},
	)
}

func (c *BookController) GetBookById() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			book, err := c.repository.GetById(vars["id"])

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
		},
	)
}

func (c *BookController) Delete() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			if _, err := c.repository.GetById(vars["id"]); err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			if err := c.repository.Delete(vars["id"]); err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
		},
	)
}

func (c *BookController) UpdateBook() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			var bookRequest book.Book

			book, err := c.repository.GetById(vars["id"])

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			if err = json.NewDecoder(r.Body).Decode(&bookRequest); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			replaceBookForBookRequest(&bookRequest, &book)

			if err = c.repository.Update(book); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
		},
	)
}

func (c *BookController) CreateBook() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var book book.Book

			book.Title = r.FormValue("title")
			book.Category = r.FormValue("category")
			book.Author = r.FormValue("author")
			book.Description = r.FormValue("description")
			book.Owner = r.FormValue("owner")
			book.Subtitle = r.FormValue("subtitle")
			book.Language = r.FormValue("language")
			book.Publisher = r.FormValue("publisher")
			book.Release_date = r.FormValue("release_date")

			pages, err := strconv.ParseInt(r.FormValue("page_number"), 10, 32)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			book.Page_number = int32(pages)

			rate, err := strconv.ParseFloat(r.FormValue("rate"), 32)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			book.Rate = float32(rate)

			if !book.Valid() {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			if err := uploadImage(r, &book); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if err := c.repository.Create(book); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(book)
		},
	)
}

func NewBookController(
	r book.BookRepository,
) *BookController {
	return &BookController{
		repository: r,
	}
}
