# Go com TDD para Pythonistas

## Visão

Esse tutorial é especial por dois motivos: ganhamos tempo assumindo que você já sabe Python, e aprendemos a fazer testes logo no início. Comparando código Go e Python lado a lado vemos na prática semelhanças e diferenças entre elas. Praticando [TDD](http://tdd.caelum.com.br/) ganhamos confiança e uma ferramenta para explorar a linguagem Go.


## Descrição

A maioria dos tutoriais e textos introdutórios de Go deixam testes para o final ou apenas ignoram essse assunto. Porém, saber como testar código em Go é requisito essencial para programar profissionalmente. Aprender a fazer testes de exemplos também pode te ajudar a dominar Go mais rapidamente, permitindo que você teste facilmente suas hipóteses enquanto estuda a linguagem e suas bibliotecas.

Nessa introdução a Go com TDD (desenvolvimento orientado a testes), veremos técnicas para ter produtividade em uma abordagem de testes em primeiro lugar. Você usará o TDD para desenvolver um utilitário do zero: sinais, um programa de linha de comando que permite pesquisar caracteres Unicode por nome (uso diariamente para encontrar símbolos especiais e emojis).

Para ganhar tempo na apresentação da linguagem Go, vamos assumir que as pessoas participantes já sabem Python. Assim explicaremos rapidamente os conceitos de Go comparando e contrastando código Python e código Go, lado a lado.


## Ementa

- Estrutura de programas e pacotes
- Declaração de funções e variáveis
- Estruturas de controle de fluxo sequencial
- Tipos compostos embutidos
- Tipos estruturados
- Métodos e interfaces
- Programação concorrente
- ALMOÇO
- Coding Dojo para praticar TDD
- Testando código produz texto na saída padrão
- Testando a função principal
- Mocando argumentos de linha de comando em testes
- MVP1: Um executável minimamente usável
- Subtestes orientados por tabela
- Fajutando variáveis ​​de ambiente em testes
- Fingindo o acesso HTTP remoto
- Testando código que cria arquivos
- Tornando testes lentos opcionais
- Testando o produto acabado
- Testando um serviço HTTP




## Preparação

- [instale](https://golang.org/doc/install) a versão mais recente de Go
- configure seu editor para programar em Go

## Conteúdo

1. [Olá, 🌎!](01-hello)
2. [Estruturas de controle básicas](02-controle)
3. [Tipos nativos](03-nativos)

## Proposta didática

Esse tutorial é uma versão extendida que daquele que apresentei na [OSCON 2018](https://conferences.oreilly.com/oscon/oscon-or-2018/public/schedule/detail/67124). A principal crítica do tutorial na OSCON foi que "gastamos" muito tempo praticando TDD num Coding Dojo, em vez de apresentar mais conteúdo. Isso é intencional: acredito que cursos presenciais devem priorizar atividades interativas e mão na massa, em vez de aulas expositivas que podem ser vistas em vídeo. A melhor forma de conhecer TDD é praticando em um dojo ao vivo. A participação no dojo é voluntária, mas mesmo quem não participar vai aprender com as discussões e a dinâmica.


