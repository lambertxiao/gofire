package main

import (
	"gofire/core"
	"gofire/example/proto"
	"gofire/generator"
)

func main() {
	endpoint := core.Endpoint{Ip: "127.0.0.1", Port: 7777}
	generator, err := generator.NewTCPServerConnGenerator(endpoint)
	// generator, err := generator.NewUDPServerConnGenerator(endpoint)
	if err != nil {
		panic(err)
	}

	pcodec := core.NewPacketCodec()
	mcodec := proto.NewCustomMsgCodec()
	server := core.NewServer(
		generator, pcodec, mcodec,
	)

	handler := &FooHandler{}
	server.RegistAction("hello", handler)
	err = server.Listen()

	if err != nil {
		panic(err)
	}
}

type FooHandler struct{}

func (h *FooHandler) Do(req core.Request) {
	msg := &proto.Message{
		Head: proto.MessageHead{
			MsgId:  "0000-0000-0000-0001",
			Action: "hello-resp",
		},
		Body: map[string]interface{}{
			"name": "bar",
		},
	}

	req.Stream.Write(msg)
}
