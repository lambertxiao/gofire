package iface

type IServer interface {
	RegistAction(actionId uint32, handler IHandler)
	GetActionHandler(actionId uint32) IHandler
	Serve(transProtocol TransProtocol) error
}
