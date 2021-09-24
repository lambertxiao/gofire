package iface

type IClient interface {
	Connect() error
	Send(IMsg) error
	OnMsg() <-chan IMsg
	SetMsgQueueSize(size int)
}
