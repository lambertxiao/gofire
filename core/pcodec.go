package core

import (
	"encoding/binary"
	"io"
)

type PacketCodec struct{}

func NewPacketCodec() IPacketCodec {
	c := &PacketCodec{}
	return c
}

func (c *PacketCodec) Encode(payload []byte, w io.Writer) error {
	payloadSize := uint32(len(payload))
	if err := binary.Write(w, binary.LittleEndian, payloadSize); err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, payload); err != nil {
		return err
	}

	return nil
}

func (c *PacketCodec) Decode(r io.Reader) ([]byte, error) {
	var size uint32
	if err := binary.Read(r, binary.LittleEndian, &size); err != nil {
		return nil, err
	}

	payload := make([]byte, size)
	if _, err := io.ReadFull(r, payload); err != nil {
		return nil, err
	}

	return payload, nil
}
