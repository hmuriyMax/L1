package main

import (
	"fmt"
	"log"
	"strconv"
)

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}

// HEX Новая структура, для которой нужен адаптер в старую
type HEX struct {
	code string
}

// HEXAdapter Сам адаптер: удовлетворяет интерфейсу Color, но работает с HEX
type HEXAdapter struct {
	HEX
}

func (h *HEX) ConvertToRGBA() (r RGBA, err error) {
	numstr, err := ConvertInt(h.code[0:2], 16, 10)
	if err != nil {
		return RGBA{}, err
	}
	atoi, err := strconv.Atoi(numstr)
	if err != nil {
		return RGBA{}, err
	}
	r.red = atoi

	numstr, err = ConvertInt(h.code[2:4], 16, 10)
	if err != nil {
		return RGBA{}, err
	}
	atoi, err = strconv.Atoi(numstr)
	if err != nil {
		return RGBA{}, err
	}
	r.green = atoi

	numstr, err = ConvertInt(h.code[4:6], 16, 10)
	if err != nil {
		return RGBA{}, err
	}
	atoi, err = strconv.Atoi(numstr)
	if err != nil {
		return RGBA{}, err
	}
	r.blue = atoi

	numstr, err = ConvertInt(h.code[6:8], 16, 10)
	if err != nil {
		return RGBA{}, err
	}
	atoi, err = strconv.Atoi(numstr)
	if err != nil {
		return RGBA{}, err
	}
	r.alpha = atoi

	return
}

func (h *HEXAdapter) UseColor() {
	rgba, err := h.ConvertToRGBA()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Hey, i am color: %s, rgb(%d, %d, %d, %d)", h.code, rgba.red, rgba.green, rgba.blue, rgba.alpha)
}

func main() {
	col := HEX{code: "AAAAAAFF"}
	colA := HEXAdapter{col}
	Print(&colA)
}
