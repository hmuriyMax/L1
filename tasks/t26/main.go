package main

import (
	"fmt"
	"strings"
)

func checkUnique(name string) bool {
	name = strings.ToLower(name)
	set := make(map[rune]struct{}, len(name))
	for _, ch := range name {
		_, ok := set[ch]
		if ok {
			return false
		}
		set[ch] = struct{}{}
	}
	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefA"
	str3 := "aabcd"
	fmt.Println(checkUnique(str1))
	fmt.Println(checkUnique(str2))
	fmt.Println(checkUnique(str3))
}
