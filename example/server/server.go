package main

import (
	"gofire"
	"gofire/iface"
)

func main() {
	// 初始化server
	// 1. 告诉server用的是什么传输协议
	// 2. 告诉server传输的消息格式

	endpoint := gofire.Endpoint{
		Ip:   "127.0.0.1",
		Port: 7777,
	}
	generator, err := gofire.NewTCPServerConnGenerator(endpoint)
	if err != nil {
		panic(err)
	}

	server := gofire.NewServer(generator)
	handler := &FooHandler{}
	server.AddRouter(0, handler)
	err := server.Listen("")

	if err != nil {
		panic(err)
	}
}

type FooHandler struct{}

func (h *FooHandler) Do(req iface.Request) {
	payload := string(req.Msg.GetPayload())
	msg := gofire.NewMsg(req.Msg.GetID(), 0, []byte("resp your request: "+payload))
	req.Stream.Write(msg)
}
