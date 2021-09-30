package core

import "io"

type IClient interface {
	Send(IMsg) (IMsg, error)
}

type IServer interface {
	RegistAction(action string, handler IHandler)
	GetHandler(action string) IHandler
	Listen() error
}

type IChannel interface {
	io.ReadWriteCloser
}

type IChannelGenerator interface {
	Gen() (IChannel, error)
}

type IMsgGenerator interface {
	Gen() IMsg
}

type IChannelPool interface {
	PutConn(conn IChannel)
}

type Request struct {
	Stream ISTransport
	Msg    IMsg
}

type IHandler interface {
	Do(req Request)
}

type IMsg interface {
	GetAction() string
	Serialize() ([]byte, error)
	Unserialize([]byte) error
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

type ISTransport interface {
	Flow()
	Write(msg IMsg)
	ReadLoop()
	WriteLoop()
}

// transport负责管理所有的stream
type ITransport interface {
	RoundTrip(msg IMsg) (IMsg, error)
}
