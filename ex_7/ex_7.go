package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Реализовать конкурентную запись данных в map.

const amountNumbers = 20

func main() {
	// создаем мапу, она хранит число и его наличие
	numbers := make(map[int]bool, amountNumbers)

	// создаем переменные для контроля над потоками
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)
	// запускаем цикл в котором запускаем 10 горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeNumbersToMap(numbers, &wg, &mutex)
	}

	wg.Wait()
	fmt.Println(numbers)
}

// функция записывает 10 чисел в мапу
func writeNumbersToMap(mp map[int]bool, wg *sync.WaitGroup, mutex *sync.Mutex) {

	defer wg.Done()
	// блокируем мьютекс чтобы только одна горутина имела доступ к данным
	mutex.Lock()
	// разблокируем мьютекс в конце выполнения функции
	defer mutex.Unlock()
	for {
		numb := rand.Intn(amountNumbers)
		if !mp[numb] {
			mp[numb] = true
			break
		}

	}

}
