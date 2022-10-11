package main

import (
	"log"
	"time"
)

func Sleep(t time.Duration) {
	start := time.Now()
	finish := start.Add(t)
	for finish.After(time.Now()) {
		continue
	}
}

func main() {
	log.Default().SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println("start")
	go func() {
		Sleep(1 * time.Second)
		log.Println("finished my sleep")
	}()
	time.Sleep(1 * time.Second)
	log.Println("finished standard sleep")
}
