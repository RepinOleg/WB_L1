package main

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	// Если оставить такую запись, то после выхода из этой функции
	// в памяти будет храниться огромная строка которая нам не нужна
	// justString = v[:100]

	// Чтобы не хранить в памяти огромную строку после того как эта функция завершится
	// Нужно с помощью функции Clone скопировать срез в новый участок памяти и тогда GC почистит память от большой строки
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
	fmt.Println(len(justString))
}

func createHugeString(size int) string {
	return string(make([]byte, size))
}
