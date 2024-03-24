package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализовать структуру-счетчик,
// которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

type Count struct {
	count int
}

func (c *Count) Increment(mutex *sync.Mutex) {
	// в конце выполнения функции разблокируем переменную
	defer mutex.Unlock()
	// Перед тем как инкрементировать счетчик блокируем доступ к переменной count
	// чтобы только один поток имел доступ к ней
	mutex.Lock()
	c.count++
}

func main() {

	var mutex sync.Mutex
	c := Count{0}

	for i := 0; i < 1000; i++ {
		// передаем в функцию мьютекс
		go c.Increment(&mutex)
	}

	time.Sleep(time.Second)
	fmt.Println(c.count)
}
