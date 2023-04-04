package core

import (
	"io"
	"time"
)

type MsgCB func(msg Msg, err error)

type Client interface {
	SyncSend(Msg) (Msg, error)
	AsyncSend(Msg, MsgCB) error
}

type IServer interface {
	RegistAction(action string, handler Handler)
	GetHandler(action string) Handler
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

type Request struct {
	Stream ISTransport
	Msg    Msg
}

type Handler interface {
	Do(req Request)
}

type Msg interface {
	GetID() string
	GetAction() string
	Serialize() ([]byte, error)
	Unserialize([]byte) error
	GetTimeout() time.Duration
	GetPriority() int
}

type MsgCodec interface {
	Encode(msg Msg) ([]byte, error)
	Decode(payload []byte) (Msg, error)
}

type IPacketCodec interface {
	Encode(payload []byte, w io.Writer) error
	Decode(r io.Reader) ([]byte, error)
}

type ISTransport interface {
	Flow()
	Write(msg Msg)
	ReadLoop()
	WriteLoop()
}

type ClientTransport interface {
	// Open()
	RoundTrip(msg Msg) (Msg, error)
}

type MsgQueue interface {
	Push(msg Msg)
	Pop() Msg
}
