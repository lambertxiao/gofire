package core

import (
	"context"
	"log"
)

type ServerTransport struct {
	ch         IChannel
	server     *FireServer
	ctx        context.Context
	cancel     context.CancelFunc
	msgChannel chan IMsg
}

func NewServerTransport(ch IChannel, server *FireServer) ISTransport {
	ctx, cancel := context.WithCancel(context.Background())
	s := &ServerTransport{
		ch:         ch,
		server:     server,
		ctx:        ctx,
		cancel:     cancel,
		msgChannel: make(chan IMsg),
	}
	return s
}

func (t *ServerTransport) Flow() {
	go t.ReadLoop()
	go t.WriteLoop()
}

func (t *ServerTransport) Close() {
	t.ch.Close()
	close(t.msgChannel)
}

func (t *ServerTransport) ReadLoop() {
	// defer t.Close()
	for {
		select {
		case <-t.ctx.Done():
			return
		default:
			data, err := t.server.pcodec.Decode(t.ch)
			if err != nil {
				log.Println("pcodec decode error", err)
				return
			}

			msg, err := t.server.mcodec.Decode(data)
			if err != nil {
				log.Println("mcodec decode error", err)
				return
			}

			req := Request{
				Stream: t,
				Msg:    msg,
			}

			log.Printf("server receive msg: %+v\n", msg)
			handler := t.server.GetHandler(msg.GetAction())
			if handler == nil {
				log.Println("not support action for action:", msg.GetAction())
				return
			}

			go handler.Do(req)
		}
	}
}

func (t *ServerTransport) WriteLoop() {
	for {
		select {
		case <-t.ctx.Done():
			return
		case msg := <-t.msgChannel:
			log.Println("write msg ....", msg)
			msgData, err := t.server.mcodec.Encode(msg)
			if err != nil {
				log.Println("mcodec encode msg error", err)
				continue
			}

			err = t.server.pcodec.Encode(msgData, t.ch)
			if err != nil {
				log.Println("pcodec encode msg error", err)
				continue
			}

			log.Println("write done")
		}
	}
}

func (t *ServerTransport) Write(msg IMsg) {
	log.Println("server write msg to channel")
	t.msgChannel <- msg
}
