package iface

import "io"

type IConn interface {
	io.ReadWriteCloser
	// Handle()
	// ReadLoop()
	// WriteLoop()
	// WriteMsg(IMsg)
}
