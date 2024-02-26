package controller

import (
	"chat-server/chat"
	"chat-server/domain"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func extractForwardedUser(c *gin.Context) (*domain.User, error) {
	userJson := c.GetHeader("X-Forwarded-User")
	if userJson == "" {
		return nil, fmt.Errorf("X-Forwarded-User is empty.")
	}
	var u domain.User
	if err := json.Unmarshal([]byte(userJson), &u); err != nil {
		return &u, err
	}
	return &u, nil
}

// serveWs handles websocket requests from the peer.
func ServeWs(hub *chat.Hub, c *gin.Context) {
	user, err := extractForwardedUser(c)

	if err != nil {
		log.WithError(err).Error("failed to parse user json")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.WithError(err).Error("failed to upgrade http connection to websocket")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	chatter := chat.NewChatter(
		chat.WithConn(conn),
		chat.WithHub(hub),
		chat.WithUser(user),
	)

	hub.Register(chatter)
}
