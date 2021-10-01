package core

type Transport struct {
	// 传输通道
	ch IChannel
	// 传输包编解码器
	pcodec IPacketCodec
	mcodec IMsgCodec
}

func NewTransport(ch IChannel, pcodec IPacketCodec, mcodec IMsgCodec) ITransport {
	t := &Transport{
		ch:     ch,
		pcodec: pcodec,
		mcodec: mcodec,
	}
	return t
}

func (t *Transport) RoundTrip(msg IMsg) (IMsg, error) {
	var err error
	payload, err := t.mcodec.Encode(msg)
	if err != nil {
		return nil, err
	}

	if err = t.pcodec.Encode(payload, t.ch); err != nil {
		return nil, err
	}

	repayload, err := t.pcodec.Decode(t.ch)
	if err != nil {
		return nil, err
	}

	msg, err = t.mcodec.Decode(repayload)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
