package proto

import "encoding/json"

type Message struct {
	Head MessageHead
	Body map[string]interface{}
}

type MessageHead struct {
	MsgId  string
	Action string
}

func (m Message) GetAction() string {
	return m.Head.Action
}

func (m Message) Serialize() ([]byte, error) {
	return json.Marshal(m)
}

func (m Message) Unserialize(data []byte) error {
	return json.Unmarshal(data, &m)
}

func (m Message) Size() uint32 {
	return 0
}
