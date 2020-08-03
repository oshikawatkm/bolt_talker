package network

import (
	"log"
	"net"
)

type Client struct {
	Conn net.Conn
}

func NewClient(address string) *Client {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{Conn: conn}
}

func (c *Client) SendMessage(msg protocol.Message) (int, error) {
	message := common.NewMessage(msg.Command(), msg.Encode())
	log.Printf("send		: %s", string(message.Command[:]))
	return c.Conn.Wirte(message.Encode())
}
