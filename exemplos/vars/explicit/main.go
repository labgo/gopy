package main

import (
	"fmt"
	"os"
)

func main() {
	var a int16
	var b, c = "B", 'C'
	var d, e = os.Open("arquivo.x")
	var f float32 = 1.6180339887498948482
	var nomes = "abcdef"
	var valores = []interface{}{a, b, c, d, e, f}
	var (
		i int
		v interface{}
	)
	for i, v = range valores {
		var nome = nomes[i]
		fmt.Printf("var %c %-13T = %#[2]v\n", nome, v)
	}
}
