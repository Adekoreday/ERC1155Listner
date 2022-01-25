package controllers

import (
	"encoding/json"
	"log"
	"unchain/configs"
	"unchain/internal/response"
	"net/http"
	"unchain/internal/models"
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var transactionCollection *mongo.Collection = configs.GetCollection(configs.DB, "transaction")


func GetTransactions() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var results []*models.Transaction
	//	var findQuery []bson.M
		if r.URL.Query().Get("sender") != "" {
			fmt.Printf(`this is the value %s\n`, r.URL.Query().Get("sender"))
			
		}
		defer cancel()
        curr, err := transactionCollection.Find(ctx, bson.D{{}})
		
		for curr.Next(ctx) {
			var elem models.Transaction
			err := curr.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}
	
			results = append(results, &elem)
		} 

		if err != nil {
			log.Fatal(err)
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.TransactionResponse{Status: http.StatusOK, Message: "success",Data: map[string]interface{}{"data": results}}
		json.NewEncoder(rw).Encode(response)
	}
}