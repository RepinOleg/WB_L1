package main

import "fmt"

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
	test := []int{1, 2, 3, 4, 5, 6, 22}
	fmt.Println(binarySearch(test, 4))
}
