package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	int
	sync.Mutex
}

// Tick В функции инкриминирования используем мьютекс, чтобы все посчиталось корректно
func (c *Counter) Tick() {
	c.Lock()
	c.int++
	c.Unlock()
}

func (c *Counter) Get() int {
	return c.int
}

func main() {
	t := Counter{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			t.Tick()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(t.Get())
}
