package workerpool

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Task - опис інтрефейсу роботи
type Task interface {
	Execute(timeMS int)
}

// Pool - структура, нам знадобиться Мутекс, для гарантій атомарності змін самого об'єкта
// Канал вхідних завдань
// Канал скасування, для завершення роботи
// WaitGroup для контролю завершення робіт
type DataTask struct {
	Task  Task
	Param int
}

type Pool struct {
	mu     sync.Mutex
	taskNo int
	size   int
	tasks  chan DataTask
	kill   chan struct{}
	wg     sync.WaitGroup
}

// Прихуємо внутрішній пристрій за конструктором, користувач може впливати тільки на розмір пула
func NewPool(size int) *Pool {
	pool := &Pool{
		// Канал завдань - буферизований, щоб основна програма не блокувалася під час постановки завдань
		tasks: make(chan DataTask, 128),
		// Канал kill для вбивства "зайвих воркерів"
		kill: make(chan struct{}),
	}
	// Викличемо метод resize, щоб встановити відповідний розмір пулу
	pool.Resize(size)
	return pool
}

// Життєвий цикл воркера
func (p *Pool) worker(no int) {
	defer p.wg.Done()
	log.Printf("worker #%d started", no)
	for {
		select {
		// Якщо є завдання, її потрібно обробити
		case task, ok := <-p.tasks:
			if !ok {
				log.Printf("worker #%d closed", no)
				return
			}
			log.Printf("worker #%d got new task:%v\n", no, task)
			task.Task.Execute(task.Param)
			// Якщо прийшов сигнал помирати, виходимо
		case <-p.kill:
			log.Printf("worker #%d killed", no)
			return
		}
	}
}

func (p *Pool) Resize(n int) {
	// Захоплюємо лок, щоб уникнути одночасної зміни стану
	p.mu.Lock()
	defer p.mu.Unlock()
	for p.size < n {
		p.taskNo++
		p.size++
		p.wg.Add(1)
		go p.worker(p.taskNo)
	}
	for p.size > n {
		p.size--
		p.kill <- struct{}{}
	}
}

func (p *Pool) Close() {
	close(p.tasks)
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func (p *Pool) Exec(task Task, param int) {
	p.tasks <- DataTask{Task: task, Param: param}
}

type ExampleTask string

func (e ExampleTask) Execute(timeMS int) {
	fmt.Println("executing:", string(e))
	time.Sleep(time.Duration(timeMS) * time.Millisecond)
	fmt.Println("finishing:", string(e))
}
