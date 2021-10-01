package main

import (
	gofire "gofire/core"
	"gofire/example/proto"
	"gofire/generator"
)

var endpoint gofire.Endpoint
var gen gofire.IChannelGenerator
var pcodec gofire.IPacketCodec
var mcodec gofire.IMsgCodec

func init() {
	endpoint = gofire.Endpoint{Ip: "127.0.0.1", Port: 7777}
	sgen, err := generator.NewTCPServerConnGenerator(endpoint)
	// sgen, err := generator.NewUDPServerConnGenerator(endpoint)
	if err != nil {
		panic(err)
	}
	gen = sgen
	pcodec = gofire.NewPacketCodec(gofire.TransProtocol{Name: 1, Version: 1})
	mcodec = proto.NewCustomMsgCodec()
}

func main() {
	server := gofire.NewServer(
		gen, pcodec, mcodec,
	)

	handler := &FooHandler{}
	server.RegistAction("hello", handler)
	err := server.Listen()

	if err != nil {
		panic(err)
	}
}

type FooHandler struct{}

func (h *FooHandler) Do(req gofire.Request) {
	msg := &proto.Message{
		MsgId:  req.Msg.GetAction(),
		Action: "hello-resp",
		Body: map[string]interface{}{
			"name": "bar",
		},
	}

	req.Stream.Write(msg)
}
