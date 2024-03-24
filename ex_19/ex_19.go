package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»).
// Символы могут быть unicode.

func main() {
	test := "главрыба"
	test2 := "English"
	fmt.Println(reverseString(test))
	fmt.Println(reverseString(test2))
}

func reverseString(input string) string {
	// Переводим строку в слайс рун
	runeStr := []rune(input)

	// идем с двух сторон по слайсу и меняем элементы местами
	for i, j := 0, len(runeStr)-1; i < j; i, j = i+1, j-1 {
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i]
	}

	// возвращаем строку
	return string(runeStr)
}
