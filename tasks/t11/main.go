package main

import "fmt"

// Set структура множества
type Set struct {
	data   *map[interface{}]bool
	length int
}

// Init инициализирует множество
func (s *Set) Init() {
	newMap := make(map[interface{}]bool)
	s.data = &newMap
	s.length = 0
}

// Add добавлет элемент в множество
func (s *Set) Add(v interface{}) {
	(*s.data)[v] = true
	s.length++
}

// Remove удаляет элемент из множества
func (s *Set) Remove(v interface{}) {
	(*s.data)[v] = false
	s.length--
}

// Len возвращает мощность множества
func (s *Set) Len() int {
	return s.length
}

// IsIn проверяет, входит ли элемент во множество
func (s *Set) IsIn(v interface{}) bool {
	val, ok := (*s.data)[v]
	return ok && val
}

// String возаращает текстовое представление множества
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

// Intersection получает множество-пересечение двух
func Intersection(s1, s2 Set) (res Set) {
	res.Init()
	if s2.Len() < s1.Len() {
		tmp := s1
		s1 = s2
		s2 = tmp
	}
	for key, val := range *s1.data {
		if val && s2.IsIn(key) {
			res.Add(key)
		}
	}
	return
}

func main() {
	var set1, set2 Set

	set1.Init()
	set2.Init()

	set1.Add(2)
	set1.Add(3)
	set2.Add(4)
	set2.Add(3)
	set1.Add("hi")
	set2.Add("hi")

	res := Intersection(set1, set2)
	fmt.Printf("set1: %v\nset2: %v\ninter: %v", set1.String(), set2.String(), res.String())
}
