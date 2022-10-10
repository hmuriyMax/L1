package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	// Там под капотом вроде как quicksort
	sort.Ints(arr[:])
	fmt.Println(arr)
}
