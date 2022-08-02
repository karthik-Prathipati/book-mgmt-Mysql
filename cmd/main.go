package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karthik-Prathipati/book-mgmt-Mysql/pkg/routes"
)

func main() {
	mux := mux.NewRouter()
	routes.RegisterBookRoutes(mux)
	http.Handle("/", mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
