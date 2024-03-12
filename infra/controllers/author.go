package controllers

import (
	"api/domain/author"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type AuthorController struct {
	repository author.AuthorRepository
}

func (c *AuthorController) GetAuthors() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			Authors, _ := c.repository.Get()
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{"authors": Authors})
		},
	)
}

func (c *AuthorController) GetAuthorById() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			author, err := c.repository.GetById(vars["id"])

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(author)
		},
	)
}

func (c *AuthorController) Delete() http.HandlerFunc {
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

func (c *AuthorController) UpdateAuthor() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			var AuthorRequest author.Author

			Author, err := c.repository.GetById(vars["id"])

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			if err = json.NewDecoder(r.Body).Decode(&AuthorRequest); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			if len(AuthorRequest.Name) != 0 && AuthorRequest.Name != "" {
				Author.Name = AuthorRequest.Name
			}

			if err = c.repository.Update(Author); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Author)
		},
	)
}

func (c *AuthorController) CreateAuthor() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var author author.Author

			err := json.NewDecoder(r.Body).Decode(&author)

			if err != nil || !author.Valid() {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			author.Id = uuid.NewString()

			if err = c.repository.Create(author); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			w.Header().Set("Location", "/author/"+author.Id)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(author)
		},
	)
}

func NewAuthorController(
	r author.AuthorRepository,
) *AuthorController {
	return &AuthorController{
		repository: r,
	}
}
