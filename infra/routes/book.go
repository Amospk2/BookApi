package routes

import (
	"api/infra/controllers"

	"github.com/gorilla/mux"
)

type BookRouter struct {
	controller *controllers.BookController
}

func (p *BookRouter) Load(mux *mux.Router) {
	mux.HandleFunc("/book", p.controller.GetBooks()).Methods("GET")
	mux.HandleFunc("/book/{id}", p.controller.GetBookById()).Methods("GET")
	mux.HandleFunc("/book/{id}", p.controller.UpdateBook()).Methods("PUT")
	mux.HandleFunc("/book/{id}", p.controller.Delete()).Methods("DELETE")
	mux.HandleFunc("/book", p.controller.CreateBook()).Methods("POST")

}

func NewBookRouter(
	controller *controllers.BookController,
) *BookRouter {
	return &BookRouter{
		controller: controller,
	}
}
