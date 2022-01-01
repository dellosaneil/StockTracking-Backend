package main

import (
	"log"
	"net/http"

	"github.com/dellosaneil/stocktracking-backend/db"
	"github.com/dellosaneil/stocktracking-backend/handlers"
	"github.com/gorilla/mux"
)

func init() {
	db.ConnectDatabase()
	defer db.StockTrackingDb.Close()
}

func main() {
	r := mux.NewRouter()
	handlers.HandleRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
