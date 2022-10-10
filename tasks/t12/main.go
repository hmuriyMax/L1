package main

import "fmt"

type Set struct {
	data   *map[interface{}]bool
	length int
}

func (s *Set) Init() {
	newMap := make(map[interface{}]bool)
	s.data = &newMap
	s.length = 0
}

func (s *Set) Add(v interface{}) {
	(*s.data)[v] = true
	s.length++
}

func (s *Set) Remove(v interface{}) {
	(*s.data)[v] = false
	s.length--
}

func (s *Set) Len() int {
	return s.length
}

func (s *Set) IsIn(v interface{}) bool {
	val, ok := (*s.data)[v]
	return ok && val
}

func (s *Set) String() (out string) {
	out += "{"
	for key, val := range *s.data {
		if val {
			out += fmt.Sprintf("%v, ", key)
		}
	}
	if s.length > 0 {
		out = out[:len(out)-2]
	}
	out += "}"
	return
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	set := Set{}
	set.Init()
	for _, word := range words {
		set.Add(word)
	}
	fmt.Println(set.String())
}
