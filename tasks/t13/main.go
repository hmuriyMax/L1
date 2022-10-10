package main

import "fmt"

func main() {
	a := 2
	b := 52
	fmt.Printf("a = %d, b = %d\n", a, b)

	// Используем свойство xor: a ^ b ^ b = a, b ^ a ^ a = b
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("a = %d, b = %d\n", a, b)
}
