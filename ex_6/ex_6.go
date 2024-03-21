package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// флаг для третьей горутины
var stopFlag bool

func main() {
	// вызов функции с первой горутиной
	closeGo1()

	// создание контекста с функцией отмены для второй горутины
	ctx, cancel := context.WithCancel(context.Background())

	// вызов функции со второй горутиной
	closeGo2(ctx)

	time.Sleep(3 * time.Second)
	// Вызов функции прекращения контекста
	cancel()
	time.Sleep(3 * time.Second)

	// создание wg для ожидания пока третья горутина завершится
	var wg sync.WaitGroup
	wg.Add(1)

	// Вызов функции с третьей горутиной
	closeGo3(&wg)

	time.Sleep(3 * time.Second)
	// устанавливаем флаг в true для того чтобы третья горутина вышла из цикла
	stopFlag = true
	wg.Wait()

}
func closeGo1() {
	// Канал который будем закрывать чтобы остановить горутину
	quit := make(chan bool)
	// запускаем горутину
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Остановка")
				// когда канал закрылся выходим из функции и горутина останавливается
				return
			default:
				fmt.Println("Горутина №1 работает")
				time.Sleep(time.Second)

			}
		}
	}()
	time.Sleep(3 * time.Second)
	// закрываем канал
	quit <- true
}

func closeGo2(ctx context.Context) {
	go func() {
		for {
			select {
			// Когда контекст зваершился останавливаем горутину
			case <-ctx.Done():
				fmt.Println("Остановка")
				return
			default:
				fmt.Println("Горутина №2 работает")
				time.Sleep(time.Second)
			}
		}
	}()
}

func closeGo3(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for {
			// когда флаг останавливаем горутину
			if stopFlag {
				fmt.Println("Остановка")
				return
			}

			fmt.Println("Горутина №3 работает")
			time.Sleep(time.Second)
		}
	}()
}
