package main

import (
	"fmt"
	"sync"
)

func PrintRandom() {
	nums := []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{} // Вейтгруппа для контроля за выполнением горутин
	for _, el := range nums {
		wg.Add(1)
		go func(el int) {
			fmt.Printf("%d ", el*el) // Вывод происходит в рандомном порядке
			wg.Done()
		}(el)
	}
	wg.Wait()
	fmt.Println()
}

func PrintCorrect() {
	nums := []int{2, 4, 6, 8, 10}
	res := make([]int, len(nums)) // слайс чтобы расположить результаты в правильном порядке
	wg := &sync.WaitGroup{}
	for i, el := range nums {
		wg.Add(1)
		go func(el, i int, res *[]int) {
			(*res)[i] = el * el
			wg.Done()
		}(el, i, &res)
	}
	wg.Wait()
	fmt.Println(res)
}

func main() {
	PrintRandom()
	PrintCorrect()
}
