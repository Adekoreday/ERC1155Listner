package main

import (
	"log"
	"unchain/configs"
	"unchain/internal/routes"
	"unchain/internal/transaction"
	"unchain/internal/middlewares"
	"net/http"
	
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
    router.Use(common.CommonMiddleware)
	//run database
	configs.ConnectDB()

	go 	 transaction.Listen();
	//routes
	routes.Transaction(router) //add this

	go log.Fatal(http.ListenAndServe(":8000", router))


}