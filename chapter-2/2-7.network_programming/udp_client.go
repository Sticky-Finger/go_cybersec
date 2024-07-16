package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	host := "localhost"
	port := "8053"
	server, err := net.ResolveUDPAddr("udp", host+":"+port)

	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, server)
	if err != nil {
		log.Fatal(err)
	}

	// close the connection
	defer conn.Close()

	data := "我是客户端"
	_, err = conn.Write([]byte(data))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("客户端[%s] >>> 服务端[%s] : %s\n", conn.LocalAddr(), conn.RemoteAddr(), data)

	received := make([]byte, 1024)
	_, err = conn.Read(received)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("客户端[%s] <<< 服务器[%s] : %s\n", conn.LocalAddr(), conn.RemoteAddr(), received)
}
