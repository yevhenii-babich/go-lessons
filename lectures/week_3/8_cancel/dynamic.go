package main

import (
	"fmt"
	"sync"
	"time"
)

// Динамічно створюємо горутини, додаємо в лічильник очікування і чекаємо на завершення.
func main() {
	x := 11
	var wg sync.WaitGroup
	for i := 0; i < x; i++ {
		wg.Add(1)
		go func(in int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Printf("started %d\n", in)
			// Додаємо "час виконання"
			time.Sleep(time.Duration(in+1) * 300 * time.Millisecond)
			fmt.Printf("done %d\n", in)
		}(i, &wg)
	}
	wg.Wait()
}
