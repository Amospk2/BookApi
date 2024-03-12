package controllers

import (
	"api/domain/category"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CategoryController struct {
	repository category.CategoryRepository
}

func (c *CategoryController) GetCategorys() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			categorys, err := c.repository.Get()
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{"categorys": categorys})
		},
	)
}

func (c *CategoryController) GetCategoryById() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			category, err := c.repository.GetById(vars["id"])

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(category)
		},
	)
}

func (c *CategoryController) Delete() http.HandlerFunc {
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

func (c *CategoryController) UpdateCategory() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			var categoryRequest category.Category

			category, err := c.repository.GetById(vars["id"])

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			if err = json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			if len(categoryRequest.Name) != 0 && categoryRequest.Name != "" {
				category.Name = categoryRequest.Name
			}

			if err = c.repository.Update(category); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(category)
		},
	)
}

func (c *CategoryController) CreateCategory() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var category category.Category

			err := json.NewDecoder(r.Body).Decode(&category)

			if err != nil || !category.Valid() {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			if err = c.repository.Create(category); err != nil {
				log.Fatal(err)
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(category)
		},
	)
}

func NewCategoryController(
	r category.CategoryRepository,
) *CategoryController {
	return &CategoryController{
		repository: r,
	}
}
