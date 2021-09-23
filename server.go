package gofire

import (
	"fmt"
	"gofire/iface"
	"log"
	"net"
)

type FireServer struct {
	ip      string
	port    int
	network string
}

func NewServer(ip string, port int) iface.IServer {
	s := &FireServer{
		ip:      ip,
		port:    port,
		network: "tcp4",
	}
	return s
}

func (s *FireServer) Listen() error {
	addr, err := net.ResolveTCPAddr(s.network, fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP(s.network, addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println("accept tcp err", err)
			continue
		}

		log.Println("get conn from addr ", conn.RemoteAddr().String())
		fconn := NewFireConn(conn)
		go fconn.Handle()
	}
}

func (s *FireServer) HandleConn(conn net.Conn) {

}
