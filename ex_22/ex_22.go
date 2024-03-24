package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит,
// складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	x, _ := big.NewInt(0).SetString("123123123123123123", 10)
	y, _ := big.NewInt(0).SetString("678678678678678678678", 10)

	// используем библиотеку big
	var res big.Int

	// функции возвращают *big.Int
	fmt.Println(res.Add(x, y))
	fmt.Println(res.Div(x, y))
	fmt.Println(res.Mul(x, y))
	fmt.Println(res.Sub(x, y))
}
