package main

import (
	"log"
	"os"
)

// # simple applying

// func main() {
// 	log.SetPrefix("Login: ")
// 	log.Printf("%s login, age:%d", "se", 2)
// }

// # self-defined log context

// func main() {
// 	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
// 	log.Printf("hello %s", "secself")
// }

// # output log to file

func main() {
	file, err := os.OpenFile("custom.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.Lshortfile|log.Ldate|log.Lmicroseconds)
	logger.Println("hello secself")
}
