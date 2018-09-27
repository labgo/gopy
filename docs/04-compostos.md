---
permalink: 04-compostos
---

# Tipos compostos

## Coleções embutidas

Coleções armazenam itens de um tipo pré-determinado, e permitem o acesso a cada item individualmente.

Go tem apenas 4 tipos de coleção embutidas:

| tipo | descrição | `len(x)` | `cap(x)` | pode ser `nil` | inicializado por `make` |
| ---- | --------- | -------- | -------- | -------------- | ----------------------- |
| `array`   | sequência pré-alocada de valores | fixo | fixo | não | não |
| `slice`   | janela sobre um `array` interno | variável | fixo | sim | sim |
| `map`     | dicionário *(hash table)* | variável | N/A | sim | sim |
| `channel` | fila FIFO manipulada com operador `<-` | variável | fixo | sim | sim |

> 🔍 Arrays são tratados como valores: todo o seu conteúdo é copiado na atribuição e na passagem como argumento de função, não importa se o tamanho do `array` é 1 byte ou 1 terabyte. Variáveis dos tipos `slice`, `map` e `channel` se comportam como referências: o conteúdo é compartilhado. Por esse motivo, é comum ver ponteiros para arrays, mas é raro ver ponteiros para `slice`, `map` ou `channel`.

> 📖 Leitura essencial sobre slices: [Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals).

## Exemplos

### Array

> 📖 Fonte: https://gobyexample.com/arrays

```go
package main

import "fmt"

func main() {

    var a [5]int
    fmt.Println("emp:", a)

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
```

### Slice

> 📖 Fonte: https://gobyexample.com/slices

```go

package main

import "fmt"

func main() {

    s := make([]string, 3)
    fmt.Println("emp:", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
```

### Map

> 📖 Fonte: https://gobyexample.com/maps

```go
package main

import "fmt"

func main() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}

```

### Channel

> 📖 Fonte: https://gobyexample.com/channels

```go
package main

import "fmt"

func main() {

    messages := make(chan string)

    go func() { messages <- "ping" }()

    msg := <-messages
    fmt.Println(msg)
}
```