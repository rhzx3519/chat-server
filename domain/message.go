package domain

import (
    "chat-server/domain/serialnumber"
    "fmt"
    "time"
)

type MessageType int

const (
    PRIVATE MessageType = iota
    GROUP
)

type Message struct {
    SerialNo    int64       `json:"serialNo" bson:"serialNo"`
    From        string      `json:"from" bson:"from"`
    To          string      `json:"to" bson:"to"`
    Content     string      `json:"content" bson:"content"`
    CreatedAt   int64       `json:"createdAt" bson:"createdAt"`
    IsRead      bool        `json:"isRead" bson:"isRead"`
    MessageType MessageType `json:"messageType" bson:"messageType"`
}

func NewMessage(from, to string, content string) *Message {
    serialNo, err := serialnumber.NextSerialNo(from, to)
    if err != nil {
        fmt.Println(err)
    }
    return &Message{
        SerialNo:    serialNo,
        From:        from,
        To:          to,
        Content:     content,
        CreatedAt:   time.Now().Unix(),
        IsRead:      false,
        MessageType: GROUP,
    }
}
