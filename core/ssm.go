package core

import (
	"gofire/iface"
	"sync"
)

type MsgSSM struct {
	sync.WaitGroup
	Resp iface.IMsg
}

func NewMsgSSM() *MsgSSM {
	m := &MsgSSM{}
	return m
}

func (m *MsgSSM) Return() iface.IMsg {
	m.WaitGroup.Wait()
	return m.Resp
}
