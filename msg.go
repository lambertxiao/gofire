package gofire

import (
	"bytes"
	"encoding/binary"
	"gofire/iface"
)

const HeaderLength uint32 = 8

type FireMsg struct {
	ID            string
	Action        string
	PayloadLength uint32
	Payload       []byte
}

func NewMsg() iface.IMsg {
	m := &FireMsg{}
	return m
}

func (h *FireMsg) SetPayload(payload []byte) {
	h.Payload = payload
}

func (h *FireMsg) GetPayloadLen() uint32 {
	return h.PayloadLength
}

func (h *FireMsg) LoadHead(data []byte) error {
	r := bytes.NewReader(data)
	err := binary.Read(r, binary.LittleEndian, &h.ID)
	if err != nil {
		return err
	}

	err = binary.Read(r, binary.LittleEndian, &h.PayloadLength)
	if err != nil {
		return err
	}

	err = binary.Read(r, binary.LittleEndian, &h.Action)
	if err != nil {
		return err
	}

	return nil
}
