---
permalink: 01-hello
---

# Olá, 🌎!

## Primeiro exemplo

Todo arquivo-fonte em Go deve ser salvo na codificação UTF-8. Isso permite usar caracteres Unicode em qualquer lugar que a sintaxe permita: strings, comentários e identificadores.

```go
package main

import "fmt"

func main() {
    fmt.Println("Olá, 🌎!")
}
```

## Como executar um programa em Go

### Go playground

Visite https://play.golang.org/. Copie e cole o código e use os botões na página.

### Go run

Vá até o diretório `…/gopy/exemplos/hello/` e execute o seguinte comando:

```bash
$ go run ola.go
Olá, 🌎!
```

Isso compila o programa em um diretório temporário, e o executa. Nenhum arquivo é gerado no diretório atual.

### Go build

O comando `go build` compila o programa e gera um binário executável. Por padrão, o executável tem o mesmo nome do diretório onde está o programa:

```bash
$ go build
$ ls
hello  ola.go
$ ./hello
Olá, 🌎!
```

Este executável não depende de mais nada para rodar *(stand-alone executable)*. Por padrão, as dependências externas são incorporadas ao binário *(statically linked binary)*.

### Use a sua IDE

Qualquer IDE corretamente configurada deve ser capaz de executar o programa e mostrar a saída em uma área de mensagens.

No VS Code, a tecla `[F5]` roda o programa, se o depurador [`delve`](https://github.com/derekparker/delve) estiver instalado.

## Olá, linha por linha

```go
package main  // ➊

import "fmt"  // ➋

func main() {  // ➌
	fmt.Println("Olá, 🌎!")  // ➍
}
```

➊ Declaração do nome do pacote.

➋ Importamos o pacote `fmt` da biblioteca padrão.

➌ Declaramos a função `main`, que será invocada quando este program for executado. A indentação se faz com um tab por nível. Não discuta.

➍ Invocamos a função `fmt.Println`. Nomes definidos em outros devem ser prefixado com o nome do pacote. Por isso o prefixo `fmt` nesta linha.


## Declaração `package`

Para programas executáveis, o nome pacote deve ser `main`. Bibliotecas devem usar um nome igual à última parte do caminho até seu código-fonte. Por exemplo, os arquivos do pacote `strset`, disponível em "github.com/standupdev/strset", começam com esta declaração:

```go
package strset
```

## Declaração `import`

É comum precisar importar mais de um pacote. A sintaxe convencional é colocar os nomes dos pacotes entre parêntesis, em linhas separadas:

```go
import (
    "encoding/json"
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"

    "github.com/standupdev/strset"

    "appengine"
    "appengine/urlfetch"
)
```

> 🔍 Existe uma sintaxe chamada *dot import* que corresponde ao `from pacote import *` em Python, mas ela é usada raramente, principalmente para testes e DSLs. Detalhes: [especificação](https://golang.org/ref/spec#Import_declarations) e [discussão](https://github.com/golang/lint/issues/179).


### Convenções

- Declare os pacotes em ordem alfabética.
- Se o pacote está na Web, coloque o caminho até ele, incluindo o domínio mas não o protocolo (`https://`)
- Opcionalmente, divida os pacotes em grupos separados por uma linha: biblioteca padrão, meus pacotes, etc. Dentro de cada grupo, respeite a ordem alfabética.

> 🔍 As ferramentas `go fmt` e `goimports` ajudam a organizar os imports, mas elas têm comportamentos distintos, como explicado nesse [issue](https://github.com/golang/go/issues/8963).
