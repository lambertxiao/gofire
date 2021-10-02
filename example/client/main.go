package main

import (
	gofire "gofire/core"
	"gofire/example/proto"
	"gofire/generator"
	"log"
	"sync"

	"github.com/gofrs/uuid"
)

var endpoint gofire.Endpoint
var gen gofire.IChannelGenerator
var pcodec gofire.IPacketCodec
var mcodec gofire.IMsgCodec

func init() {
	endpoint = gofire.Endpoint{Ip: "127.0.0.1", Port: 7777}
	_gen, err := generator.NewTCPClientConnGenerator(endpoint)
	// _gen, err := generator.NewUDPClientConnGenerator(endpoint)
	if err != nil {
		panic(err)
	}

	gen = _gen
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

	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			id, _ := uuid.NewV4()
			helloMsg := &proto.Message{
				MsgId:  id.String(),
				Action: "hello",
				Body: map[string]interface{}{
					"name": "foo",
				},
			}

			resp, err := client.Send(helloMsg)
			if err != nil {
				log.Println(err)
				return
			}

			log.Println(resp)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
