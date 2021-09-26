package iface

type IServer interface {
	AddRouter(actionId uint32, handler IHandler)
	GetActionHandler(actionId uint32) IHandler
	Listen(transProtocol TransProtocol) error
}
