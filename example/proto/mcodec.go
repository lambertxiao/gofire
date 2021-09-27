package proto

import "gofire/iface"

type MsgCodec struct{}

func NewMsgCodec() iface.IMsgCodec {
	c := &MsgCodec{}
	return c
}

func (c *MsgCodec) Encode(msg iface.IMsg) ([]byte, error) {
	payload, err := msg.Serialize()
	if err != nil {
		return payload, err
	}

	return payload, nil
}

func (c *MsgCodec) Decode(payload []byte) (iface.IMsg, error) {
	msg := new(Message)
	err := msg.Unserialize(payload)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
