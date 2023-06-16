package bank

import (
	"sync"
)

type BankMonitor interface {
	Init()
	Wait()
	Signal()
	GetBalance() int
	SetBalance()
	Deposit(int)
	Withdraw(int) bool
}

type Bank struct { // all variable are private, i.e., non-capital letter
	mutex         *sync.Mutex
	balance       int
	isInitialized bool
}

func (b *Bank) Init() {
	b.mutex = &sync.Mutex{}
	b.balance = 0
	b.isInitialized = true
}
func (b *Bank) Wait() {
	if b.isInitialized {
		b.mutex.Lock()
	}
}
func (b *Bank) Signal() {
	if b.isInitialized {
		b.mutex.Unlock()
	}
}

func (b Bank) GetBalance() int {
	return b.balance
}
func (b *Bank) SetBalance(balance int) {
	b.Wait()
	// critical section
	b.balance = balance
	// critical section done
	b.Signal()
}
func (b *Bank) Deposit(amount int) {
	b.Wait()
	b.balance = b.balance + amount
	b.Signal()
}

func (b *Bank) Withdraw(amount int) bool {
	var r bool

	b.Wait()
	if (b.balance - amount) >= 0 {
		b.balance = b.balance - amount
		r = true
	} else {
		r = false
	}
	b.Signal()
	return r
}


