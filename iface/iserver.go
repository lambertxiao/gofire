package iface

type IServer interface {
	RegistAction(action string, handler IHandler)
	GetHandler(action string) IHandler
	Listen() error
	// RegistProto(s IProto, c IProto)
}
