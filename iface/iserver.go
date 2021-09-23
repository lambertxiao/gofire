package iface

type IServer interface {
	AddAction(actionId uint32)
	Listen() error
}
