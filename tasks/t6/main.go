package main

import (
	"context"
	"fmt"
	"log"
)

func eventPrinter(events chan interface{}, stopChan chan struct{}) {
	for {
		select {
		case event := <-events:
			fmt.Printf("printer 1 got event: %v\n", event)
		case <-stopChan: // в случае получения данных здесь или закрытия канала завершаемся
			log.Println("stopping event printer 1")
			return
		}
	}
}

func eventPrinterClose(events chan interface{}) {
	for {
		select {
		case event, ok := <-events: // проверяем закрытие канала с данными. Если закрыт, выходим
			if !ok {
				log.Println("stopping event printer 2")
				return
			}
			fmt.Printf("printer 2 got event: %v\n", event)
		}
	}
}

func eventPrinterContext(ctx context.Context, events chan interface{}) {
	for {
		select {
		case event := <-events:
			fmt.Printf("printer 3 got event: %v\n", event)
		case <-ctx.Done(): // Выходим, когда контекст завершен
			log.Println("stopping event printer 3")
			return
		}
	}
}

func main() {
	log.Default().SetFlags(log.Ldate | log.Lmicroseconds)
	ch := make(chan interface{})

	// Создаём отдельный канал, который будет сигнализировать об окончании горутины
	closed := make(chan struct{})
	go eventPrinter(ch, closed)

	// Пользуемся контекстом
	ctx, cancelFunc := context.WithCancel(context.Background())
	go eventPrinterContext(ctx, ch)

	// Просто закрываем канал с данными
	go eventPrinterClose(ch)

	for i := 0; i < 1000; i++ {
		ch <- i
	}
	close(closed)
	cancelFunc()
	close(ch)
}
