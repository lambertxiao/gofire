package iface

type IMsg interface {
	GetAction() string
	Serialize() ([]byte, error)
	Unserialize([]byte) error
	// Size() uint32
	// GetPayload() []byte
	// GetID() uint32
}

type IMsgCodec interface {
	Encode(msg IMsg) ([]byte, error)
	Decode(payload []byte) (IMsg, error)
}
