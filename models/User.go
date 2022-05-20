package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email              string             `json:"email" bson:"email,omitempty"`
	Password           string             `json:"password" bson:"password,omitempty"`
	CreatedAt          time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt          time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`
	Username           string             `json:"username" bson:"username,omitempty"`
}
