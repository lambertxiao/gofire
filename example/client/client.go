package main

import (
	"fmt"
	"gofire"
	"log"
	"time"
)

func main() {
	client := gofire.NewClient("127.0.0.1:7777")
	client.SetMsgQueueSize(100)
	err := client.Connect()
	if err != nil {
		panic(err)
	}

	var i uint32 = 0

	go func() {
		for ; i < 10; i++ {
			msg := gofire.NewMsg(i, 0, []byte(fmt.Sprintf("msg id: %d", i)))
			err := client.Send(msg)
			if err != nil {
				panic(err)
			}

			time.Sleep(500 * time.Millisecond)
		}
	}()

	log.Println("wait resp msg")
	for msg := range client.OnMsg() {
		log.Println("receive resp for msg id: ", msg.GetID())
	}
}
