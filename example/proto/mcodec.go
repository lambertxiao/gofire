package proto

import (
	"gofire/core"
)

type CustomMsgCodec struct{}

func NewCustomMsgCodec() core.MsgCodec {
	c := &CustomMsgCodec{}
	return c
}

func (c *CustomMsgCodec) Encode(msg core.Msg) ([]byte, error) {
	payload, err := msg.Serialize()
	if err != nil {
		return payload, err
	}

	return payload, nil
}

func (c *CustomMsgCodec) Decode(payload []byte) (core.Msg, error) {
	msg := new(Message)
	err := msg.Unserialize(payload)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
