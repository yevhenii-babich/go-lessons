// Go рутини, що одночасно працюють із загальними даними самі собою не можуть синхронізуватися
package main

import (
	"fmt"
	"log"
)

// Нехай у нас є Рахунок
type Account struct {
	balance float64
}

func (a *Account) Balance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) {
	log.Printf("depositing: %f (%f)", amount, a.balance)
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) {
	if amount > a.balance {
		log.Printf("not enouth : %f", a.balance)
		return
	}
	log.Printf("withdrawing: %f (%f)", amount, a.balance)
	a.balance -= amount
}

func main() {
	acc := Account{}

	// Стартуємо 10 go-lessons рутин
	for i := 0; i < 10; i++ {
		go func() {
			// Кожна з яких проводить операції з акаунтом
			for j := 0; j < 10; j++ {
				// Іноді знімає гроші
				if j%2 == 1 {
					acc.Withdraw(50)
					continue
				}
				// іноді кладе
				acc.Deposit(50)
			}
		}()
	}
	_, _ = fmt.Scanln()
	// Що ж вийде в результаті
	fmt.Printf("balance: %f", acc.Balance())
}

// // //
// func closure() {
// // Функції захоплюють змінні у сфері видимості
// // Але щоб передавати значення, потрібно явно передавати їх у функцію
// for i: = 0; i < 10; i++ {
// go func() {
// fmt.Println("Got", i)
// } ()
// }
// fmt.Scanln()
// }
