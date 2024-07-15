package main

import (
	"log"
	"net/rpc"
	"fmt"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	var result string
	err = client.Call("HelloSecself.Hello", "secself", &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
