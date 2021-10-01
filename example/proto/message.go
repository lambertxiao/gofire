package proto

import "encoding/json"

type Message struct {
	MsgId  string
	Action string
	Body   map[string]interface{}
}

func (m *Message) GetID() string {
	return m.MsgId
}

func (m *Message) GetAction() string {
	return m.Action
}

func (m *Message) Serialize() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Message) Unserialize(data []byte) error {
	return json.Unmarshal(data, m)
}
