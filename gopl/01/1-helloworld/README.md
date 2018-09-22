# helloworld

O primeiro programa em [A Linguagem de Programação Go](https://novatec.com.br/livros/linguagem-de-programacao-go/) de Alan A. A. Donovan e Brian W. Kernighan.

## Notas

- Suporte nativo a Unicode. Código-fonte em UTF-8.
- Executação rápida: `go run hello.go`
- Compilar executável: `go build hello.go`
- Todo código pertence a algum pacote. Declaração `package` é obrigatória.
- Declaração `package main` define um executável (e não uma biblioteca).
- Compilador rejeita programas com importações faltando ou sobrando.
- Pacote `fmt` contém funções para formatar saídas e processar entradas. formatadas.
- Função `main`: "O que `main` faz é o que o programa faz."
- Estrutura geral do programa:
	- 1º `package …`
	- 2º `import …`
	- demais declarações em qualquer ordem:
		- `func`, `var`, `const`, `type`
- Parser (analisador sintático) faz inserção automática de `;`. Por isso:
	- `{` tem que estar no final linha que declara `func`
	- expressão `x + y` pode ter `\n` após o `+`, mas não antes
- Use `gofmt` para formatar seu código de forma padrão (como `black` em Python)
- Use `goimports` para organizar declarações `import`.

```bash
$ go get golang.org/x/tools/cmd/goimports
```

## Fonte

`gopl.io/ch1/helloworld`

Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.

License: https://creativecommons.org/licenses/by-nc-sa/4.0/

See page 1.

Helloworld is our first Go program.
