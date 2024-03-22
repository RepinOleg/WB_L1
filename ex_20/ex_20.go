package main

import (
	"fmt"
	"strings"
)

func main() {
	test := "snow dog sun"
	fmt.Println(reverseWords(test))
}

func reverseWords(input string) string {
	// Разбиваем строку на подстроки разделяя слова по пробелам
	slWords := strings.Split(input, " ")

	// инициализируем стринг Билдер
	var bld strings.Builder

	// Идем по массиву строк начиная с конца
	for i := len(slWords) - 1; i >= 0; i-- {
		// Добавляем в билдер слово и пробел
		bld.WriteString(slWords[i])
		bld.WriteString(" ")
	}
	// возвращаем строку из билдера
	return bld.String()
}
