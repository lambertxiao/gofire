package iface

type IConn interface {
	Handle()
	ReadLoop()
	WriteLoop()
	WriteMsg(IMsg)
}
