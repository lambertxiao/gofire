package main

import (
	"gofire"
)

func main() {
	client := gofire.NewClient("127.0.0.1:7777")
	err := client.Connect()
	if err != nil {
		panic(err)
	}

	msg := gofire.NewMsg(1, 1, []byte("hello, i am client"))
	client.Send(msg)
}
