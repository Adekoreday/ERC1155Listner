package models
import (
	"time"
)

type Transaction struct {
	TokenId    string             `json:"tokenId,omitempty" validate:"required"`
	Sender     string             `json:"sender,omitempty" validate:"required"`
	Receiver string             `json:"receiver,omitempty" validate:"required"`
	Token    string             `json:"token,omitempty" validate:"required"`
	Operator string          `json:"operator,omitempty" validate:"required"`
	CreatedAt time.Time      `json:"createdAt"`
}