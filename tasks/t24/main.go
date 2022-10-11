package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

// GenPoint Конструктор
func GenPoint(x, y float64) Point {
	return Point{x, y}
}

// Distance Считает расстояние
func Distance(a, b Point) float64 {
	return math.Sqrt(math.Pow(b.x-a.x, 2) + math.Pow(b.y-a.y, 2))
}

func main() {
	a := GenPoint(0, 0)
	b := GenPoint(3, 4)
	fmt.Println(Distance(a, b))
}
