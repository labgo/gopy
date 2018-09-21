---
permalink: 01-hello
---

# Olá, 🌎!

## Primeiro exemplo

Todo arquivo-fonte em Go deve ser salvo na codificação UTF-8. Isso permite usar caracteres Unicode em qualquer lugar que a sintaxe permita: strings, comentários e identificadores.

```go
package main  // ➊

import "fmt"  // ➋

func main() {  // ➌
	fmt.Println("Olá, 🌎!")  // ➍
}
```

➊ Todo arquivo-fonte em Go precisa declarar o pacote ao qual ele pertence. 

➋ Importamos o pacote `fmt` da biblioteca padrão.

➌ Declaramos a função `main`, que será invocada quando este program for executado.

➍ Invocamos a função `fmt.Println`. Por padrão, qualquer nome definido em outro pacote deve ser prefixado com o nome do pacote. Por isso o prefixo `fmt` nesta linha.


## Nome do pacote

Para programas executáveis, o nome pacote deve ser `main`. Bibliotecas devem usar um nome igual à última parte do caminho até seu código-fonte. Por exemplo, os arquivos do pacote `strset`, disponível em "github.com/standupdev/strset", começam com esta declaração:

```go
package strset
```

## Importação

