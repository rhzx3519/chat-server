package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const CollChannel = "channels"

type Channel struct {
	ID        primitive.ObjectID `bson:"_id"`
	Owner     string             `bson:"owner,omitempty"`
	Name      string             `bson:"name,omitempty"`
	MaxMember int                `bson:"max_member,omitempty"`
}
