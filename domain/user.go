package domain

import "github.com/gorilla/websocket"

type User struct {
    Id   int64           `json:"id"`
    Conn *websocket.Conn `json:"conn"`
}

type Channel struct {
    Id          int64   `json:"id"`
    Subscribers []int64 `json:"subscribers"`
}

type InBox struct {
    UserId   int64      `json:"userId"`
    Messages []*Message `json:"messages"`
}

type ConnectionPool map[int64]*websocket.Conn
