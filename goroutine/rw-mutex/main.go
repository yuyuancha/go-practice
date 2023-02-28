package main

import (
	"fmt"
	"sync"
	"time"
)

// Bank 模擬「銀行」的結構
type Bank struct {
	balance int
	// 透過 RWMutex 讀寫鎖，允許多個 goroutine 對讀取同時使用。
	rwMutex sync.RWMutex
}

// Deposit 存款
func (b *Bank) Deposit(amount int) {
	b.rwMutex.Lock()
	time.Sleep(time.Second)
	b.balance += amount
	b.rwMutex.Unlock()
}

// GetBalance 取得銀行餘額
func (b *Bank) GetBalance() (balance int) {
	// 在讀取的時候，允許多個 goroutine 同時操作。
	b.rwMutex.RLock()
	time.Sleep(time.Second)
	balance = b.balance
	b.rwMutex.RUnlock()

	return
}

func main() {
	var waitGroup sync.WaitGroup
	var bank = &Bank{}
	var num = 5

	waitGroup.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			bank.Deposit(100)
			fmt.Printf("向銀行存入(%d): %d\n", i, 100)

			waitGroup.Done()
		}(i)
	}

	waitGroup.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			fmt.Printf("查詢餘額(%d): %d\n", i, bank.GetBalance())

			waitGroup.Done()
		}(i)
	}

	waitGroup.Wait()

	fmt.Println("程式結束。")
}
