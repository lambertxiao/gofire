package iface

type IMsg interface {
	LoadHead(data []byte) error
	GetPayloadLen() uint32
	SetPayload(data []byte)
}
