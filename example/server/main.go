package main

import (
	gofire "gofire/core"
	"gofire/example/proto"
	"gofire/generator"

	"github.com/sirupsen/logrus"
)

func main() {
	cg, err := generator.NewTCPServerConnGenerator(
		gofire.Endpoint{Ip: "127.0.0.1", Port: 7777},
	)
	if err != nil {
		panic(err)
	}
	server := gofire.NewServer(
		cg,
		gofire.NewPacketCodec(gofire.TransProtocol{Name: 1, Version: 1}),
		proto.NewCustomMsgCodec(),
		onRecvMsg,
	)

	err = server.Listen()

	if err != nil {
		panic(err)
	}
}

func onRecvMsg(tp gofire.Transport, msg gofire.Msg) {
	logrus.Infof("recv msg_id: %s msg_len: %d", msg.GetID(), len(msg.GetPayload()))

	replyMsg := &proto.Message{
		MsgId:   msg.GetID(),
		Payload: []byte("reply: " + string(msg.GetPayload())),
	}
	tp.SendMsg(replyMsg)
}
