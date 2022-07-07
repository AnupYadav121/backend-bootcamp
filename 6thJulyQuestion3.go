package main

import (
	"errors"
	"fmt"
	"sync"
)

// name different
// make account structure also
type Bank struct {
	mu       sync.Mutex
	balances map[string]int
}

func (c *Bank) add(name string, rupee int) {
	c.mu.Lock()
	c.balances[name] += rupee
	c.mu.Unlock()
}

func (c *Bank) subtract(name string, rupee int) {
	c.mu.Lock()
	c.balances[name] -= rupee
	c.mu.Unlock()
}

func main() {
	c := Bank{
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
		} else {
			c.subtract(name, n)
		}
		wg.Done()
	}

	fmt.Println(c.balances)

	wg.Add(4)
	go doDeposit("anup", 10000)
	go doDeposit("kashi", 10000)

	go doWithdraw("anup", 1004)
	go doWithdraw("kashi", 104)

	wg.Wait()
	fmt.Println(c.balances)
}
