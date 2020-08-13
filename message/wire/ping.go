package wire

type PingPayload []byte

type Ping struct {
	NumPongBytes uint16
	PaddingBytes PingPayload
}

func NewPing(numBytes uint16) *Ping {
	return &Ping{
		NumPongBytes: numBytes
	}
}

func (p *Ping) Decode(r io.Reader, prev uint32) error {
	return ReadElements(r,
		&p.NumPongBytes,
		&p.PaddingBytes
	)
}

func (p *Ping) Encode(w io.Witer, prev uint32) error {
	return WiteElemets(w,
		p.NumPongBytes,
		p.PaddingBytes
	)
}
