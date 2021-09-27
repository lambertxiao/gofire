package iface

type IConnPool interface {
	PutConn(conn IConn)
}
