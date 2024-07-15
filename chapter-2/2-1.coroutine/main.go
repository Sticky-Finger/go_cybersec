package main

// import (
// 	"fmt"
// 	"time"
// )

// func worker(count int) {
// 	for i := 1; i < count; i++ {
// 		time.Sleep(500)
// 		fmt.Print(" ", i)
// 	}
// 	fmt.Println()
// }

// func main() {
// 	// not use coroutine
// 	now1 := time.Now()
// 	worker(10)
// 	worker(10)
// 	worker(10)
// 	fmt.Println(time.Now().Sub(now1))

// 	// start 3 goroutines at the same time
// 	go worker(10)
// 	go worker(10)
// 	go worker(10)
// 	time.Sleep(time.Second * 5)
// }

// # channel

// func main() {
// 	channel := make(chan string)

// 	go func() {
// 		time.Sleep(time.Second * 3)
// 		channel <- "Hello from goroutine!"
// 	}()

// 	fmt.Println("Waiting for channel...")
// 	message := <-channel
// 	fmt.Println(message)
// }

// # select

// import (
// 	"fmt"
// )

// func main() {
// 	channel1 := make(chan string)
// 	channel2 := make(chan string)

// 	go func() {
// 		channel1 <- "Hello from channel 1!"
// 	}()

// 	go func() {
// 		channel2 <- "Hello from channel 2!"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case message1 := <-channel1:
// 			fmt.Println(message1)
// 		case message2 := <-channel2:
// 			fmt.Println(message2)
// 		}
// 	}
// }

// # WaitGroup

import (
	"fmt"
	"sync"
	"time"
)

var sw sync.WaitGroup

func worker(count int) {
	for i := 1; i < count; i++ {
		time.Sleep(500)
		fmt.Print(" ", i)
	}
	fmt.Println()
	sw.Done()
}

func main() {
	now2 := time.Now()

	go worker(10)
	sw.Add(1)
	go worker(10)
	sw.Add(1)
	go worker(10)
	sw.Add(1)
	sw.Wait()
	now3 := time.Now()
	fmt.Println(now3.Sub(now2))
}
