package core

import (
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
}

type DefaultServer struct {
	cg     ConnGenerator
	mcodec MsgCodec
	pcodec PacketCodec
	cb     MsgRecvCallback
}

func NewServer(
	cg ConnGenerator,
	pcodec PacketCodec,
	mcodec MsgCodec,
	cb MsgRecvCallback,
) Server {
	s := &DefaultServer{
		cg:     cg,
		mcodec: mcodec,
		pcodec: pcodec,
		cb:     cb,
	}

	return s
}

func (s *DefaultServer) Listen() error {
	logrus.Info("server listening...")
	for {
		conn, err := s.cg.Gen()
		logrus.Info("get conn from generator")
		if err != nil {
			logrus.Info(err)
			break
		}

		tp := NewTransport(conn, s.pcodec, s.mcodec)
		tp.SetMsgCB(s.cb)
		go tp.Open()
	}

	return nil
}
