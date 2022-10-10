package main

import "fmt"

// Функция перебирает различные типы и выводит подходящий тип
func printType(val interface{}) {
	switch val.(type) {
	case string:
		fmt.Printf("%v is string\n", val)
	case int:
		fmt.Printf("%v is int\n", val)
	case bool:
		fmt.Printf("%v is bool\n", val)
	case chan any:
		fmt.Printf("%v is channel\n", val)
	default:
		// Можно просто вывести тип переменной, но сама программа тип не узнает
		fmt.Printf("%v is of unknown type \"%T\"\n", val, val)
	}
}

func main() {
	printType(4)
	printType("string")
	printType(true)
	ch := make(chan int)
	printType(ch)
}
