package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(createSet(s))
}

func createSet(s []string) map[string]bool {

	result := make(map[string]bool)
	// идем по словам в массиве строк и устанавливаем для них значение true
	for _, word := range s {
		result[word] = true
	}
	return result
}
