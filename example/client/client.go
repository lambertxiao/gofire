package main

import (
	"fmt"
	"gofire"
	"log"
	"time"
)

func main() {
	client := gofire.NewClient("127.0.0.1:7777")
	err := client.Connect()
	if err != nil {
		panic(err)
	}

	var i uint32 = 0

	for {
		msg := gofire.NewMsg(i, 0, []byte(fmt.Sprintf("msg id: %d", i)))
		resp, err := client.SyncSend(msg)
		if err != nil {
			panic(err)
		}

		log.Println(string(resp))
		i++
		time.Sleep(500 * time.Millisecond)
	}
}
