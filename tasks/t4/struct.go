package t4

import (
	"fmt"
	"log"
	"sync"
)

// ChannelPrinter Структура канала. Num - количество воркеров, wg - вейтгруппа для них
type ChannelPrinter struct {
	ch  chan interface{}
	Num int
	wg  sync.WaitGroup
}

func (cp *ChannelPrinter) Launch() {
	cp.ch = make(chan interface{})
	if cp.Num <= 0 {
		cp.Num = 1
	}
	for i := 0; i < cp.Num; i++ {
		cp.wg.Add(1)
		log.Printf("Launching worker %d", i)
		go cp.worker()
	}
}

func (cp *ChannelPrinter) GetChan() *chan interface{} {
	return &cp.ch
}

func (cp *ChannelPrinter) Stop() {
	close(cp.ch)
	cp.wg.Wait()
}

// Воркеры завершаются при закрытии канала. Закрытие канала происходит при окончании работы и вызове функции Stop().
func (cp *ChannelPrinter) worker() {
	defer cp.wg.Done()
	for {
		select {
		case data, ok := <-cp.ch:
			if !ok {
				log.Printf("Worker stopped")
				return
			} else {
				fmt.Println(data)
			}
		}
	}
}
