package iface

type IClient interface {
	Send(IMsg) error
	Connect() error
}
