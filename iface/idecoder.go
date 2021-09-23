package iface

import "net"

type IDecoder interface {
	Decode(conn net.Conn) (IMsg, error)
}
