package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloSecself struct {}

func (p *HelloSecself) Hello(name string, result *string) error {
	*result = "hello " + name
	return nil
}

func main() {
	rpc.RegisterName("HelloSecself", new(HelloSecself))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}
