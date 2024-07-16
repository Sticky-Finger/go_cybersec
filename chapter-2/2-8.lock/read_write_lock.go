package main

import (
	"fmt"
	"sync"
	"time"
)

var money = 100
var rw sync.RWMutex

// 余额
func balance() int {
	return money
}

// 存钱
func depositMoney(m int) {
	rw.Lock()
	money += m
	rw.Unlock()
}

func main() {
	go depositMoney(100)
	go depositMoney(100)
	go depositMoney(100)

	for {
		rw.RLock()
		time.Sleep(time.Second)
		fmt.Println(balance())
		rw.RUnlock()
	}
}
