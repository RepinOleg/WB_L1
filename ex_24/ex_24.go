package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
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
