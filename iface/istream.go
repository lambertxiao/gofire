package iface

type IStream interface {
	Run()
	Write(IMsg)
	ReadLoop()
}
