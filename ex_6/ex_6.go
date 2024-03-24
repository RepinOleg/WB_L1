package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

var stopFlag bool

func main() {
	var wg sync.WaitGroup
	// вызов функции с первой горутиной
	closeGo1(&wg)

	// вызов функции со второй горутиной
	closeGo2(&wg)

	// создание wg для ожидания пока третья горутина завершится

	// Вызов функции с третьей горутиной
	closeGo3(&wg)
	// устанавливаем флаг в true для того чтобы третья горутина вышла из цикла
	stopFlag = true
	wg.Wait()

}

// Остановка с помощью канала
func closeGo1(wg *sync.WaitGroup) {
	// Канал который будем закрывать чтобы остановить горутину
	quit := make(chan bool)
	wg.Add(1)
	// запускаем горутину
	go func() {
		defer wg.Done()
		for {
			select {
			case <-quit:
				fmt.Println("Остановка 1 горутины")
				// когда канал закрылся выходим из функции и горутина останавливается
				return
			default:
				fmt.Println("Горутина №1 работает")
				time.Sleep(time.Second)

			}
		}
	}()
	time.Sleep(time.Second)
	// закрываем канал
	quit <- true
}

// Остановка с помощью контекста
func closeGo2(wg *sync.WaitGroup) {
	// создание контекста с функцией отмены для второй горутины
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			// Когда контекст зваершился останавливаем горутину
			case <-ctx.Done():
				fmt.Println("Остановка 2 горутины")
				return
			default:
				fmt.Println("Горутина №2 работает")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(time.Second)
	// Вызов функции прекращения контекста
	cancel()
}

// остановка с помощью глобальной переменной
func closeGo3(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// когда флаг = true останавливаем горутину
			if stopFlag {
				fmt.Println("Остановка 3 горутины")
				return
			}

			fmt.Println("Горутина №3 работает")
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second)
}
