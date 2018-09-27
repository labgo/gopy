---
permalink: 03-vars-simples
---

# Variáveis e tipos simples

Todos os tipos simples são tratados como valores, exceto `string`, que é uma referência. 


> 🔍 Strings em Go são imutáveis. Então o fato de uma variável `string` ser implementada como referência é um detalhe de implementação: uma otimização que não tem impacto na lógica do programa.

## Declaração explícita de variável

Variáveis podem ser declaradas com `var` no nível mais alto do pacote (variável global) ou dentro de funções e mesmo dentro de blocos de estrutura de controle *(block scope)*.

Alguns exemplos de uso de `var`:

```go
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
		fmt.Printf("var %c %-13T = %#[2]v\n", nomes[i], v)
	}
}
```

Exercício: complete o programa acima, faça ele rodar e observe a saída. Alguma surpresa?

## Declaração curta

A sintaxe de *short declaration* pode ser usada somente dentro de funções (não no nível global).

O tipo da variável é inferido (deduzido) em função do valor da expressão à direita.

```go
func main() {
	a := 1
	b, c := int16(2), 'C'
	d, e := os.Open("arquivo.x")
	f := 1.6
	g := nil
	nomes := "abcdefg"
	valores := []interface{}{a, b, c, d, e, f}
	for i, v := range valores {
		fmt.Printf("%c := %#v  // %[2]T\n", nomes[i], v)
	}
}
```

No exemplo acima, as variáveis `i` e `v` só existem dentro do bloco do laço `for`.

## Tipos simples

### Booleano

Tipo `bool`, não compatível com números inteiros.

Constantes: `true`, `false` (inicial minúscula). 

### Números inteiros

#### Inteiros com sinal

> Quando precisar de um valor inteiro deverá usar `int`, a menos que tenha um motivo específico para usar um tipo de inteiro com tamanho especificado ou sem sinal. *[Tour, Tipos Básicos](https://go-tour-br.appspot.com/basics/11)*


```go
int8   // inteiro de  8 bits (-128 a 127)
int16  // inteiro de 16 bits (-32768 a 32767)
int32  // inteiro de 32 bits (-2147483648 a 2147483647)
int64  // inteiro de 64 bits (-9223372036854775808 a 9223372036854775807)
int    // inteiro de 32 ou 64 bits, dependendo da arquitetura da máquina
rune   // sinônimo de int32, para código de caractere Unicode
```

`rune` não é um tipo distinto de `int32`, é apenas um *alias* (apelido).

>  Conversions are required when different numeric types are mixed in an expression or assignment. For instance,  `int32` and `int` are not the same type even though they may have the same size on a particular architecture. *[The Go Programming Language Specification, Numeric types](https://golang.org/ref/spec#Numeric_types)*


#### Inteiros sem sinal

> [...] números sem sinal tendem a ser usados somente quando operadores bit a bit ou operações aritméticas peculiares forem necessárioa, por exemplo, quando implementamos conjuntos de bits (bitsets), fazemos parse de formatos de arquivos binários ou para hashing e criptografia. Normalmente, eles não são usados em casos que sejam apenas de quantidades não negativas. *GOPL, 3.1 Inteiros*

```go
uint8   // inteiro de  8 bits (0 a 255)
uint16  // inteiro de 16 bits (0 a 65535)
uint32  // inteiro de 32 bits (0 a 4294967295)
uint64  // inteiro de 64 bits (0 a 18446744073709551615)
uint    // inteiro de 32 ou 64 bits, dependendo da arquitetura da máquina
uintptr // para armazenar os bits de um ponteiro
byte    // sinônimo de uint8
```

`byte` não é um tipo distinto de `uint8`, é apenas um *alias* (apelido).



### Números de ponto flutuante

```go
float32     número IEEE-754 de 32 bits
float64     número IEEE-754 de 64 bits

complex64   número complexo com partes real e imaginária float32
complex128  número complexo com partes real e imaginária float64
```

### Strings

Tipo `string` define uma sequência imutável de bytes normalmente usada para armazenar texto Unicode codificado como UTF-8. O tipo `string` de Go se parece com o tipo `str` de Python 2 (diferente do tipo `str` de Python 3):

- `len(s)` devolve o número de bytes, e não o número de caracteres. Para obter o número de caracteres, use `utf8.RuneCountInString(s)`.
- Para iterar caractere a caractere, use `for...range`.

```go
package main

import (
	"fmt"
	"utf8"
)

func main() {
	ss := []string{"bola", "café", "世界", "🌎"} 
	for _, s := range(ss) {
		fmt.Println("%s %d %d %t", s, len(s), utf8.RuneCountInString(s), []runne(s))
	}
}
```
