package main

import (
	"fmt"
	"sync"
)

//Написать программу,
//которая конкурентно рассчитает значение квадратов чисел
//взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

// Функция выводит в терминал квадрат числа
func squaring(number int) {
	fmt.Println(number * number)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, number := range numbers {
		// Добавляем к счетчику 1 за каждый проход цикла
		wg.Add(1)
		// Запускаем функцию внутри горутины
		go func(number int) {
			// убавляем счетчик
			defer wg.Done()

			// Вызываем функцию возведения в квадрат
			squaring(number)
		}(number)
	}
	// Ожидаем пока счетчик станет равен 0
	wg.Wait()
}
