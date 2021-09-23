package iface

type IPacket interface {
	GetHeadLength() uint32
	UnPackHeadData(data []byte) error
	UnPackPayload(data []byte)
}
