// Приклад використання  функції генератори, які захоплюють передані значення
// запускають go-lessons рутину і повертають канал, до якого надходитимуть сигнали
// Приклад time.After(t time.Duration)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Функція повертає канал
	c := boring("boring!")
	for i := 0; i < 15; i++ {
		// Читаємо з каналу
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string) <-chan string { // Повертаємо канал рядків тільки для читання.
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
	}()
	return c
}
