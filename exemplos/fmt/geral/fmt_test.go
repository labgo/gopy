/* Exemplo de verbos de formatação gerais (funcionam com qualquer tipo)
   %v, %#v, %T

   Referência: https://golang.org/pkg/fmt/

   Nota: a ordem dos elementos de um map é indefinida, portanto testes de exemplo
   só funcionam de forma confiável com um mapa contendo apenas um elemento.
*/

package main

import "fmt"

func Example_general_verbs() {
	type Coordenada struct {
		lat, long float32
	}
	values := []interface{}{
		true, 1, 1.0, 'A', "abc",
		[]int{},
		[]int{1, 2, 3},
		[]string{"A", "B", "C"},
		[2]float64{10, 20},
		map[string]int{"um": 1}, // Exemplo só funciona com 1 elemento (*)
		Coordenada{23.55, 46.63},
	}
	for _, v := range values {
		fmt.Printf("%#v \t| %v \t| %T\n", v, v, v)
	}
	// Output:
	// true 	| true 	| bool
	// 1 	| 1 	| int
	// 1 	| 1 	| float64
	// 65 	| 65 	| int32
	// "abc" 	| abc 	| string
	// []int{} 	| [] 	| []int
	// []int{1, 2, 3} 	| [1 2 3] 	| []int
	// []string{"A", "B", "C"} 	| [A B C] 	| []string
	// [2]float64{10, 20} 	| [10 20] 	| [2]float64
	// map[string]int{"um":1} 	| map[um:1] 	| map[string]int
	// main.Coordenada{lat:23.55, long:46.63} 	| {23.55 46.63} 	| main.Coordenada
}

func Example_int_verbs_plain() {
	values := []interface{}{
		74, byte(75), 'L',
	}
	formato := "%T \t| %#v \t| %d \t| %o \t| %x \t| %X \t| %c \t| %q \t| %b \t| %U\n"
	fmt.Print(formato)
	for _, v := range values {
		fmt.Printf(formato, v, v, v, v, v, v, v, v, v, v)
	}
	// Output:
	// 	%T 	| %#v 	| %d 	| %o 	| %x 	| %X 	| %c 	| %q 	| %b 	| %U
	// int 	| 74 	| 74 	| 112 	| 4a 	| 4A 	| J 	| 'J' 	| 1001010 	| U+004A
	// uint8 	| 0x4b 	| 75 	| 113 	| 4b 	| 4B 	| K 	| 'K' 	| 1001011 	| U+004B
	// int32 	| 76 	| 76 	| 114 	| 4c 	| 4C 	| L 	| 'L' 	| 1001100 	| U+004C
}

func Example_int_verbs_alternate() {
	values := []interface{}{
		74, byte(75), 'L',
	}
	formato := "%T \t| %#v \t| %#o \t| %#q \t| %#x \t| %#X \t| %#U\n"
	fmt.Print(formato)
	for _, v := range values {
		fmt.Printf(formato, v, v, v, v, v, v, v)
	}
	// Output:
	// %T 	| %#v 	| %#o 	| %#q 	| %#x 	| %#X 	| %#U
	// int 	| 74 	| 0112 	| 'J' 	| 0x4a 	| 0X4A 	| U+004A 'J'
	// uint8 	| 0x4b 	| 0113 	| 'K' 	| 0x4b 	| 0X4B 	| U+004B 'K'
	// int32 	| 76 	| 0114 	| 'L' 	| 0x4c 	| 0X4C 	| U+004C 'L'
}
