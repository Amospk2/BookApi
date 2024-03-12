package routes

import (
	"api/infra/controllers"

	"github.com/gorilla/mux"
)

type AuthorRoute struct {
	controller *controllers.AuthorController
}

func (p *AuthorRoute) Load(mux *mux.Router) {
	mux.HandleFunc("/author", p.controller.GetAuthors()).Methods("GET")
	mux.HandleFunc("/author/{id}", p.controller.GetAuthorById()).Methods("GET")
	mux.HandleFunc("/author/{id}", p.controller.UpdateAuthor()).Methods("PUT")
	mux.HandleFunc("/author/{id}", p.controller.Delete()).Methods("DELETE")
	mux.HandleFunc("/author", p.controller.CreateAuthor()).Methods("POST")

}

func NewAuthorRouter(
	controller *controllers.AuthorController,
) *AuthorRoute {
	return &AuthorRoute{
		controller: controller,
	}
}
