// Але набагато простіше поєднувати логіку каналів через оператор select
// Він дозволяє перевірити можливість операції з кількома каналами відразу.
// Важливо, якщо немає доступних варіантів і немає блоку default, підпрограма заблокується
// Якщо доступні до роботи більше одного каналу, вибирається довільний
package main

import (
	"fmt"
	"time"
)

func main() {
	// Створюємо пару каналів
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "chan 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "chan 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
	Out:
		for {
			select {
			case msg1, ok := <-c1:
				if !ok {
					break Out
				}
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
				// time.After повертає канал, в який запис відбудеться через 1 секунду
			case <-time.After(time.Second):
				fmt.Println("timeout")
			}
			fmt.Println("next read")
		}
		fmt.Println("chanel closed")
	}()

	fmt.Scanln()
	close(c1)
	fmt.Scanln()
}
