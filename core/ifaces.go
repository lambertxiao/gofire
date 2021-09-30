package core

import "io"

type IClient interface {
	// Connect() error
	Send(IMsg) (IMsg, error)
	OnMsg() <-chan IMsg
	SetMsgQueueSize(size int)
}

type IConn interface {
	io.ReadWriteCloser
}

type IConnGenerator interface {
	Gen() (IConn, error)
}

type IConnPool interface {
	PutConn(conn IConn)
}

type Request struct {
	Stream IStream
	Msg    IMsg
}

type IHandler interface {
	Do(req Request)
}

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

type IProto interface {
	// 每种协议的head长度都是固定的
	GetHeadLen() uint32
	GetPayloadSize(head []byte) uint32
}

type IServer interface {
	RegistAction(action string, handler IHandler)
	GetHandler(action string) IHandler
	Listen() error
	// RegistProto(s IProto, c IProto)
}

type IStream interface {
	Flow()
	Write(IMsg)
	ReadLoop()
	WriteLoop()
}
