// similar counts all Unicode character names that are similar
// (same words except the last)
package main

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/text/unicode/runenames"
)

const lastCodePoint = 0x10FFFF // http://unicode.org/faq/utf_bom.html

func main() {
	fmt.Println("Unicode version", runenames.UnicodeVersion)
	var r rune
	counts := make(map[string]int)
	for r = 0; r <= lastCodePoint; r++ {
		name := runenames.Name(r)
		if len(name) == 0 || name[0] == '<' {
			continue // ignore nameless and named ranges
		}
		words := strings.Fields(name)
		if len(words) == 1 {
			continue // ignore single-word names
		}
		prefix := strings.Join(words[:len(words)-1], " ")
		counts[prefix]++
	}
	report := []string{}
	for line, n := range counts {
		if n > 1 {
			report = append(report, fmt.Sprintf("%6d\t%s", n, line))
		}
	}
	// sort.Strings(report)  // ascending order
	sort.Sort(sort.Reverse(sort.StringSlice(report))) // descending order
	fmt.Println(strings.Join(report, "\n"))
}
