package main

import "fmt"

func main() {
	str := "hello world"
	fmt.Println(str)
	res := make([]byte, len(str))
	for i, v := range str {
		res[len(str)-1-i] = byte(v)
	}
	str = string(res)
	fmt.Println(str)
}
