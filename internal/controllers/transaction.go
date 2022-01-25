package controllers

import (
	"encoding/json"
	"unchain/configs"
	"unchain/internal/response"
	"net/http"
	
	"go.mongodb.org/mongo-driver/mongo"
)

var transaction *mongo.Collection = configs.GetCollection(configs.DB, "transactions")


func GetTransactions() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		response := responses.TransactionResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "success"}}
		json.NewEncoder(rw).Encode(response)
	}
}