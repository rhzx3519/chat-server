package connectionpool

import "github.com/gorilla/websocket"

type ConnectionPool map[string]*websocket.Conn
