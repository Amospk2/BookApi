package routes

import (
	"api/infra/controllers"
	"api/infra/middleware"

	"github.com/gorilla/mux"
)

type BookRouter struct {
	controller *controllers.BookController
}

func (p *BookRouter) Load(mux *mux.Router) {
	mux.HandleFunc("/book", middleware.AuthenticationMiddleware(p.controller.GetBooks())).Methods("GET")
	mux.HandleFunc("/book/{id}", middleware.AuthenticationMiddleware(p.controller.GetBookById())).Methods("GET")
	mux.HandleFunc("/book/{id}", middleware.AuthenticationMiddleware(p.controller.UpdateBook())).Methods("PUT")
	mux.HandleFunc("/book/{id}", middleware.AuthenticationMiddleware(p.controller.Delete())).Methods("DELETE")
	mux.HandleFunc("/book", middleware.AuthenticationMiddleware(p.controller.CreateBook())).Methods("POST")
}

func NewBookRouter(
	controller *controllers.BookController,
) *BookRouter {
	return &BookRouter{
		controller: controller,
	}
}
