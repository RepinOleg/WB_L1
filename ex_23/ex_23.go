package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 4, 5}
	res, _ := deleteElem(test, 2)
	fmt.Println(res, test)
}

// Функция принимает слайс и индекс элемента, возвращает обрезаный слайс и ошибку
func deleteElem(numb []int, index int) ([]int, error) {
	if index < 0 || index >= len(numb) {
		return nil, fmt.Errorf("index out of range")
	}
	// Нарезаем слайс - берем левую часть от индекса и к ней прибавляем значения из правой части от индекса
	res := append(numb[:index], numb[index+1:]...)
	return res, nil
}
