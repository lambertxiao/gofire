package iface

type IProto interface {
	// 每种协议的head长度都是固定的
	GetHeadLen() uint32
	GetPayloadSize(head []byte) uint32
}
