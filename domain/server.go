package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const CollSever = "servers"

type Server struct {
	ID    primitive.ObjectID `bson:"_id"`
	Owner string             `bson:"owner,omitempty"`
	Name  string             `bson:"name,omitempty"`
}
