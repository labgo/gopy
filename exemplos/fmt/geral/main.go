package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func generalVerbs() {
	type Cidade struct {
		nome      string
		lat, long float32
	}
	values := []interface{}{
		true, 1, 1.0, 'A', "abc", 3 - 4i,
		[]int{},
		[]int{1, 2, 3},
		[]string{"A", "B", "C"},
		[2]float64{10, 20},
		map[string]int{"um": 1, "dois": 2},
		Cidade{"São Paulo", 23.55, 46.63},
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)
	format := "%T \t│ %v \t│ %#v\n"
	fmt.Fprint(w, format)
	for _, v := range values {
		fmt.Fprintf(w, format, v, v, v)
	}
	w.Flush()
}

func intVerbs() {
	values := []interface{}{
		74, byte(75), 'L',
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)
	format := "%T \t│ %#v \t│ %d \t│ %o \t│ %x \t│ %X \t│ %c \t│ %q \t│ %b \t│ %U\n"
	fmt.Fprint(w, format)
	for _, v := range values {
		fmt.Fprintf(w, format, v, v, v, v, v, v, v, v, v, v)
	}
	w.Flush()
}

func intVerbsAlternate() {
	values := []interface{}{
		74, byte(75), 'L',
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)
	formato := "%T \t│ %#v \t│ %#o \t│ %#x \t│ %#X \t│ %#q \t│ %#U\n"
	fmt.Fprint(w, formato)
	for _, v := range values {
		fmt.Fprintf(w, formato, v, v, v, v, v, v, v)
	}
	w.Flush()
}

func title(s string) {
	fmt.Printf("\n%s%s\n", s, strings.Repeat("_", 79-len(s)))
}

func main() {
	title("generalVerbs")
	generalVerbs()
	title("intVerbs")
	intVerbs()
	title("intVerbsAlternate")
	intVerbsAlternate()
}
