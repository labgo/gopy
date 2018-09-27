---
permalink: 02-sintaxe-basica
---

# Sintaxe básica


## Estrutura de um arquivo-fonte

Estas são as palavras reservadas que podem ser usadas no primeiro nível de um arquivo-fonte:

- `package`
- `import`
- `const`
- `type`
- `var`
- `func`

Todo arquivo-fonte em Go precisa declarar o pacote ao qual ele pertence, e esta declaração deve ser o primeiro código no arquivo. Ela pode ser precedida de um comentário de pacote (ver **Comentários**, a seguir).

As demais declarações podem ser usadas em qualquer ordem, uma ou mais vezes. No caso de `import`, a convenção é usar apenas um `import` com parêntesis para declarar múltiplas dependências.

## Comentários

Go aceita as sintaxes de comentários de C++:

- `//` comenta até o final da linha
- `/* */` comenta um bloco de código

O mais comum é usar `//`, mesmo em múltiplas linhas.

A sintaxe de bloco `/* */` é mais usada para omitir temporariamente um trecho, ou então para comentários muito longos no nível do pacote (no topo de um arquivo-fonte).

A ferramenta `godoc` gera documentação a partir de comentários escritos no topo de um arquivo-fonte, ou imediatamente antes de uma declaração de função, tipo, constante, ou variável global. Tais comentários não são incorporados ao programa compilado, como ocorre com as *docstrings* de Python.

> 📖 Veja [
Godoc: documenting Go code](https://blog.golang.org/godoc-documenting-go-code) para saber mais.

> 📖 Veja [Package testing→Overview→Examples](https://golang.org/pkg/testing/#hdr-Examples) para conhecer o uso de comentários em testes com exemplos, similar aos *doctests* de Python.


## Inserção automática de `;`

Como em Python, o sinal `;` raramente aparece em um arquivo-fonte. Ele só é necessário quando se quer escrever mais de uma instrução na mesma linha, ou nas variantes mais complexas de `if` e `for`.

Por baixo dos panos, o parser de Go insere `;` automaticamente ao final de cada instrução. Para que o algoritmo funcione, quando uma instrução ou declaração continua na próxima linha, ela deve deixar uma dica de que vai continuar.

Por isso o sinal `{` ou `(` deve ser sempre colocado no final da primeira linha de uma declaração de várias linhas.

Pelo mesmo motivo, expressões muito extensas devem ter quebras de linhas após operadores, e não antes. Assim:

```go
x +
y
```

E não assim:

```go
x
+ y
```

No primeiro caso, o parser sabe que a expressão precisa continuar. No segundo caso, não há como adivinhar a continuação.

> 📖 Veja [Effective Go→Semicolons](https://golang.org/doc/effective_go.html#semicolons) para saber mais.


## Estilo

Não existe o equivalente do PEP-8 no nível da sintaxe. Em vez disso o comando `go fmt` e o utilitário `gofmt` são **fortemente** recomendados para arrumar seu código antes de um commit. Uma IDE bem configurada deve rodar `gofmt` no arquivo-fonte toda vez que ele é salvo. Não discuta, apenas faça assim e seja feliz.

Questões de estilo de mais alto nível são tratadas no documento oficial [*Effective Go*](https://golang.org/doc/effective_go.html). Além de tratar brevemente de formatação, *Effective Go*  discute a programação idiomática de Go para produzir código mais fácil de entender, menos sujeito a erros, e com melhor desempenho.


## Identificadores

Qualquer caractere Unicode considerado uma *letra* pode ser usado para formar identificadores. Isso inclui letras acentuadas, letras gregas, ideogramas chineses etc. Além de letras, o `_` pode ser usado. Isoladamente, ele tem um significado especial em alguns contextos. Depois de uma ou mais letras, dígitos de 0 a 9 também podem ser usados.

> 📖 Veja [Effective Go→Names](https://golang.org/doc/effective_go.html#names) para saber mais.

### Caixa alta e baixa: Público e privado

Para identificadores com várias palavras, Go usa *camelCase*  e não *snake_case* como Python.

Se o identificador é o nome de uma declaração de primeiro nível (`type`, `func`, `var`, `const`), ou o nome de um campo em um `struct`, a primeira letra deve ser maiúscula para indicar que este nome é público e pode ser acessado via `import` em outros pacotes. Se a primeira letra for minúscula, o nome é privado e não pode ser acessado por outros pacotes.
