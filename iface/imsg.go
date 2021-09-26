package iface

type IMsg interface {
	UnpackHead(headData []byte) error
	Pack() ([]byte, error)
	GetPayloadLen() uint32
	SetPayload(data []byte)
	GetAction() uint32
	GetPayload() []byte
	GetID() uint32
}

type IMsgCodec interface {
	Encode(msg IMsg) []byte
	Decode(payload []byte) IMsg
}
