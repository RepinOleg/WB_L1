package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

func binarySearch(numbers []int, target int) int {
	// Переменный для левого и правого края поиска
	left, right := 0, len(numbers)-1

	for left <= right {
		// Элемент в середине между left и right для поиска target относительно центрального элемента
		middle := (left + right) / 2

		// Сдвигаем right до middle - 1, потому что искомый элемент находится слева от middle элемента
		if numbers[middle] > target {
			right = middle - 1
		} else if numbers[middle] < target {
			// Сдвигаем left до middle + 1, потому что искомый элемент находится справа от middle элемента
			left = middle + 1
		} else {
			// Элемент найден возвращаем его индекс
			return middle
		}

	}
	return -1
}

func main() {
	test := []int{1, 3, 4, 6, 8, 10, 55, 56, 59, 70, 79, 81, 91, 10001}
	fmt.Println(binarySearch(test, 55))
}
