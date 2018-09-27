package main

import (
	"fmt"
	"strings"
)

func main() {
	a := 3.14159265358979323846264338327950288
	b := float32(3.14159265358979323846264338327950288)
	c := (a - 2i)
	d := complex(b, -2*b)
	e := complex128(d)
	codigo := `a := 3.14159265358979323846264338327950288
	b := float32(3.14159265358979323846264338327950288)
	c := a - 2*a
	d := complex(b, -2*b)
	e := complex128(d)`
	valores := []interface{}{a, b, c, d, e}
	expressoes := strings.Split(codigo, "\n")
	for i, v := range valores {
		expr := strings.TrimSpace(expressoes[i])
		fmt.Printf("| %T | %#[1]v | %v |\n", v, expr)
	}
}
