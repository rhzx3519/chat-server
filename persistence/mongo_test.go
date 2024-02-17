package persistence

import (
    "context"
    "fmt"
    "github.com/stretchr/testify/assert"
    "go.mongodb.org/mongo-driver/bson"
    "testing"
    "time"
)

func TestInsert(t *testing.T) {
    InitMongoDB()
    defer func() {
        PostMongoDB()
    }()

    collection := Client.Database("testing").Collection("numbers")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})

    assert.ErrorIs(t, err, nil)
    fmt.Println(res)
}
