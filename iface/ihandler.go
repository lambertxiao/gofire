package iface

type Request struct {
	Conn IConn
	Msg  IMsg
}

type IHandler interface {
	Do(req Request)
}
