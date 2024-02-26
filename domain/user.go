package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollUser = "users"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	No       string             `json:"no" bson:"no"`
	Email    string             `json:"email" bson:"email"`
	Nickname string             `json:"nickname" bson:"nickname"`
	Fullname string             `json:"fullname"bson:"fullname"`
}
