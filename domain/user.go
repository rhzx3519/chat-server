package domain

import (
    "chat-server/persistence"
    "context"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "time"
)

const (
    collectionName = "Users"
)

type User struct {
    No       string `json:"no" bson:"no"`
    Email    string `json:"email" bson:"email"`
    Nickname string `json:"nickname" bson:"nickname"`
    Presence bool   `json:"presence" bson:"presence"`
}

type Channel struct {
    Id          string   `json:"id" bson:"id"`
    Subscribers []string `json:"subscribers" bson:"subscribers"`
}

type InBox struct {
    UserNo   string     `json:"userNo" bson:"userNo"`
    Messages []*Message `json:"messages" bson:"messages"`
}

func HeartBeat(userNo string) (err error) {
    coll := persistence.Database().Collection(collectionName)
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    filter := bson.D{{"no", userNo}}
    var u User
    err = coll.FindOne(ctx, filter).Decode(&u)
    if errors.Is(err, mongo.ErrNoDocuments) {
        _, err = coll.InsertOne(ctx, bson.D{
            {"no", userNo},
            {"presence", true},
        })
        return err
    } else if err != nil {
        return err
    }
    _, err = coll.UpdateOne(ctx, filter, bson.D{
        {
            "$set", bson.D{{"presence", false}},
        },
    })

    return
}
