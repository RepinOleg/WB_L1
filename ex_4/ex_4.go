package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.

// Функция для вывода данных из канала
func readFromChannel(idGoroutine int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("Goroutine number - %d, Number from channel: %d\n", idGoroutine, val)
	}
}

func main() {
	// Считываем количество воркеров
	var workers int
	fmt.Println("Enter amount of workers")
	count, err := fmt.Scanln(&workers)
	if err != nil || count != 1 {
		log.Fatal("Wrong input")
	}

	ch := make(chan int)
	var wg sync.WaitGroup
	// Запускаем горутины которые будут ждать пока в канал попадут данные
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go readFromChannel(i, ch, &wg)
	}

	// Создаем еще один канал который отвечает за получение сигнала
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT)

	// запускаем цикл который будет писать в канал пока не поступит SIGINT
	// метка(label) для внешнего цикла
loop:
	for {
		select {
		// когда сигнал поступил, закрываем канал, выходим из внешнего цикла
		case <-signalCh:
			close(ch)
			wg.Wait()
			fmt.Println("Sigint - channel closed, exit")
			break loop // выходим из внешнего цикла
		case ch <- rand.Int(): // пишем числа в канал
			fmt.Println("writing to channel")
			time.Sleep(time.Second)
		}
	}
}
