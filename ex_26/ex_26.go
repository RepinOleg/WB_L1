package main

import (
	"fmt"
	"strings"
)

func main() {

	test1 := "abcd"
	test2 := "abCdefAaf"
	test3 := "aabcd"
	test4 := "a"
	res := checkUniqueCharacters(test1)
	fmt.Println(res)
	res = checkUniqueCharacters(test2)
	fmt.Println(res)
	res = checkUniqueCharacters(test3)
	fmt.Println(res)
	res = checkUniqueCharacters(test4)
	fmt.Println(res)
}

func checkUniqueCharacters(str string) bool {
	// мапа для подсчета количества каждого символа в строке
	chars := make(map[rune]int)

	// Приводим строку к нижнему регистру
	lowerStr := strings.ToLower(str)

	// Идем по строке и прибавляем счетчик каждого символа
	for _, ch := range lowerStr {
		chars[ch]++
		if chars[ch] > 1 {
			return false
		}
	}
	return true
}
