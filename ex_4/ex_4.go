package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Функция для вывода данных из канала
func makeWorkers(idGoroutine int, ch chan string) {
	for {
		fmt.Printf("Goroutine number - %d, String from channel: %s\n", idGoroutine, <-ch)
	}
}
func main() {
	// Считываем количество воркеров
	var workers int
	fmt.Println("Input amount of workers")
	count, err := fmt.Scanln(&workers)
	if err != nil || count != 1 {
		log.Fatal("Wrong input")
	}

	ch := make(chan string)
	// Запускаем горутины которые будут ждать пока в канал попадут данные
	for i := 1; i <= workers; i++ {
		go makeWorkers(i, ch)
	}

	// Создаем еще один канал который отвечает за получение сигнала
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGKILL)

	// запускаем горутину которая будет ожидать пока в канал поступит сигнал
	go func() {
		<-signalCh
		// закрываем канал с данными
		close(ch)
	}()

	// Читаем в канал
	for {
		var input string
		fmt.Println("Enter a string")
		fmt.Scanln(&input)
		ch <- input
	}
}
