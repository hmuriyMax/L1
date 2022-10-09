package main

import (
	"log"
	"sync"
)

type IntConcurrentMap struct {
	m map[int]int
	sync.Mutex
}

func (m *IntConcurrentMap) Init() {
	m.m = make(map[int]int)
}

func (m *IntConcurrentMap) Get(key int) (int, bool) {
	val, ok := m.m[key]
	return val, ok
}

// Set Используем мьютекс, чтобы избежать гонки данных
func (m *IntConcurrentMap) Set(key, value int) {
	m.Lock()
	m.m[key] = value
	m.Unlock()
}

func main() {
	log.Default().SetFlags(log.Ldate | log.Lmicroseconds)
	const N = 1000
	var cMap IntConcurrentMap
	cMap.Init()
	for i := 0; i < N; i++ {
		go cMap.Set(i, i*i)
	}
	for i := 0; i < N; i++ {
		val, ok := cMap.Get(i)
		if !ok {
			log.Fatalf("could not set value for key %v\n", i)
		}
		log.Printf("map[%d] = %d", i, val)
	}
}
