package main

import "fmt"

// Human Просто структура
type Human struct {
	height     int
	weight     int
	name       string
	secondName string
	health     int
}

// Die метод структуры Human
func (h Human) Die() {
	h.health = 0
	fmt.Printf("I am dead!")
}

// Go метод структуры Human
func (h Human) Go(distance float64) {
	h.weight -= int(0.1 * distance)
	h.health += int(0.1 * distance)
	fmt.Printf("I am tired!")
}

// Action cтруктура встраивает структуру Human
type Action struct {
	Human
}

func main() {
	act := Action{}
	act.Die()
}
