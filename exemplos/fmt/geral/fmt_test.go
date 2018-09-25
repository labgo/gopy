/* Exemplo de verbos de formatação gerais (funcionam com qualquer tipo)
   %v, %#v, %T

   Nota: a ordem dos elementos de um map é indefinida, portanto esse teste
   só funciona de forma confiável com um mapa contendo apenas um elemento.
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
		[2]float64{10, 20},
		map[string]int{"um": 1}, // Exemplo só funciona com 1 elemento (*)
		Coordenada{23.55, 46.63},
	}
	for _, v := range values {
		fmt.Printf("| %v \t| %#v \t| %T\n", v, v, v)
	}
	// Output:
	// | true 	| true 	| bool
	// | 1 	| 1 	| int
	// | 1 	| 1 	| float64
	// | 65 	| 65 	| int32
	// | abc 	| "abc" 	| string
	// | [] 	| []int{} 	| []int
	// | [1 2 3] 	| []int{1, 2, 3} 	| []int
	// | [10 20] 	| [2]float64{10, 20} 	| [2]float64
	// | map[um:1] 	| map[string]int{"um":1} 	| map[string]int
	// | {23.55 46.63} 	| main.Coordenada{lat:23.55, long:46.63} 	| main.Coordenada
}
