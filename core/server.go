package core

import (
	"fmt"
	"log"
)

type FireServer struct {
	routers map[string]IHandler
	// 连接生成器，server不关心具体conn底层的实现
	connG IConnGenerator
	// 包编解码器
	pcodec IPacketCodec
	// 消息编解码器
	mcodec IMsgCodec
}

type Endpoint struct {
	Ip   string
	Port int
}

func (e Endpoint) String() string {
	return fmt.Sprintf("%s:%d", e.Ip, e.Port)
}

func NewServer(
	connG IConnGenerator,
	pcodec IPacketCodec,
	mcodec IMsgCodec,
) IServer {
	s := &FireServer{
		connG:   connG,
		pcodec:  pcodec,
		mcodec:  mcodec,
		routers: make(map[string]IHandler),
	}

	return s
}

func (s *FireServer) Listen() error {
	log.Println("server listening...")
	for {
		c, err := s.connG.Gen()
		if err != nil {
			return err
		}

		stream := NewServerStream(c, s)
		go stream.Flow()
	}
}

func (s *FireServer) RegistAction(action string, handler IHandler) {
	log.Println("regist action id", action)
	s.routers[action] = handler
}

func (s *FireServer) GetHandler(action string) IHandler {
	h, exist := s.routers[action]
	if !exist {
		return nil
	}

	return h
}
