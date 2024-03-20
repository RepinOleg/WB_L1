package main

import (
	"fmt"
	"sync"
)

// Функция возвращает квадрат числа
func squaring(number int) int {
	return number * number
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	var sum int
	for _, number := range numbers {
		// Добавляем к счетчику 1 за каждый проход цикла
		wg.Add(1)
		// Запускаем функцию внутри горутины
		go func(number int) {
			// убавляем счетчик
			defer wg.Done()

			// Вызываем функцию возведения в квадрат
			sum += squaring(number)
		}(number)
	}

	// Ожидаем пока счетчик станет равен 0
	wg.Wait()
	fmt.Println(sum)
}
