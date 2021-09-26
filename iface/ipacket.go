package iface

import "io"

type IPacket interface {
	GetHeadLength() uint32
	UnPackHeadData(data []byte) error
	UnPackPayload(data []byte)
}

type IPacketCodec interface {
	// 将数据写入io流
	Encode(payload []byte, w io.Writer) error
	Decode(r io.Reader) ([]byte, error)
}
