package chat

import (
    "chat-server/domain"
    "chat-server/network"
    "encoding/json"
    "github.com/gorilla/websocket"
)

type Chatter struct {
    client *network.Client
    hub    *Hub
    user   *domain.User
}

type ChatterOpt func(c *Chatter)

func WithConn(conn *websocket.Conn) ChatterOpt {
    return func(c *Chatter) {
        //client := chat.NewClient(hub, conn, make(chan []byte, 256))
        client := network.NewClient(
            network.WithConn(conn),
            network.WithSendBuffer(make(chan []byte, 256)))
        c.client = client
        network.WithReadCallback(c.Receive)(client)
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
    c := &Chatter{}
    for _, opt := range opts {
        opt(c)
    }
    return c
}

func (c *Chatter) Receive(content []byte) ([]byte, error) {
    message := domain.NewMessage(c.user.No, "group1", string(content))
    bytes, err := json.Marshal(message)
    if err != nil {
        return nil, err
    }
    c.hub.broadcast <- bytes
    return bytes, nil
}

func (c *Chatter) Send(message []byte) {
    c.client.Buffer() <- message
}

func (c *Chatter) Run() {
    defer func() {
        c.hub.unregister <- c
    }()
    // Allow collection of memory referenced by the caller by doing all work in
    // new goroutines.
    go c.client.ReadPump()
    go c.client.WritePump()
}

func (c *Chatter) Exit() {
    c.client.Clear()
}
