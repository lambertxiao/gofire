package main

import (
	gofire "gofire/core"
	"gofire/example/proto"
	"gofire/generator"
	"gofire/iface"
)

func main() {
	endpoint := gofire.Endpoint{Ip: "127.0.0.1", Port: 7777}
	sgenerator, err := generator.NewTCPServerConnGenerator(endpoint)
	if err != nil {
		panic(err)
	}

	pcodec := gofire.NewPacketCodec()
	mcodec := proto.NewMsgCodec()
	server := gofire.NewServer(
		sgenerator, pcodec, mcodec,
	)

	handler := &FooHandler{}
	server.RegistAction("hello", handler)
	err = server.Listen()

	if err != nil {
		panic(err)
	}
}

type FooHandler struct{}

func (h *FooHandler) Do(req iface.Request) {
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
