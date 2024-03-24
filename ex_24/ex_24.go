package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния
// между двумя точками, которые представлены в виде структуры Point
// с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x, y float64 // инкапсуляция полей
}

// NewPoint конструктор принимает координаты точки
func NewPoint(x, y float64) *Point {
	// возвращает указатель на обьект класса Point
	return &Point{x, y}
}

// Функия вычисляет дистанцию между двумя точками
func distanceCalculation(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func main() {
	firstPoint := NewPoint(1, 1)
	secondPoint := NewPoint(-1, -1)
	fmt.Println(distanceCalculation(firstPoint, secondPoint))
}
