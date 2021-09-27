package main

import (
	gofire "gofire/core"
	"gofire/example/proto"
	"gofire/generator"
	"log"
)

func main() {
	endpoint := gofire.Endpoint{Ip: "127.0.0.1", Port: 7777}
	// generator := generator.NewTCPClientConnGenerator(endpoint)
	generator, err := generator.NewUDPClientConnGenerator(endpoint)
	if err != nil {
		panic(err)
	}

	pcodec := gofire.NewPacketCodec()
	mcodec := proto.NewMsgCodec()
	client := gofire.NewClient(
		endpoint.String(),
		generator,
		pcodec,
		mcodec,
	)
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
