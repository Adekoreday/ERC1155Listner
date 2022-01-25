package main

import (
	"log"
	"unchain/configs"
	"unchain/internal/routes"
	"unchain/internal/transaction"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//run database
	configs.ConnectDB()
    transaction.Listen();
	//routes
	routes.Transaction(router) //add this

	log.Fatal(http.ListenAndServe(":8000", router))
}