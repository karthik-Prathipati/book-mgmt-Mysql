package routes

import (
	"github.com/gorilla/mux"
	"github.com/karthik-Prathipati/book-mgmt-Mysql/pkg/controllers"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")

}
