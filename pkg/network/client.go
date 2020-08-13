package network

import (
	"log"
	"net"

	"github.com/oshikawatkm/bolt_talker/message"
)

type Client struct {
	Conn net.Conn
}

func NewClient(address string) *Client {
	log.Printf("Connecting: " + address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{Conn: conn}
}

func (c *Client) SendMessage(msg message.Message) (int, error) {
	message := common.NewMessage(msg.Command(), msg.Encode())
	log.Printf("send		: %s", string(message.Command[:]))
	return c.Conn.Wirte(message.Encode())
}
