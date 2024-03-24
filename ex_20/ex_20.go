package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func main() {
	test := "snow dog sun"
	fmt.Println(reverseWords(test))
	fmt.Println(len(test), len(reverseWords(test)))
}

func reverseWords(input string) string {
	if input == "" {
		return ""
	}
	// Разбиваем строку на подстроки разделяя слова по пробелам
	slWords := strings.Split(input, " ")

	// инициализируем стринг Билдер
	var bld strings.Builder

	// Идем по массиву строк начиная с конца
	for i := len(slWords) - 1; i >= 0; i-- {
		// Добавляем в билдер слово и пробел
		bld.WriteString(slWords[i])

		// после последнего слова добавлять пробел не нужно
		if i != 0 {
			bld.WriteString(" ")
		}
	}
	// возвращаем строку из билдера
	return bld.String()
}
