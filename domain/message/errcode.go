package message

type ErrCode int

const (
    COVERSATION ErrCode = iota
    ROOM_INFO

    DISCONNECTED = 10
)
