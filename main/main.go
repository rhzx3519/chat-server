package main

import (
    "chat-server/domain"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "github.com/joho/godotenv"
    "log"
    "net/http"
    "os"
)

var (
    upgrader       = websocket.Upgrader{}
    connectionPool = domain.ConnectionPool{}
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

func echo(ctx *gin.Context) {
    w, r := ctx.Writer, ctx.Request
    c, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("upgrade:", err)
        return
    }
    fmt.Printf("Auth-User-No: %s", ctx.GetHeader("Auth-User-No"))

    defer c.Close()
    for {
        mt, message, err := c.ReadMessage()
        if err != nil {
            log.Println("read:", err)
            break
        }
        log.Printf("recv:%s", message)
        err = c.WriteMessage(mt, message)
        if err != nil {
            log.Println("write:", err)
            break
        }
    }
}

func main() {
    r := gin.Default()
    r.Use(corsHeader)

    ws := r.Group("/ws")
    v1 := ws.Group("/v1")
    {
        v1.GET("/ping", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{
                "message": "pong",
            })
        })

        v1.GET("/echo", echo)
    }

    port := fmt.Sprintf(":%v", os.Getenv("PORT"))
    r.Run(port)
}
