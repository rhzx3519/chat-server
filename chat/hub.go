package chat

import (
	"chat-server/domain"
	"chat-server/domain/msg"
	"chat-server/domain/serialno"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	chatters map[*Chatter]bool

	// Inbound messages from the clients.
	broadcast chan *msg.Message

	// Register requests from the clients.
	register chan *Chatter

	// Unregister requests from clients.
	unregister chan *Chatter
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan *msg.Message),
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
				for existChatter := range h.chatters {
					if existChatter.user.No == chatter.user.No {
						existChatter.Status = OFFLINE
						// close the connection to terminal goroutine
						close(existChatter.client.Done)
						break
					}
				}

				log.Debugf("%v joined chatroom", chatter.user.Nickname)
				h.chatters[chatter] = true
				chatter.Run()
				h.doBroadcast(h.packageRoomInfo())
			}
		case chatter := <-h.unregister:
			if _, ok := h.chatters[chatter]; ok {
				log.Debugf("%v left chatroom", chatter.user.Nickname)
				delete(h.chatters, chatter)
				h.doBroadcast(h.packageRoomInfo())
			}
		case message := <-h.broadcast:
			h.doBroadcast(message)
		}
	}
}

func (h *Hub) packageRoomInfo() *msg.Message {
	var users []*domain.User
	for chatter := range h.chatters {
		users = append(users, chatter.user)
	}
	bs, _ := json.Marshal(users)
	message := msg.NewMessage(bs)
	message.MsgCode = msg.GROUP_INFO
	return message
}

func (h *Hub) doBroadcast(message *msg.Message) {
	var err error
	if message.MsgCode == msg.GROUP_CONVERSATION {
		message.SerialNo, err = serialno.NextSerialNo(message.From.No, message.To.(*domain.Channel).Name)
		if err != nil {
			log.WithError(err).Error("failed to generate serial no.")
		}
	}

	msg.Save(message)
	for chatter := range h.chatters {
		if chatter.Status == ACTIVE {
			chatter.Send(message)
		}
	}
}
