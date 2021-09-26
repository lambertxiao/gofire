package gofire

import (
	"fmt"
	"gofire/iface"
	"log"
)

type FireServer struct {
	endpoint Endpoint
	routers  map[uint32]iface.IHandler

	// 连接生成器，server不关心具体conn底层的实现
	connG iface.IConnGenerator
	// 包编解码器
	pCodec iface.IPacketCodec
	// 消息编解码器
	mCodec iface.IMsgCodec
}

type Endpoint struct {
	Ip   string
	Port int
}

func (e Endpoint) String() string {
	return fmt.Sprintf("%s:%d", e.Ip, e.Port)
}

func NewServer(
	connG iface.IConnGenerator,
	pCodec iface.IPacketCodec,
	mCodec iface.IMsgCodec,
) iface.IServer {
	s := &FireServer{
		connG:   connG,
		pCodec:  pCodec,
		mCodec:  mCodec,
		routers: make(map[uint32]iface.IHandler),
	}

	return s
}

func (s *FireServer) Listen(transProtocol iface.TransProtocol) error {
	for {
		c := s.connG.Gen()
		stream := NewFireStream(c, s)
		go stream.Run()
	}
}

func (s *FireServer) AddRouter(actionId uint32, handler iface.IHandler) {
	log.Println("regist action id", actionId)
	s.routers[actionId] = handler
}

func (s *FireServer) GetActionHandler(actionId uint32) iface.IHandler {
	h, exist := s.routers[actionId]
	if !exist {
		return nil
	}

	return h
}
