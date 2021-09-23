package iface

type IClient interface {
	SyncSend(IMsg) ([]byte, error)
	Connect() error
}
