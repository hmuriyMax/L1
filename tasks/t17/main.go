package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var arr []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(10))
	}
	sort.Ints(arr)
	fmt.Println(arr)
	res := sort.SearchInts(arr, 5)
	if arr[res] != 5 {
		fmt.Printf("Not found")
	} else {
		fmt.Printf("Found %d", res)
	}
}
