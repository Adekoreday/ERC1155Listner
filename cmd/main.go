package main

import (
	log "github.com/sirupsen/logrus"
	"unchain/configs"
	"unchain/internal/routes"
	"unchain/internal/transaction"
	"unchain/internal/middlewares/common"
	"unchain/internal/middlewares/logger"
	"net/http"
	
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	log.SetFormatter(&log.JSONFormatter{})
    router.Use(common.CommonMiddleware)
	router.Use(logger.LoggingMiddleware)
	//run database
	configs.ConnectDB()

	go 	 transaction.Listen();
	//routes
	routes.Transaction(router) //add this

	go log.Fatal(http.ListenAndServe(":8000", router))


}