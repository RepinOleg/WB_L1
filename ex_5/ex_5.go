package main

import (
	"fmt"
	"math/rand"
	"time"
)

func writeNumbers(ch chan<- int, N int) {
	// Функция After возвращает канал в который придет значение после истечения времени
	finish := time.After(time.Duration(N) * time.Second)

	for {
		// Рандомим число
		numb := rand.Int()
		select {
		// Если в канал finish пришло значение, то закрываем канал с данными и выходим из функции
		case <-finish:
			close(ch)
			return
		default:
			// Иначе пишем число в канал
			ch <- numb
		}
	}
}

func main() {
	ch := make(chan int)
	//N количество секунд
	var N = 2
	//Запуск потока
	go writeNumbers(ch, N)
	// чтение из канала
	for numb := range ch {
		fmt.Println(numb)
	}

}
