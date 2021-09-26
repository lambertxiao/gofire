package gofire

import (
	"context"
	"gofire/iface"
	"log"
)

type FireStream struct {
	conn       iface.IConn
	server     *FireServer
	ctx        context.Context
	cancel     context.CancelFunc
	msgChannel chan iface.IMsg
}

func NewFireStream(conn iface.IConn, server *FireServer) iface.IStream {
	ctx, cancel := context.WithCancel(context.Background())
	s := &FireStream{
		conn:       conn,
		server:     server,
		ctx:        ctx,
		cancel:     cancel,
		msgChannel: make(chan iface.IMsg),
	}
	return s
}

func (s *FireStream) Run() {
	go s.ReadLoop()
	go s.WriteLoop()
}

func (s *FireStream) Close() {
	s.conn.Close()
	close(s.msgChannel)
}

func (s *FireStream) ReadLoop() {
	defer s.Close()
	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			data, err := s.server.pCodec.Decode(s.conn)
			if err != nil {
				log.Println("pCodec decode error", err)
				return
			}

			msg := s.server.mCodec.Decode(data)
			req := iface.Request{
				Stream: s,
				Msg:    msg,
			}

			// headData := make([]byte, HeaderLength)

			// if _, err := io.ReadFull(c.conn, headData); err != nil {
			// 	log.Println("read head data error", err)
			// 	return
			// }

			// msg := &FireMsg{}
			// if err := msg.UnpackHead(headData); err != nil {
			// 	log.Println("read msg head error", err)
			// 	return
			// }

			// payloadData := make([]byte, msg.GetPayloadLen())
			// if _, err := io.ReadFull(c.conn, payloadData); err != nil {
			// 	log.Println("read msg payload error", err)
			// 	return
			// }

			// msg.SetPayload(payloadData)
			handler := s.server.GetActionHandler(msg.GetAction())
			if handler == nil {
				log.Println("not support action")
				s.Write(NewMsg(0, 0, []byte("not support action")))
				return
			}

			go handler.Do(req)
		}
	}
}

func (s *FireStream) WriteLoop() {
	for {
		select {
		case <-s.ctx.Done():
			return
		case msg := <-s.msgChannel:
			msgData := s.server.mCodec.Encode(msg)
			err := s.server.pCodec.Encode(msgData, s.conn)
			if err != nil {
				log.Println("write msg data to connection error", err)
				return
			}
		}
	}
}

func (s *FireStream) Write(msg iface.IMsg) {
	s.msgChannel <- msg
}
