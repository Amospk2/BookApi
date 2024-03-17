package routes

import (
	"api/infra/controllers"
	"api/infra/middleware"

	"github.com/gorilla/mux"
)

type AuthorRoute struct {
	controller *controllers.AuthorController
}

func (p *AuthorRoute) Load(mux *mux.Router) {
	mux.HandleFunc("/author", middleware.AuthenticationMiddleware(p.controller.GetAuthors())).Methods("GET")
	mux.HandleFunc("/author/{id}", middleware.AuthenticationMiddleware(p.controller.GetAuthorById())).Methods("GET")
	mux.HandleFunc("/author/{id}", middleware.AuthenticationMiddleware(p.controller.UpdateAuthor())).Methods("PUT")
	mux.HandleFunc("/author/{id}", middleware.AuthenticationMiddleware(p.controller.Delete())).Methods("DELETE")
	mux.HandleFunc("/author", middleware.AuthenticationMiddleware(p.controller.CreateAuthor())).Methods("POST")

}

func NewAuthorRouter(
	controller *controllers.AuthorController,
) *AuthorRoute {
	return &AuthorRoute{
		controller: controller,
	}
}
