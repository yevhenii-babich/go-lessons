// Якщо при запуску програми всі go-lessons рутини виявляться в стані очікування,
// runtime викине паніку, з повідомленням all goroutines are sleeping
package main

import (
	"fmt"
	"time"
)

// Ball is just a ball
type Ball struct{ hits int }

func main() {
	// Створюємо канал взаємодії гравців
	table := make(chan *Ball)
	// Стартуємо пару гравців
	go player("ping", table)
	go player("pong", table)

	// table <- new(Ball) // Запуск м'яча у гру
	time.Sleep(1 * time.Second)
	<-table // Кінець гри, забираємо м'яч
}

func player(name string, table chan *Ball) {
	for {
		// Чекаємо, коли м'яч потрапив до гравця
		ball := <-table
		// Збільшуємо лічильник ударів
		ball.hits++
		fmt.Println(name, ball.hits)
		// Чекаємо трохи
		time.Sleep(100 * time.Millisecond)
		// Відправляємо м'яч назад у канал
		// Важливо, програма заблокується, доки інший гравець звідти не прочитає
		table <- ball
	}
}
