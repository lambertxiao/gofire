package gofire

import (
	"fmt"
	"gofire/iface"
	"log"
	"net"
)

type FireServer struct {
	ip             string
	port           int
	network        string
	actionHandlers map[uint32]iface.IHandler
}

func NewServer(ip string, port int) iface.IServer {
	s := &FireServer{
		ip:             ip,
		port:           port,
		network:        "tcp4",
		actionHandlers: make(map[uint32]iface.IHandler),
	}
	return s
}

func (s *FireServer) Serve(transProtocol iface.TransProtocol) error {
	addr, err := net.ResolveTCPAddr(s.network, fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP(s.network, addr)
	if err != nil {
		return err
	}

	log.Printf("serve on %s:%d\n", s.ip, s.port)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println("accept tcp err", err)
			continue
		}

		log.Println("get conn from addr ", conn.RemoteAddr().String())
		fconn := NewFireConn(conn, s)
		go fconn.Handle()
	}
}

func (s *FireServer) RegistAction(actionId uint32, handler iface.IHandler) {
	log.Println("regist action id", actionId)
	s.actionHandlers[actionId] = handler
}

func (s *FireServer) GetActionHandler(actionId uint32) iface.IHandler {
	h, exist := s.actionHandlers[actionId]
	if !exist {
		return nil
	}

	return h
}
