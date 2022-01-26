package models
import (
	"time"
)

type Transaction struct {
	TokenId    string             `bson:"tokenId,omitempty" validate:"required"`
	Sender     string             `bson:"sender,omitempty" validate:"required"`
	Receiver string             `bson:"receiver,omitempty" validate:"required"`
	Token    string             `bson:"token,omitempty" validate:"required"`
	SenderBal string         `bson:"senderbal,omitempty" validate:"required"`
	ReceiverBal string       `bson:"receiverbal,omitempty" validate:"required"`
	CreatedAt time.Time      `bson:"createdAt"`
}