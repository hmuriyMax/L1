package main

import (
	"fmt"
	"sync"
)

func SquareSum() {
	nums := []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	mu := sync.Mutex{}
	sum := 0
	for _, el := range nums {
		wg.Add(1)
		go func(el int) {
			sqr := el * el
			mu.Lock() // Мьютекс, чтобы избежать гонки
			sum += sqr
			mu.Unlock()
			wg.Done()
		}(el)
	}
	wg.Wait()
	fmt.Println(sum)
}

func main() {
	SquareSum()
}
