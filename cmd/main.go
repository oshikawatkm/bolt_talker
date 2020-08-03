package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/oshikawatkm/bolt_talker/pkg/network"
)

func main() {
	// if len(os.Args) < 2 {
	// 	os.Exit(1)
	// }

	var online bool
	var ip string
	var port string
	var ip_port string

	online = false

	log.Printf("[*] START BOLT TALKER")
	online = true

	// Start Message Read handler
	go message_handler()

	// !!! Start CLI Operation !!!
	// c := network.NewClient("10.4.255.146:10011")
	stdin := bufio.NewScanner(os.Stdin)

	fmt.Print("[+] INPUT LN Node IP ADDRESS : ")
	if stdin.Scan() {
		ip = stdin.Text()
	}

	fmt.Print("[+] INPUT PORT NUMBER: ")
	if stdin.Scan() {
		port = stdin.Text()
	}
	ip_port = ip + ":" + port

	c := network.NewClient(ip_port)
	defer c.Conn.Close()

	log.Printf("[*] remote addre: %s", c.Conn.RemoteAddr().String())

	// COMMAND SWITCH
	command := os.Args[1]
	switch command {
	case "ping":

	default:
		log.Printf("[-] NO COMMAND")
	}
	log.Printf("[*] FINISH")
	online = false
}

func message_handler() {

	for onlien == true {
		msg := 
		switch msg {
		case *wire.Pong:
			log.Printf("[!] PONG MESSAGE")

		case *wire.Ping:
			log.Printf("[!] PING MESSAGE")
		}
	}
}
