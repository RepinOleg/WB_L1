package main

import (
	"fmt"
	"log"
)

func main() {
	var number int64 = 10
	SetBit(4, 1, &number)
	fmt.Println(number)
}

func SetBit(pos, val int, number *int64) {
	// Сдвигаем 1 на количество pos
	var mask int64 = 1 << pos

	if val == 1 {
		// Применяем операцию ИЛИ. Если в маске или числе есть 1 то в результате будет 1
		*number = *number | mask
	} else if val == 0 {
		// Применяем операцию И НЕ. Если в маске 1, то в результате будет 0. Если в маске 0, то берется значение бита из числа
		*number = *number &^ mask
	} else {
		log.Fatal("Wrong val")
	}
}
