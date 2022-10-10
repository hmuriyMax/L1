package main

import "fmt"

func multiplier(in <-chan int, out chan<- int) {
	for {
		select {
		case num, ok := <-in:
			if !ok {
				return
			}
			out <- num * 2
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	senderChan := make(chan int)
	receiverChan := make(chan int)
	go multiplier(senderChan, receiverChan)
	for _, num := range nums {
		fmt.Printf("sent: %d\n", num)
		senderChan <- num
		fmt.Printf(" got: %d\n", <-receiverChan)
	}
}
