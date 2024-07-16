package main

// # 同步锁/互斥锁

import (
	"fmt"
	"sync"
	"time"
)

var money = 100
var sw sync.WaitGroup

// withdrawMoney 取钱
func withdrawMoney(m int) {
	if money > m {
		// 模拟执行业务代码产生延迟
		time.Sleep(time.Second * 2)
		money = money - m
		fmt.Printf("取款 %d 成功，余额 %d\n", m, money)
	} else {
		fmt.Printf("取款 %d 失败，余额 %d\n", m, money)
	}
	sw.Done()
}

func main() {
	sw.Add(2)
	go withdrawMoney(99)
	go withdrawMoney(99)
	sw.Wait()
}
