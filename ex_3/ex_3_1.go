package main

import (
	"fmt"
	"sync"
)

func squaring2(number int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Пишем в канал квадраты чисел
	ch <- number * number
}

func main() {
	numbers := []int{1, 1, 1, 2, 2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	// создаем канал
	ch := make(chan int)

	for _, number := range numbers {
		wg.Add(1)
		go squaring2(number, ch, &wg)
	}

	// запускаем поток в котором ждем завершения всех горутин после чего закрываем канал
	go func() {
		wg.Wait()
		close(ch)
		fmt.Println("Channel closed")
	}()

	// считываем все данные из канала пока он открыт
	var sum int
	for number := range ch {
		fmt.Println("reading")
		sum += number
	}

	fmt.Println(sum)
}
