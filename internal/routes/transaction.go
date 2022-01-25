package routes

import (
	"unchain/internal/controllers"

	"github.com/gorilla/mux"
)

func Transaction(router *mux.Router) {
	router.HandleFunc("/transaction", controllers.GetTransactions()).Methods("GET")
}