package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться
// флаг для третьей горутины

func writeNumbers(ch chan<- int, N int) {
	// Функция After возвращает канал в который придет значение после истечения времени
	finish := time.After(time.Duration(N) * time.Second)

	for {
		select {
		// Если в канал finish пришло значение, то закрываем канал с данными и выходим из функции
		case <-finish:
			close(ch)
			fmt.Println("channel closed")
			return
		case ch <- rand.Int():
		}
	}
}

func main() {
	ch := make(chan int)
	//N количество секунд
	var N = 1
	//Запуск потока
	go writeNumbers(ch, N)
	// чтение из канала
	for numb := range ch {
		fmt.Println(numb)
	}
}
