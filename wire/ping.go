package lnwire

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

func Decode(r io.Reader, prev uint32) error {
	return ReadElements(r,
		&p.NumPongBytes,
		&p.PaddingBytes
	)
}