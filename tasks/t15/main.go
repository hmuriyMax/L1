package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Глобальные переменные - не очень хорошо, здесь можно не использовать её
//var justString string

func createHugeString(len int) (s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		s += string(rune(rand.Intn(1 << 16)))
	}
	return strings.ToValidUTF8(s, "[incorrect]")
}

func someFunc() {
	// Перенесли сюда
	var justString string
	v := createHugeString(1 << 10)
	justString = v[:100]
	fmt.Printf(justString)
}

func main() {
	someFunc()
}
