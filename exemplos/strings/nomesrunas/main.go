// Nomes exibe os nomes das runas que formam uma string
package main

import (
	"fmt"
	"unicode/utf8"

	"golang.org/x/text/unicode/runenames"
)

func main() {
	fmt.Printf("s\tlen(s)\tcount\trunes\n")
	fmt.Printf("────\t──────\t─────\t───────────────────────────────\n")

	ss := []string{"cafe", "café", "café", "世界", "🌎", "🖖🏿"}
	for _, s := range ss {
		runas := []rune(s)
		fmt.Printf("%s\t%5d\t%5d\t%#v\n", s, len(s),
			utf8.RuneCountInString(s), runas)
		for i, r := range runas {
			fmt.Printf("\t\t\t\t%d  U+%04X  %c\t%s\n", i, r, r,
				runenames.Name(r))
		}
	}
}
