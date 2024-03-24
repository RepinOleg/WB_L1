package main

import (
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить? Приведите корректный пример реализации.
//
//
//	var justString string
//	func someFunc() {
//  	v := createHugeString(1 << 10)
//  	justString = v[:100]
//	}
//
//	func main() {
//  	someFunc()
//	}

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
}

func createHugeString(size int) string {
	return string(make([]byte, size))
}
