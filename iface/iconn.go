package iface

type IConn interface {
	Handle()
	ReadLoop()
	WriteLoop()
	GetMsg() (IMsg, error)
}
