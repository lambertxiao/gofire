package iface

type ITransport interface {
	RegistClientProto()
	RegistServerProto()
}
