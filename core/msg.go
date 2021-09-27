package core

const HeaderLength uint32 = 12

type FireMsg struct {
	ID            uint32
	PayloadLength uint32
	ActionId      uint32
	Payload       []byte
}
