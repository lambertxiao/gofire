package core

import (
	"sync"
)

type MsgSSM struct {
	sync.WaitGroup
	Resp IMsg
	Err  error
}

func NewMsgSSM() *MsgSSM {
	m := &MsgSSM{}
	return m
}

func (m *MsgSSM) Go() {
	m.Add(1)
}

func (m *MsgSSM) Done() {
	m.WaitGroup.Done()
}

func (m *MsgSSM) Return() (IMsg, error) {
	m.WaitGroup.Wait()
	return m.Resp, nil
}
