// Поки в прикладах нам було не так важливо, що всі горутини закінчили роботу,
// тк при виході основного процесу, всі горутини завершаться,
// Але, що якщо нам потрібно переконатися в завершенні всіх виконаних робіт
// Можна це зробити за допомогою каналу
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(die chan bool) <-chan string { // Повертаємо канал рядків тільки для читання.
	c := make(chan string)
	go func() {
		for {
			select {
			case c <- fmt.Sprintf("boring %d", rand.Intn(100)):
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			case <-die:
				fmt.Println("Jobs done!")
				die <- true
				return
			}
		}
	}()
	return c
}

func main() {
	die := make(chan bool)
	res := boring(die)

	for i := 0; i < 5; i++ {
		// Читаємо з каналу
		fmt.Printf("You say: %q\n", <-res)
	}
	die <- true
	// Чекаємо, поки всі горутини закінчать виконуватися
	<-die
}
