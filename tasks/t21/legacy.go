package main

import "fmt"

// Color Целевой интерфейс
type Color interface {
	UseColor()
}

// RGBA Какая-то структура, удовлетворявшая нас ранее
type RGBA struct {
	red, green, blue, alpha int
}

func (r *RGBA) UseColor() {
	fmt.Printf("Hey, i am rgb(%d, %d, %d, %d)\n", r.red, r.green, r.blue, r.alpha)
}

func Print(color Color) {
	color.UseColor()
}
