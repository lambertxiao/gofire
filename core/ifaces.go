package core

import (
	"io"
	"time"
)

type MsgCallback interface {
	Success(msg Msg)
	Timeout(msg Msg)
}

type MsgRecvCallback func(tp Transport, msg Msg)

type Client interface {
	SyncSend(Msg) (Msg, error)
	AsyncSend(Msg, MsgCallback) error
}

type Server interface {
	Listen() error
}

type Conn interface {
	io.ReadWriteCloser
}

type ConnGenerator interface {
	Gen() (Conn, error)
}

type ConnPool interface {
	PutConn(conn Conn)
}

type Msg interface {
	GetID() string
	GetPayload() []byte
	GetTimeout() time.Duration
	GetPriority() int
}

type MsgCodec interface {
	Encode(msg Msg) ([]byte, error)
	Decode(payload []byte) (Msg, error)
}

type PacketCodec interface {
	Encode(payload []byte, w io.Writer) error
	Decode(r io.Reader) ([]byte, error)
}

// 消息传输站
type Transport interface {
	Open()
	SendMsg(msg Msg) error
	SetMsgCB(cb MsgRecvCallback)
	IsActive() bool
}

type MsgQueue interface {
	Push(msg Msg)
	Pop() Msg
}
