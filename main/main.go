package main

import (
    "chat-server/chat"
    "chat-server/controller"
    "chat-server/domain/connectionpool"
    "chat-server/persistence"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "github.com/joho/godotenv"
    "log"
    "net/http"
    "os"
    "time"
)

var (
    upgrader       = websocket.Upgrader{}
    connectionPool = connectionpool.ConnectionPool{}
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

// This is used to avoid cors(request different domains) problem from the client
func corsHeader(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "*")
    c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")

    // When React calls an API, it first sends an OPTIONS request to detect if the API available
    // So return 204 whenever receive an OPTIONS request to avoid CORS error
    if c.Request.Method == "OPTIONS" {
        c.AbortWithStatus(204)
        return
    }
}

func main() {
    persistence.InitMongoDB()
    defer func() {
        persistence.PostMongoDB()
    }()

    hub := chat.NewHub()
    go hub.Run()

    r := gin.Default()
    r.Use(corsHeader)

    v1 := r.Group("/v1")
    ws := v1.Group("/ws")
    {
        ws.GET("/chat", func(c *gin.Context) {
            controller.ServeWs(hub, c)
        })
    }

    {
        v1.GET("/ping", func(c *gin.Context) {
            go func() {
                time.Sleep(time.Second * 10)
                fmt.Println("subroutine end...")
            }()
            c.JSON(http.StatusOK, gin.H{
                "message": "pong",
            })
        })
    }

    port := fmt.Sprintf(":%v", os.Getenv("PORT"))
    r.Run(port)
}
