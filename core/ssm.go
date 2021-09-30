package core

import (
	"sync"
)

type MsgSSM struct {
	sync.WaitGroup
	Resp IMsg
}

func NewMsgSSM() *MsgSSM {
	m := &MsgSSM{}
	return m
}

func (m *MsgSSM) Return() IMsg {
	m.WaitGroup.Wait()
	return m.Resp
}
