package main

import (
	gofire "gofire/core"
	"gofire/example/proto"
	"gofire/generator"
	"strconv"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
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
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		time.Sleep(1 * time.Second)

		go func(i int) {
			id, _ := uuid.NewV4()
			helloMsg := &proto.Message{
				MsgId:   id.String(),
				Payload: []byte("hello: " + strconv.Itoa(i)),
			}

			resp, err := client.SyncSend(helloMsg)
			if err != nil {
				logrus.Info(err)
				return
			}

			logrus.Infof("msg_id: %s, %s", resp.GetID(), string(resp.GetPayload()))
			wg.Done()
		}(i)
	}

	wg.Wait()
}
