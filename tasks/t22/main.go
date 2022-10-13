package main

import (
	"fmt"
	"math/big"
)

type LongInt struct {
	num string
}

func Sum(lhs, rhs LongInt) (res LongInt) {
	var sum, num1, num2 big.Int
	num1.SetString(lhs.num, 10)
	num2.SetString(rhs.num, 10)
	sum.Add(&num1, &num2)
	return LongInt{num: sum.Text(10)}
}

func Minus(lhs, rhs LongInt) (res LongInt) {
	var sum, num1, num2 big.Int
	num1.SetString(lhs.num, 10)
	num2.SetString(rhs.num, 10)
	sum.Sub(&num1, &num2)
	return LongInt{num: sum.Text(10)}
}

func Multiply(lhs, rhs LongInt) (res LongInt) {
	var mult, num1, num2 big.Int
	num1.SetString(lhs.num, 10)
	num2.SetString(rhs.num, 10)
	mult.Mul(&num1, &num2)
	return LongInt{num: mult.Text(10)}
}

func Divide(lhs, rhs LongInt) (res LongInt) {
	var div, num1, num2 big.Int
	num1.SetString(lhs.num, 10)
	num2.SetString(rhs.num, 10)
	div.Div(&num1, &num2)
	return LongInt{num: div.Text(10)}
}

func main() {
	a := LongInt{"1000000000000000000000000000"}
	b := LongInt{"2"}
	fmt.Println(Divide(a, b).num)
}
