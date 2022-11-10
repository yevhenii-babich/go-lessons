// Але найідеалогічне правильним шляхом, буде забезпечення синхронізації, використовуючи канали
// Don't communicate by sharing memory; share memory by communicating

package main

import (
	"errors"
	"fmt"
	"sync"
)

// Тепер обліковий запис це стуктура, що містить у собі канали, для змін
type AccountAsync struct {
	balance     float64
	deltaChan   chan float64
	balanceChan chan float64
	errChan     chan error
}

// Потрібен конструктор, щоб приховати від користувачів тонкощі реалізації
func NewAccount(balance float64) (a *AccountAsync) {
	a = &AccountAsync{
		balance:     balance,
		deltaChan:   make(chan float64),
		balanceChan: make(chan float64),
		errChan:     make(chan error, 1),
	}
	// і запустимо горутину, обслуговування операцій з акаунтом
	go a.run()
	return
}

// Просто читаємо з каналу балансу
func (a *AccountAsync) Balance() float64 {
	return <-a.balanceChan
}

// Записуємо кількість каналу змін
func (a *AccountAsync) Deposit(amount float64) error {
	a.deltaChan <- amount
	return <-a.errChan
}

// Аналогічно, насправді ця функція потрібна лише задля збереження семантики
func (a *AccountAsync) Withdraw(amount float64) error {
	a.deltaChan <- -amount
	return <-a.errChan
}

// Застосування змін до рахунку
func (a *AccountAsync) applyDelta(amount float64) error {
	stateStr := "Кладемо на рахунок"
	if amount < 0 {
		stateStr = "Знімаємо"
	}
	fmt.Println(stateStr, amount)

	newBalance := a.balance + amount
	if newBalance < 0 {
		return errors.New("Insufficient funds")
	}
	a.balance = newBalance
	return nil
}

// Нескінченний цикл обробника рахунку
// тепер скільки б горутин не робили операції над цим акаунтом
// Всі вони будуть синхронізовані тут і блокування вже не потрібні
func (a *AccountAsync) run() {
	var delta float64
	for {
		select {
		// Якщо надійшли зміни
		case delta = <-a.deltaChan:
			// Спробуємо їх застосувати
			a.errChan <- a.applyDelta(delta)
			// Якщо хтось запитує баланс
		case a.balanceChan <- a.balance:
			// Не робимо нічого, тому ми вже відправили відповідь
		}
	}
}

func main() {
	acc := NewAccount(20)

	// Стартуємо 10 go-lessons рутин
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(myG *sync.WaitGroup) {
			// Кожна з яких проводить операції з акаунтом
			defer myG.Done()
			for j := 0; j < 10; j++ {
				// Іноді знімає гроші
				if j%2 == 1 {
					if err := acc.Withdraw(50); err != nil {
						fmt.Printf("can't Withdraw :%v", err)
					}
					continue
				}
				// іноді кладе
				if err := acc.Deposit(50); err != nil {
					fmt.Printf("can't Deposit :%v", err)
				}
			}
		}(&wg)
	}
	// fmt.Scanln()
	// Тепер баланс завжди сходитиметься в 20
	wg.Wait()
	fmt.Println(acc.Balance())
}
