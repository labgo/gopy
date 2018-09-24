// duplicate counts all duplicate Unicode character names

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
		if len(name) == 0 {
			name = "(no name)"
		}
		counts[name]++
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
