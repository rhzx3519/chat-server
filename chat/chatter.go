package chat

import (
	"chat-server/domain"
	"chat-server/domain/msg"
	"chat-server/network"
	"encoding/json"
	"github.com/gorilla/websocket"
)

type ChatterStatus int

const (
	ACTIVE ChatterStatus = iota
	OFFLINE
)

type Chatter struct {
	client *network.Client
	hub    *Hub
	user   *domain.User
	Status ChatterStatus
}

type ChatterOpt func(c *Chatter)

func WithConn(conn *websocket.Conn) ChatterOpt {
	return func(c *Chatter) {
		client := network.NewClient(
			network.WithConn(conn),
			network.WithSendBuffer(make(chan []byte, 256)))
		c.client = client
		network.WithReadCallback(c.readHandler)(client)
		network.WithCloseCallback(c.closeHandler)(client)
	}
}

func WithUser(user *domain.User) ChatterOpt {
	return func(c *Chatter) {
		c.user = user
	}
}

func WithHub(hub *Hub) ChatterOpt {
	return func(c *Chatter) {
		c.hub = hub
	}
}

func NewChatter(opts ...ChatterOpt) *Chatter {
	c := &Chatter{Status: ACTIVE}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Chatter) readHandler(content []byte) {
	message := msg.NewMessage(content)
	message.From = c.user
	message.To = &domain.Channel{
		Name:      ROOM_NO,
		MaxMember: 10,
	}
	message.MsgCode = msg.GROUP_CONVERSATION
	c.hub.broadcast <- message
}

func (c *Chatter) closeHandler() {
	c.hub.unregister <- c
}

func (c *Chatter) Send(message *msg.Message) {
	bytes, _ := json.Marshal(message)
	c.client.Buffer() <- bytes
}

func (c *Chatter) Run() {
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go c.client.ReadPump()
	go c.client.WritePump()
}
