package persistence

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "os"
    "time"
)

var (
    Client   *mongo.Client
    database string
)

func getenv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func InitMongoDB() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    var err error
    Client, err = mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v",
        getenv("DBHOST", "127.0.0.1"),
        getenv("DBPORT", "27017"))))
    if err != nil {
        panic(err)
    }
    database = getenv("DATABASE", "testing")
}

func PostMongoDB() {
    if err := Client.Disconnect(context.TODO()); err != nil {
        panic(err)
    }
}

func Database() *mongo.Database {
    return Client.Database(database)
}
