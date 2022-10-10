package main

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// будет записывать группы по 10 градусов в соответствующую мапу
	groups := make(map[int][]float64)
	for _, temp := range temperatures {
		// Приводим вещественное число к целому, целочисленно делим на 10, результат умножаем на 10.
		// таким образом получим число, делящееся на 10, соответствующее данной переменной
		groups[int(temp)/10*10] = append(groups[int(temp)/10*10], temp)
	}
	fmt.Println(groups)
}
