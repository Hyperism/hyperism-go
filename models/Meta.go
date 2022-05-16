package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meta struct {
    ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Owner        string             `json:"owner,omitempty" bson:"owner,omitempty"`
    Price        int64              `json:"price,omitempty" bson:"price,omitempty"`
    Minter       string             `json:"minter,omitempty" bson:"minter,omitempty"`
    Title        string             `json:"title,omitempty" bson:"title,omitempty"`
    Description  string             `json:"description,omitempty" bson:"description,omitempty"`
    IpfsHash     string             `json:"ipfshash,omitempty" bson:"ipfshash,omitempty"`
    MintDate     time.Time          `json:"mintdate,omitempty" bson:"mintdate,omitempty"`
    
}