package msg

import (
	"chat-server/domain"
	"chat-server/persistence"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	collName = "messages"
)

type ErrCode int

type MsgCode int

const (
	PRIVATE_COVERSATION MsgCode = iota
	GROUP_CONVERSATION
	GROUP_INFO
)

type Message struct {
	SerialNo  int64           `json:"serial_no" bson:"serial_no"`
	From      *domain.User    `json:"from" bson:"from"`
	To        *domain.Channel `json:"to" bson:"to"`
	Content   string          `json:"content" bson:"content"`
	CreatedAt int64           `json:"created_at" bson:"created_at"`
	IsRead    bool            `json:"is_read" bson:"is_read"`
	MsgCode   MsgCode         `json:"msg_code" bson:"msg_code"`
	ErrCode   ErrCode         `json:"err_code" bson:"err_code"`
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

func List() []*Message {
	coll := persistence.Database().Collection(collName)
	// Creates a query filter to match documents in which the "cuisine"
	// is "Italian"
	filter := bson.D{{}}

	// Retrieves documents that match the query filer
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	// end find

	var results []*Message
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	//// Prints the results of the find operation as structs
	//for _, result := range results {
	//	cursor.Decode(&result)
	//	output, err := json.MarshalIndent(result, "", "    ")
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("%s\n", output)
	//}
	return results
}
