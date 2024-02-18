package msg

import (
    "chat-server/persistence"
    "context"
    "time"
)

type MessageType int

const (
    PRIVATE MessageType = iota
    GROUP
)

const (
    collName = "messages"
)

type Message struct {
    SerialNo    int64       `json:"serialNo" bson:"serialNo"`
    From        string      `json:"from" bson:"from"`
    To          string      `json:"to" bson:"to"`
    Content     []byte      `json:"content" bson:"content"`
    CreatedAt   int64       `json:"createdAt" bson:"createdAt"`
    IsRead      bool        `json:"isRead" bson:"isRead"`
    MessageType MessageType `json:"messageType" bson:"messageType"`
    MsgCode     MsgCode     `json:"msgCode" bson:"msgCode"`
    ErrCode     ErrCode     `json:"errCode" bson:"errCode"`
}

func NewMessage(content []byte) *Message {
    return &Message{
        Content:     content,
        CreatedAt:   time.Now().Unix(),
        IsRead:      false,
        MessageType: GROUP,
    }
}

func Save(msg *Message) (err error) {
    coll := persistence.Database().Collection(collName)
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err = coll.InsertOne(ctx, msg)
    return
}
