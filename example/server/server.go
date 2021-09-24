package main

import (
	"gofire"
	"gofire/iface"
)

func main() {
	server := gofire.NewServer("127.0.0.1", 7777)
	handler := &FooHandler{}
	server.RegistAction(0, handler)
	err := server.Serve("")

	if err != nil {
		panic(err)
	}
}

type FooHandler struct{}

func (h *FooHandler) Do(req iface.Request) {
	payload := string(req.Msg.GetPayload())
	msg := gofire.NewMsg(0, 0, []byte("resp your request: "+payload))
	req.Conn.WriteMsg(msg)
}
