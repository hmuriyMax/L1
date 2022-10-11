package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Удаление без сохранения позиций элементов, работает за О(1)
func del(slice []int, position int) []int {
	ln := len(slice)
	slice[position] = slice[ln-1]
	return slice[:ln-1]
}

// Удаление с сохранением позиций элементов, работает за О(n)
func delSave(slice []int, position int) []int {
	for i := position; i < len(slice)-1; i++ {
		slice[i] = slice[i+1]
	}
	return slice[:len(slice)-1]
}

func main() {
	rand.Seed(time.Now().UnixNano())
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)
	slice = del(slice, rand.Intn(len(slice)))
	fmt.Println(slice)
	slice = delSave(slice, rand.Intn(len(slice)))
	fmt.Println(slice)
}
