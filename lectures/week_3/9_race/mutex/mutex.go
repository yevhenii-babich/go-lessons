// Go рутини, що одночасно працюють із загальними даними самі собою не можу синхронізуватися
// Як приклад захисту даних від небезпечних змін можна використовувати структуру Mutex
// У неї також немає громадських полів, проте, є два громадських способу
// Lock(), який дасть лише одній підпрограмі продовжити виконання блоку, решта заблокується в очікуванні
// Unlock(), який знімає лок, захоплений раніше
package main

import (
	"fmt"
	"log"
	"sync"
)

// Нехай у нас є Рахунок
// Вбудуємо в нього об'єкт Mutex
// тепер наш об'єкт може використати його публічні методи
type AccountProtected struct {
	sync.Mutex
	balance float64
}

func (a *AccountProtected) Balance() float64 {
	a.Lock()
	defer a.Unlock()
	return a.balance
}

func (a *AccountProtected) Deposit(amount float64) {
	a.Lock()
	defer a.Unlock()
	log.Printf("depositing: %f", amount)
	a.balance += amount
}

func (a *AccountProtected) Withdraw(amount float64) {
	a.Lock()
	defer a.Unlock()
	if amount > a.balance {
		return
	}
	log.Printf("withdrawing: %f", amount)
	a.balance -= amount
}

func main() {
	acc := AccountProtected{}
	var wg sync.WaitGroup
	// Стартуємо 10 go-lessons рутин
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(myGroup *sync.WaitGroup) {
			defer myGroup.Done()
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
		}(&wg)
	}
	wg.Wait()
	// Тепер баланс завжди сходитиметься в 0
	fmt.Println(acc.Balance())
}
