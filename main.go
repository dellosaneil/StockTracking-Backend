package main

import (
	"log"
	"net/http"

	"github.com/dellosaneil/stocktracking-backend/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	handlers.HandleRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
