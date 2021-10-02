package core

import (
	"sync"
)

type MsgSSM struct {
	sync.WaitGroup
	Resp IMsg
	Err  error
	ch   chan IMsg
}

func NewMsgSSM() *MsgSSM {
	m := &MsgSSM{
		ch: make(chan IMsg),
	}
	return m
}

func (m *MsgSSM) Go() {
	m.Add(1)
}

func (m *MsgSSM) Done() {
	m.WaitGroup.Done()
}

func (m *MsgSSM) WaitDone() {

}

func (m *MsgSSM) Return() (IMsg, error) {
	m.WaitGroup.Wait()
	return m.Resp, nil
}
