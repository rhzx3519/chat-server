package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const CollChannel = "channels"

type Channel struct {
	ID        primitive.ObjectID `bson:"_id"`
	Owner     string             `json:"owner"bson:"owner,omitempty"`
	Name      string             `json:"name"bson:"name,omitempty"`
	MaxMember int                `json:"max_member,omitempty"bson:"max_member,omitempty"`
}
