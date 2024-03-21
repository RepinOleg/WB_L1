package main

import "fmt"

func main() {
	var x, y int = 0, 10
	fmt.Println("Before", x, y)
	x, y = y, x
	fmt.Println("After", x, y)
}
