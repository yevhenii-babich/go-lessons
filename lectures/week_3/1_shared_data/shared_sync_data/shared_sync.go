package main

import (
	"fmt"
	"log"
	"time"
)

// Нехай у нас є Рахунок
type Account struct {
	balance float64
	debit   float64
	credit  float64
	changes chan float64
}

func (a *Account) BankWorker() {
	for {
		changes, ok := <-a.changes
		if !ok {
			break
		}
		log.Printf("processing request: %f, balance: %f\n", changes, a.balance)
		if a.balance+changes < 0 {
			log.Printf("declined changes: %.2f\n", changes)
			continue
		}
		a.balance += changes
		switch {
		case changes > 0:
			a.debit += changes
		case changes < 0:
			a.credit -= changes
		}
	}
	log.Println("bank closed")
}

func (a *Account) Deposit(amount float64) {
	log.Printf("depositing: %f (%f)\n", amount, a.Balance())
	a.changes <- amount
}

func (a *Account) Withdraw(amount float64) {
	log.Printf("withdrawing: %f (%f)\n", amount, a.balance)
	a.changes <- -amount
}

func (a *Account) Balance() string {
	return fmt.Sprintf("balance: %0.2f, debit: %0.2f, credit: %0.2f", a.balance, a.debit, a.credit)
}

func (a *Account) Done() {
	close(a.changes)
}

func main() {
	acc := &Account{changes: make(chan float64, 100)}
	go acc.BankWorker()
	// Стартуємо 10 go-lessons рутин
	for i := 0; i < 10; i++ {
		go func(acc *Account, no int) {
			start := time.Now()
			// Кожна з яких проводить операції з акаунтом
			for j := 0; j < 10; j++ {
				// Іноді знімає гроші
				if j%3 == 0 {
					acc.Withdraw(50)
					continue
				}
				// іноді кладе
				acc.Deposit(50)
			}
			log.Printf("# %d completed in %v", no, time.Since(start))
		}(acc, i)
	}
	_, _ = fmt.Scanln()
	acc.Done()
	// Що ж вийде в результаті
	fmt.Println(acc.Balance())
}
