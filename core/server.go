package core

import (
	"fmt"
	"log"
)

type FireServer struct {
	generator IChannelGenerator
	mcodec    IMsgCodec
	pcodec    IPacketCodec
	routers   map[string]IHandler
}

type Endpoint struct {
	Ip   string
	Port int
}

func (e Endpoint) String() string {
	return fmt.Sprintf("%s:%d", e.Ip, e.Port)
}

func NewServer(
	generator IChannelGenerator,
	pcodec IPacketCodec,
	mcodec IMsgCodec,
) IServer {
	s := &FireServer{
		generator: generator,
		mcodec:    mcodec,
		pcodec:    pcodec,
		routers:   make(map[string]IHandler),
	}

	return s
}

func (s *FireServer) Listen() error {
	log.Println("server listening...")
	for {
		ch, err := s.generator.Gen()
		log.Println("get channel from generator")
		if err != nil {
			log.Println(err)
			break
		}

		stream := NewServerTransport(ch, s)
		go stream.Flow()
	}

	return nil
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
