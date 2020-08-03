package main

import (
	"log"
	"os"

	""
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	// c := network.NewClient("10.4.255.146:10011")
	// defer c.Conn.Close()
	// log.Printf("[*] remote addre: %s", c.Conn.RemoteAddr().String())

	// COMMAND SWITCH
	command := os.Args[1]
	switch command {
	case "ping":

	default:
		log.Printf("[-] NO COMMAND")
	}
	log.Printf("[*] FINISH")
}
