package main

import "fmt"

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
