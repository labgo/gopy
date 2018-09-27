package main

import (
	"fmt"
	"os"
)

func main() {
	a := 1
	b, c := int16(2), 'C'
	d, e := os.Open("arquivo.x")
	f := 1.6
	nomes := "abcdefg"
	valores := []interface{}{a, b, c, d, e, f}
	for i, v := range valores {
		nome := nomes[i]
		fmt.Printf("%c := %#v  // %[2]T\n", nome, v)
	}
}
