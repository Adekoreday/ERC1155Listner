package controllers

import (
	"encoding/json"
	"unchain/configs"
	"unchain/internal/response"
	"net/http"
	"unchain/internal/models"
	"context"
	"time"
	"strconv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var transactionCollection *mongo.Collection = configs.GetCollection(configs.DB, "transaction")


func GetTransactions() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var results []*models.Transaction
		findQuery := make(bson.D, 2)

		opts := options.Find().SetLimit(100)

		if r.URL.Query().Get("createdAt") != "" {
            str := r.URL.Query().Get("createdAt")
            t, err := time.Parse(time.RFC3339, str)
			if err== nil {
				findQuery = append(findQuery, bson.E{"createdAt", bson.M{"$gte": primitive.NewDateTimeFromTime(t)}})
			}
		}

		if r.URL.Query().Get("sender") != "" {
			findQuery = append(findQuery, bson.E{"sender", r.URL.Query().Get("sender")})
		}

		if r.URL.Query().Get("receiver") != "" {
			findQuery = append(findQuery, bson.E{"receiver", r.URL.Query().Get("receiver")})
		}

		if r.URL.Query().Get("limit") != "" {
			intVar, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
			if err == nil {
				opts = options.Find().SetLimit(intVar)
			}
		}
		defer cancel()
        curr, err := transactionCollection.Find(ctx, findQuery, opts)
		
		for curr.Next(ctx) {
			var elem models.Transaction
			err := curr.Decode(&elem)
			if err == nil {
				results = append(results, &elem)
			}
		} 

		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			response := responses.TransactionResponse{Status: http.StatusNotFound, Message: "no data", Data: map[string]interface{}{"data": nil}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.TransactionResponse{Status: http.StatusOK, Message: "success",Data: map[string]interface{}{"data": results}}
		json.NewEncoder(rw).Encode(response)
	}
}