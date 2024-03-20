package main

import (
	"fmt"
	"sync"
)

func squaring2(number int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- number * number
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	ch := make(chan int)

	for _, number := range numbers {
		wg.Add(1)
		go func(number int) {
			squaring2(number, ch, &wg)
		}(number)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	var sum int
	for number := range ch {
		sum += number
	}
	fmt.Println(sum)
}
