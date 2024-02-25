package serialno

import (
	"chat-server/persistence"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	coll = "serialnos"
)

type SerialNo struct {
	ID   primitive.ObjectID `bson:"_id"`
	From string             `bson:"from,omitempty"`
	To   string             `bson:"to,omitempty"`
	Next int64              `bson:"next,omitempty"`
}

func NextSerialNo(from, to string) (int64, error) {
	coll := persistence.Database().Collection(coll)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"from", from}, {"to", to}}
	var r SerialNo
	err := coll.FindOne(ctx, filter).Decode(&r)

	if errors.Is(err, mongo.ErrNoDocuments) {
		_, err = coll.InsertOne(ctx, bson.D{
			{"from", from},
			{"to", to},
			{"next", 2},
		})
		if err != nil {
			return 0, err
		}
		return 1, nil
	} else if err != nil {
		return 0, err
	}

	_, err = coll.UpdateOne(ctx, filter, bson.D{
		{
			"$set", bson.D{{"next", r.Next + 1}},
		},
	})
	if err != nil {
		return 0, err
	}
	return r.Next, nil
}
