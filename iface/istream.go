package iface

type IStream interface {
	Flow()
	Write(IMsg)
	ReadLoop()
	WriteLoop()
}
