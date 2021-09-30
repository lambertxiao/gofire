package core

import (
	"encoding/binary"
	"io"
)

type TransProtocol struct {
	Name    uint32
	Version uint32
}

type TransHeader struct {
	Name          uint32
	Version       uint32
	ContentLength uint32
}

type PacketCodec struct {
	tp TransProtocol
}

func NewPacketCodec(tp TransProtocol) IPacketCodec {
	c := &PacketCodec{
		tp: tp,
	}
	return c
}

func (c *PacketCodec) Encode(payload []byte, w io.Writer) error {
	var err error
	if err = binary.Write(w, binary.LittleEndian, c.tp.Name); err != nil {
		return err
	}

	if err = binary.Write(w, binary.LittleEndian, c.tp.Version); err != nil {
		return err
	}

	if err = binary.Write(w, binary.LittleEndian, uint32(len(payload))); err != nil {
		return err
	}

	// write payload
	if err = binary.Write(w, binary.LittleEndian, payload); err != nil {
		return err
	}

	return nil
}

func (c *PacketCodec) Decode(r io.Reader) ([]byte, error) {
	th := &TransHeader{}
	if err := binary.Read(r, binary.LittleEndian, &th.Name); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &th.Version); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &th.ContentLength); err != nil {
		return nil, err
	}

	payload := make([]byte, th.ContentLength)
	if _, err := io.ReadFull(r, payload); err != nil {
		return nil, err
	}

	return payload, nil
}
