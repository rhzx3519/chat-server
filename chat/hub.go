package chat

import log "github.com/sirupsen/logrus"

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
    // Registered clients.
    chatters map[*Chatter]bool

    // Inbound messages from the clients.
    broadcast chan []byte

    // Register requests from the clients.
    register chan *Chatter

    // Unregister requests from clients.
    unregister chan *Chatter
}

func NewHub() *Hub {
    return &Hub{
        broadcast:  make(chan []byte),
        register:   make(chan *Chatter),
        unregister: make(chan *Chatter),
        chatters:   make(map[*Chatter]bool),
    }
}

func (h *Hub) Register(chatter *Chatter) {
    h.register <- chatter
}

func (h *Hub) Run() {
    for {
        select {
        case chatter := <-h.register:
            {
                log.Debugf("%v joined chatroom", chatter.user.Nickname)
                h.chatters[chatter] = true
            }
        case chatter := <-h.unregister:
            if _, ok := h.chatters[chatter]; ok {
                delete(h.chatters, chatter)
                chatter.Exit()
                log.Debugf("%v left chatroom", chatter.user.Nickname)
            }
        case message := <-h.broadcast:
            for chatter := range h.chatters {
                chatter.Send(message)
            }
        }
    }
}

func (h *Hub) Clients() {

}
