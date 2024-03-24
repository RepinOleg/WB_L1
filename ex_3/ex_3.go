package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….)
//с использованием конкурентных вычислений.

// Функция возвращает квадрат числа
func squaring(number int) int {
	return number * number
}

func main() {
	numbers := []int{1, 1, 1, 1, 1, 2, 4, 6, 8, 10}
	var (
		sum   int
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	for _, number := range numbers {
		// Добавляем к счетчику 1 за каждый проход цикла
		wg.Add(1)
		// Запускаем функцию внутри горутины
		go func(number int) {
			// убавляем счетчик
			defer wg.Done()

			// Вызываем функцию возведения в квадрат
			res := squaring(number)
			mutex.Lock()
			sum += res
			mutex.Unlock()
		}(number)
	}

	// Ожидаем пока счетчик станет равен 0
	wg.Wait()
	fmt.Println(sum)
}
