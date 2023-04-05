package proto

import (
	"gofire/core"
)

type CustomMsgCodec struct {
	IdLen int
}

func NewCustomMsgCodec() core.MsgCodec {
	c := &CustomMsgCodec{IdLen: 36}
	return c
}

func (c *CustomMsgCodec) Encode(msg core.Msg) ([]byte, error) {
	buf := []byte{}
	id := []byte(msg.GetID())
	buf = append(buf, id...)
	buf = append(buf, msg.GetPayload()...)
	return buf, nil
}

func (c *CustomMsgCodec) Decode(payload []byte) (core.Msg, error) {
	msg := new(Message)
	msg.MsgId = string(payload[:c.IdLen])
	msg.Payload = payload[c.IdLen:]
	return msg, nil
}
