// Elementos demonstra os elementos que formam uma string
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Printf("s\tlen(s)\tcount\trunes\n")
	fmt.Printf("────\t──────\t─────\t───────────────────────────────\n")

	ss := []string{"cafe", "café", "café", "世界", "🌎", "🖖🏿"}
	for _, s := range ss {
		fmt.Printf("%s\t%5d\t%5d\t%#v\n", s, len(s),
			utf8.RuneCountInString(s), []rune(s))
	}
}
