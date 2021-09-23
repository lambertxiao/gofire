package iface

import "net"

type Request struct {
	Conn net.Conn
	Msg  IMsg
}

type IHandler interface {
	Do(req Request)
}
