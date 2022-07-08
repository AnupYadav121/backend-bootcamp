package main

import (
	"errors"
	"fmt"
	"sync"
)

type Account struct {
	mu      sync.Mutex
	id      int
	balance int
	name    string
}

type Bank struct {
	accounts []Account
}

func (b *Bank) add(rupee int, id int) {
	b.accounts[id].mu.Lock()
	b.accounts[id].balance += rupee
	b.accounts[id].mu.Unlock()
}

func (b *Bank) subtract(rupee int, id int) {
	b.accounts[id].mu.Lock()
	b.accounts[id].balance -= rupee
	b.accounts[id].mu.Unlock()
}

func main() {
	account1 := Account{
		balance: 0,
		name:    "anup",
		id:      0,
	}

	account2 := Account{
		balance: 0,
		name:    "kashi",
		id:      1,
	}

	b := Bank{
		[]Account{account1, account2},
	}

	var wg sync.WaitGroup

	doDeposit := func(name string, n int, id int) {
		b.add(n, id)
		wg.Done()
	}

	doWithdraw := func(name string, n int, id int) {
		if b.accounts[id].balance < n {
			fmt.Println(errors.New("Sorry , can not withdraw more than current balance amount"))
			wg.Done()
		} else {
			b.subtract(n, id)
			wg.Done()
		}
	}

	wg.Add(4)
	go doDeposit("anup", 10000, 0)
	go doDeposit("kashi", 10000, 1)
	go doWithdraw("anup", 100, 0)
	go doWithdraw("kashi", 1001, 1)

	wg.Wait()
	fmt.Println(b.accounts)
}
