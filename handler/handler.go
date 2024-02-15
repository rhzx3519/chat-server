package handler

import (
    "chat-server/domain"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "log"
)

var (
    upgrader       = websocket.Upgrader{}
    connectionPool = domain.ConnectionPool{}
)

func HandleMessage(ctx *gin.Context) {
    w, r := ctx.Writer, ctx.Request
    c, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("upgrade:", err)
        return
    }

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
