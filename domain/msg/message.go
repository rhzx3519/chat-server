package msg

import (
	"chat-server/domain"
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
	SerialNo  int64        `json:"serial_no" bson:"serial_no"`
	From      *domain.User `json:"from" bson:"from"`
	To        any          `json:"to" bson:"to"`
	Content   string       `json:"content" bson:"content"`
	CreatedAt int64        `json:"created_at" bson:"created_at"`
	IsRead    bool         `json:"is_read" bson:"is_read"`
	MsgCode   MsgCode      `json:"msg_code" bson:"msg_code"`
	ErrCode   ErrCode      `json:"err_code" bson:"err_code"`
}

func NewMessage(content []byte) *Message {
	return &Message{
		Content:   string(content),
		CreatedAt: time.Now().Unix(),
		IsRead:    false,
		MsgCode:   GROUP_CONVERSATION,
	}
}

func Save(msg *Message) (err error) {
	coll := persistence.Database().Collection(collName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = coll.InsertOne(ctx, msg)
	return
}
