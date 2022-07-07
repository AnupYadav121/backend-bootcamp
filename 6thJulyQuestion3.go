package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Account struct {
	mu       sync.Mutex
	balances map[string]int
}

func (c *Account) add(name string, rupee int) {
	c.mu.Lock()
	c.balances[name] += rupee
	c.mu.Unlock()
}

func (c *Account) subtract(name string, rupee int) {
	c.mu.Lock()
	c.balances[name] -= rupee
	c.mu.Unlock()
}

func main() {
	c := Account{
		balances: map[string]int{"anup": 0, "kashi": 0},
	}

	var wg sync.WaitGroup

	doDeposit := func(name string, n int) {

		c.add(name, n)
		wg.Done()
	}

	doWithdraw := func(name string, n int) {
		if c.balances[name] < n {
			fmt.Println(nil, errors.New("Sorry , can not withdraw more than current balance amount"))
			wg.Done()
		} else {
			c.subtract(name, n)
			wg.Done()
		}
	}

	fmt.Println(c.balances)

	wg.Add(4)
	go doDeposit("anup", 10000)
	go doDeposit("kashi", 10000)
	time.Sleep(time.Second)
	go doWithdraw("anup", 10001)
	go doWithdraw("kashi", 10001)

	wg.Wait()
	fmt.Println(c.balances)
}
