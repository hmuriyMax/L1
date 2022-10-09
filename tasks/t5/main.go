package main

import (
	"context"
	"fmt"
	"time"
)

func channelWriter(ctx context.Context, ch chan interface{}) {
	var (
		i    int
		done bool
	)
	// Ожидаем завершения контекста
	go func() {
		select {
		case <-ctx.Done():
			done = true
			return
		}
	}()
	for {
		if done {
			return
		}
		ch <- i
		fmt.Printf("sent %v\n", i)
		i++
	}
}

func channelReader(ctx context.Context, ch chan interface{}) {
	for {
		select {
		case el := <-ch:
			fmt.Printf("got %v\n", el)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	const seconds = 4

	// Создаём контекст с таймаутом seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(seconds))
	ch := make(chan interface{})
	defer cancel()
	go channelReader(ctx, ch)
	go channelWriter(ctx, ch)
	select {
	case <-ctx.Done():
		return
	}
}
