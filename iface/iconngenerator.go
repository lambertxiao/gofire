package iface

type IConnGenerator interface {
	Gen() (IConn, error)
}
