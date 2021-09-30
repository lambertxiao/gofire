package core

import (
	"context"
	"log"
)

type ClientStream struct {
	client     *FireClient
	conn       IConn
	ssm        *MsgSSM
	ctx        context.Context
	cancel     context.CancelFunc
	msgChannel chan IMsg
}

func NewClientStream(conn IConn, client *FireClient, ssm *MsgSSM) IStream {
	ctx, cancel := context.WithCancel(context.Background())
	s := &ClientStream{
		client:     client,
		conn:       conn,
		ctx:        ctx,
		cancel:     cancel,
		ssm:        ssm,
		msgChannel: make(chan IMsg, 2),
	}
	return s
}

func (s *ClientStream) Flow() {
	go s.ReadLoop()
	go s.WriteLoop()
}

func (s *ClientStream) Write(msg IMsg) {
	s.msgChannel <- msg
}

func (s *ClientStream) ReadLoop() {
	defer s.Close()
	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			data, err := s.client.pcodec.Decode(s.conn)
			if err != nil {
				log.Println("pcodec decode error", err)
				return
			}

			msg, err := s.client.mcodec.Decode(data)
			if err != nil {
				log.Println("mcodec decode error", err)
				return
			}

			s.ssm.Resp = msg
			s.ssm.Done()
		}
	}
}

func (s *ClientStream) WriteLoop() {
	for {
		select {
		case <-s.ctx.Done():
			return
		case msg := <-s.msgChannel:
			msgData, err := s.client.mcodec.Encode(msg)
			if err != nil {
				log.Println("client mcodec encode msg error", err)
				continue
			}

			err = s.client.pcodec.Encode(msgData, s.conn)
			if err != nil {
				log.Println("client pcodec encode msg error", err)
				continue
			}
		}
	}
}

func (s *ClientStream) Close() {
}
