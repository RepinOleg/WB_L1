package main

import "fmt"

func main() {
	// создаем два канала
	var in, out = make(chan int), make(chan int)
	// запускаем горутину которая пишет в канал
	go func() {
		for i := 1; i <= 10; i++ {
			in <- i
		}
		close(in)
	}()
	// запускаем горутину которая читает из канала
	go func() {
		// в конце закрываем канал out
		defer close(out)

		for {
			// если канал закрылся выходим из функции
			num, ok := <-in
			if !ok {
				return
			}
			out <- num * num
		}
	}()
	// Считываем все значения которые лежат в канале out
	for num := range out {
		fmt.Println(num)
	}
}
