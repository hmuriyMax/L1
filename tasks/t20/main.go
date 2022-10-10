package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {
	res := make([]byte, len(str))
	for i, v := range str {
		res[len(str)-1-i] = byte(v)
	}
	return string(res)
}

func main() {
	str := "hello world, my name is MaxInt64"
	fmt.Println(str)

	words := strings.Split(str, " ")
	for i, word := range words {
		words[i] = reverse(word)
	}
	str = strings.Join(words, " ")
	fmt.Println(str)
}
