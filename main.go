package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	bank        = 100 // available money in the bank
	stingyAdded = 0   // money added by stingy
	spendySpend = 0   // money spend by spendy
	access      = sync.Mutex{}
)

// stingy will add value to Bank
func stingy() {
	for i := 1; i <= 100; i++ {
		access.Lock()
		bank += 10
		stingyAdded += 10
		access.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Stingy Added :", stingyAdded) //money added by stingy
}

// spendy will spend from the Bank
func spendy() {
	for i := 1; i <= 100; i++ {
		access.Lock()
		bank -= 10
		spendySpend += 10
		access.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Spendy spend : ", spendySpend) // money spended by spendy
}

func main() {
	start := time.Now()

	fmt.Println("Starting Balance : ", bank) //starting balance

	go stingy()
	go spendy()

	time.Sleep(3000 * time.Millisecond)

	fmt.Println("Closing Balance : ", bank) //closing balance

	//calculating time taken for execution
	elapsed := time.Since(start)
	fmt.Println("Execution Time : ", elapsed)
}
