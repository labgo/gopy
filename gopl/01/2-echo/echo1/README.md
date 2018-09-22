# echo1

Echo1 exibe os argumentos da linha de comando.

## Notas

- Comentários marcados com `//` ou `/* … */`.
- Convenção: comentário descrevendo o pacote acima da declaração `package`.
- Pacote `os`: funções, variáveis etc. para lidar com o SO. Como `sys` e `os` em Python.
- Declaração `var`: identificadores, depois tipo.
- Inicialização automática com **valor zero** *(zero value)*.
- Valor zero do tipo `string` é `""` (uma string com `len(s) == 0`)  
- Variável `os.Args` é uma fatia de strings *(slice of strings)*: um array dinâmico de strings. Superficialmente é parecida com uma `list` em Python.
- Operações básicas com slices: `len(s)`, `s[0]`, `s[2:5]`
- `os.Args[0]`: nome do programa.
- Forma clássica do comando `for`, como em C:

```go
for umaVez; condição; sempreApós {
  // zero ou mais instruções
}
```

- `:=` declaração curta de variável (ver `echo2`). O tipo de `i` é deduzido *(type inference)*.
- Variável `i` só existe dentro do bloco `for` *(block scope)*.
- `i++` é uma instrução (não uma expressão). Também existe `i--`, mas não `++i` ou `--i`.
- Operador `+` concatena strings.

- Variante *while*:

```go
for condição {
  // ...
}
```

- Laço infinito (ou até `break`, `return`]):

```go
for {
  // ...
}
```

## Fonte

`gopl.io/ch1/echo1`

Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
License: https://creativecommons.org/licenses/by-nc-sa/4.0/

See page 4.

Echo1 prints its command-line arguments.
