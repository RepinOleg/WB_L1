package main

import (
	"fmt"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	mp := make(map[int][]float64)

	for _, val := range temps {
		// находим ключ - делим число на 10, отбрасываем дробную часть и умножаем обратно на 10
		key := int(val/10) * 10
		// добавляем в слайс значений val
		mp[key] = append(mp[key], val)
	}
	fmt.Println(mp)
}
