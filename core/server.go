package core

import (
	"fmt"
	"log"
)

type FireServer struct {
	generator ConnGenerator
	mcodec    MsgCodec
	pcodec    IPacketCodec
	routers   map[string]Handler
}

type Endpoint struct {
	Ip   string
	Port int
}

func (e Endpoint) String() string {
	return fmt.Sprintf("%s:%d", e.Ip, e.Port)
}

func NewServer(
	generator ConnGenerator,
	pcodec IPacketCodec,
	mcodec MsgCodec,
) IServer {
	s := &FireServer{
		generator: generator,
		mcodec:    mcodec,
		pcodec:    pcodec,
		routers:   make(map[string]Handler),
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

func (s *FireServer) RegistAction(action string, handler Handler) {
	log.Println("regist action id", action)
	s.routers[action] = handler
}

func (s *FireServer) GetHandler(action string) Handler {
	h, exist := s.routers[action]
	if !exist {
		return nil
	}

	return h
}
