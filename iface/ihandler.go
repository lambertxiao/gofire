package iface

type Request struct {
	Stream IStream
	Msg    IMsg
}

type IHandler interface {
	Do(req Request)
}
