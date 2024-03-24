package main

import "time"

// Реализовать собственную функцию sleep.

func Sleep(t time.Duration) {
	select {
	// Функция time.After возвращает канал и пишет в него спустя t времени
	case <-time.After(t):
		// Пока в канал ничего не пришло main поток блокируется
	}
}

func main() {
	Sleep(5 * time.Second)
}
