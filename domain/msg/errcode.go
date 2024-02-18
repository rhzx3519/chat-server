package msg

type ErrCode int

const (
    DISCONNECTED = 10
)

type MsgCode int

const (
    COVERSATION MsgCode = iota
    ROOM_INFO
)
