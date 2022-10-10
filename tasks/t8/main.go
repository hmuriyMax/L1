package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// ConvertInt Функция, используемая для конвертации вывода в нужную базу
func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}

func main() {
	var (
		num int64
		i   int
		bit bool
	)
	rand.Seed(time.Now().UnixNano())
	num = rand.Int63()
	i = 2      // позиция слева
	bit = true // true - 1, false - 0
	fmt.Println(ConvertInt(fmt.Sprint(num), 10, 2))

	if bit {
		// Если нужно установить 1, можно использовать маску 00..010..00, где 1 стоит на нужной позиции
		mask := int64(1) << (64 - i)
		// Применяем к числу вышеуказанную маску с использованием побитового или. На нужной позиции гарантированно окажется 1
		num |= mask
	} else {
		// Если нужно установить 0, используем маску 111...101...111, где 0 стоит на нужной позиции
		minInt64 := int64(math.MaxInt64)
		// К максимальному числу (1111..111) применяем маску 00..010...00 с xor. Тогда на нудной позиции окажется 0
		mask := int64(1) << (64 - i)
		mask ^= minInt64
		// Примпеняем маску с логическим и. 0 гарантированно окажется на нужной позиции
		num &= mask
	}
	fmt.Println(ConvertInt(fmt.Sprint(num), 10, 2))
}
