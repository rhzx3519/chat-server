package domain

type MessageType int

const (
    INDIVIDUAL MessageType = iota
    CHANNEL
)

type Message interface {
    Type() MessageType
}

type IndividualMessage struct {
    Id        int64  `json:"id"`
    From      int64  `json:"from"`
    To        int64  `json:"to"`
    Content   string `json:"content"`
    CreatedAt int64  `json:"createdAt"`
    IsRead    bool   `json:"isRead"`
}

func (m *IndividualMessage) Type() MessageType {
    return INDIVIDUAL
}

type ChannelMessage struct {
    ChannelId int64 `json:"channelId"`
    Message
}

func (m *ChannelMessage) Type() MessageType {
    return CHANNEL
}
