package main

import (
	"errors"
	"fmt"
	"sync"
<<<<<<< HEAD
)

// name different
// make account structure also
type Bank struct {
=======
	"time"
)

type Account struct {
>>>>>>> d1e27de128dbb838b71a55a60828b34a847da776
	mu       sync.Mutex
	balances map[string]int
}

<<<<<<< HEAD
func (c *Bank) add(name string, rupee int) {
=======
func (c *Account) add(name string, rupee int) {
>>>>>>> d1e27de128dbb838b71a55a60828b34a847da776
	c.mu.Lock()
	c.balances[name] += rupee
	c.mu.Unlock()
}

<<<<<<< HEAD
func (c *Bank) subtract(name string, rupee int) {
=======
func (c *Account) subtract(name string, rupee int) {
>>>>>>> d1e27de128dbb838b71a55a60828b34a847da776
	c.mu.Lock()
	c.balances[name] -= rupee
	c.mu.Unlock()
}

func main() {
<<<<<<< HEAD
	c := Bank{
=======
	c := Account{
>>>>>>> d1e27de128dbb838b71a55a60828b34a847da776
		balances: map[string]int{"anup": 0, "kashi": 0},
	}

	var wg sync.WaitGroup
<<<<<<< HEAD
	doDeposit := func(name string, n int) {
=======

	doDeposit := func(name string, n int) {

>>>>>>> d1e27de128dbb838b71a55a60828b34a847da776
		c.add(name, n)
		wg.Done()
	}

	doWithdraw := func(name string, n int) {
		if c.balances[name] < n {
			fmt.Println(nil, errors.New("Sorry , can not withdraw more than current balance amount"))
<<<<<<< HEAD
		} else {
			c.subtract(name, n)
		}
		wg.Done()
=======
			wg.Done()
		} else {
			c.subtract(name, n)
			wg.Done()
		}
>>>>>>> d1e27de128dbb838b71a55a60828b34a847da776
	}

	fmt.Println(c.balances)

	wg.Add(4)
	go doDeposit("anup", 10000)
	go doDeposit("kashi", 10000)
<<<<<<< HEAD

	go doWithdraw("anup", 1004)
	go doWithdraw("kashi", 104)
=======
	time.Sleep(time.Second)
	go doWithdraw("anup", 10001)
	go doWithdraw("kashi", 10001)
>>>>>>> d1e27de128dbb838b71a55a60828b34a847da776

	wg.Wait()
	fmt.Println(c.balances)
}
