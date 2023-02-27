package main

import (
	"fmt"
	"sync"
	"time"
)

// Bank 模擬「銀行」的結構
type Bank struct {
	balance int
	mutex   sync.Mutex
}

// Deposit 存款
func (b *Bank) Deposit(amount int) {
	// 透過「互斥鎖」，做到 goroutine 同時間只能一個存取該變數。
	// 避免同時大量操作該變數時，發生錯誤。
	b.mutex.Lock()
	time.Sleep(time.Second)
	b.balance += amount
	b.mutex.Unlock()
}

// GetBalance 取得銀行餘額
func (b *Bank) GetBalance() (balance int) {
	b.mutex.Lock()
	time.Sleep(time.Second)
	balance = b.balance
	b.mutex.Unlock()

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
