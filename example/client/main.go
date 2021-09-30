package main

import (
	gofire "gofire/core"
	"gofire/example/proto"
	"gofire/generator"
	"log"
)

var endpoint gofire.Endpoint
var gen gofire.IChannelGenerator
var pcodec gofire.IPacketCodec
var mcodec gofire.IMsgCodec

func init() {
	endpoint = gofire.Endpoint{Ip: "127.0.0.1", Port: 7777}
	gen = generator.NewTCPClientConnGenerator(endpoint)
	pcodec = gofire.NewPacketCodec(gofire.TransProtocol{Name: 1, Version: 1})
	mcodec = proto.NewCustomMsgCodec()
}

func main() {
	ch, err := gen.Gen()
	if err != nil {
		panic(err)
	}

	transport := gofire.NewTransport(ch, pcodec, mcodec)
	client := gofire.NewClient(transport)
	helloMsg := &proto.Message{
		Head: proto.MessageHead{
			MsgId:  "0000-0000-0000-0001",
			Action: "hello",
		},
		Body: map[string]interface{}{
			"name": "foo",
		},
	}

	resp, err := client.Send(helloMsg)
	if err != nil {
		panic(err)
	}

	log.Println(resp)
}
