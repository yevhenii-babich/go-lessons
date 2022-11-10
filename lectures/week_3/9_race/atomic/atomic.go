// Також, коли саме загальне значення - це число
// Можна використовувати пакет atomic, щоб гарантувати черговість змін об'єкта
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type atomicCounter struct {
	val int64
}

func (c *atomicCounter) Add(x int64) {
	atomic.AddInt64(&c.val, x)
}

func (c *atomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}

func main() {
	counter := atomicCounter{}

	// Якщо запустити цей код із ключем race, можна помітити, що ми ніяк не гарантуємо
	// Завершення всіх робіт, щоб гарантії були, варто використовувати WaitGroup
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(no int) {
			// Важливо, що в рамках цього циклу управління між
			// горутинами не може перемикатися, якщо ми хочемо, щоб робота йшла паралельно
			// потрібно використовувати виклик Gosched (управління перемикається тільки на операторі select та роботі з ОС, такий як читання з файлів або мережі)
			for i := 0; i < 10000; i++ {
				counter.Add(1)
				if i%100 == 0 {
					runtime.Gosched()
				}
			}
			wg.Done()
		}(i)
	}

	// time.Sleep(time.Second)
	wg.Wait()
	fmt.Println(counter.Value())
}
