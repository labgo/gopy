---
permalink: 04-compostos
---

# Tipos compostos

## Coleções embutidas

Coleções armazenam itens de um tipo pré-determinado, e permitem o acesso a cada item individualmente.

Go tem apenas 4 tipos de coleção embutidas:

| tipo | descrição | `len(x)` | `cap(x)` | pode ser `nil` | inicializado por `make` |
| ---- | --------- | -------- | -------- | -------------- | ----------------- |
| `array`   | sequência pré-alocada de valores | fixo | fixo | não | não |
| `slice`   | janela sobre um `array` interno | variável | fixo | sim | sim |
| `map`     | dicionário *(hash table)* | variável | N/A | sim | sim |
| `channel` | fila FIFO manipulada com operador `<-` | variável | fixo | sim | sim |

> 🔍 Arrays são tratados como valores: todo o seu conteúdo é copiado na atribuição e na passagem como argumento de função, não importa se o tamanho do `array` é 1 byte ou 1 terabyte. Variáveis dos tipos `slice`, `map` e `channel` se comportam como referências: o conteúdo é compartilhado. Por esse motivo, é comum ver ponteiros para arrays, mas é raro ver ponteiros para `slice`, `map` ou `channel`.

> 📖 Leitura essencial sobre slices: [Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals).
