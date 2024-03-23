package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{1, 232, 1, 5, 1, 4, 1, 1}
	fmt.Println(QuickSort(arr))
}

func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	left, right := 0, len(array)-1
	// рандомим опорный элемент
	pivot := rand.Int() % len(array)

	// перемещаем его в конец
	array[pivot], array[right] = array[right], array[pivot]

	// проходим по слайсу и перемещаем наименьший элемент в конец слайса
	for i, _ := range array {
		if array[i] < array[right] {
			array[left], array[i] = array[i], array[left]
			left++
		}
	}
	// ставим наименьший элемент в конец
	array[left], array[right] = array[right], array[left]

	// рекурсивно вызываем ту же функцию для левой и правой части
	QuickSort(array[:left])
	QuickSort(array[left+1:])
	return array
}
