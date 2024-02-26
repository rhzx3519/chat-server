package msg

type ErrCode int

const (
	DISCONNECTED = 10
)

type MsgCode int

const (
	PRIVATE_COVERSATION MsgCode = iota
	GROUP_CONVERSATION
	GROUP_INFO
)
