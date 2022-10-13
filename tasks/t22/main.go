package main

import (
	"fmt"
	"strings"
)

type LongInt struct {
	num string
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

var (
	numToRune = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
)

func runeToNum(r rune) int {
	for i, el := range numToRune {
		if el == r {
			return i
		}
	}
	return -1
}

func tinySum(lhs, rhs rune) (this, promise rune) {
	lInt := runeToNum(lhs)
	rInt := runeToNum(rhs)
	tInt := lInt + rInt
	this = numToRune[tInt%10]
	promise = numToRune[tInt/10]
	return
}

func Sum(lhs, rhs *LongInt) (res *LongInt) {
	//var maxLen int
	if len(lhs.num) > len(rhs.num) {
		delta := len(lhs.num) - len(rhs.num)
		rhs.num = strings.Repeat("0", delta) + rhs.num
	} else {
		delta := len(rhs.num) - len(lhs.num)
		lhs.num = strings.Repeat("0", delta) + lhs.num
	}

	num1 := Reverse(lhs.num)
	num2 := Reverse(rhs.num)

	var this, prom, nprom rune
	for i := 0; i < len(num1); i++ {
		this, prom = tinySum(rune(num1[i]), numToRune[prom])
		this, nprom = tinySum(this, rune(num2[i]))
		prom += nprom
		res.num += fmt.Sprint(this)
	}
	if prom != '0' {
		res.num += fmt.Sprint(this)
	}
	return
}

func Multiply(lhs, rhs *LongInt) (res *LongInt) {

	return
}

func main() {

}
