package main

import (
	gofire "gofire/core"
	"gofire/example/proto"
	"gofire/generator"
	"log"
	"sync"

	"github.com/gofrs/uuid"
)

func main() {
	endpoint := gofire.Endpoint{Ip: "127.0.0.1", Port: 7777}
	client, err := gofire.NewClient(
		generator.NewTCPClientConnGenerator(endpoint),
		gofire.NewPacketCodec(gofire.TransProtocol{Name: 1, Version: 1}),
		proto.NewCustomMsgCodec(),
		gofire.NewDefaultMsgQueue(100),
		10,
	)
	if err != nil {
		panic(err)
	}

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

			resp, err := client.SyncSend(helloMsg)
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
