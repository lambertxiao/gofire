package proto

import (
	"time"
)

type Message struct {
	MsgId   string
	Payload []byte
}

func (m *Message) GetID() string {
	return m.MsgId
}

func (m *Message) GetPayload() []byte {
	return m.Payload
}

func (m *Message) GetTimeout() time.Duration {
	return 5 * time.Second
}

func (m *Message) GetPriority() int {
	return 0
}
