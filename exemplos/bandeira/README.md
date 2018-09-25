# bandeira

Programas para exibir bandeiras usando os caracteres Unicode da faixa de [REGIONAL INDICATOR SYMBOL LETTER](https://runefinder2018.appspot.com/?q=regional+indicator+symbol+letter) de A até Z (U+1F1E6 até U+1F1FF).

Alguns terminais, navegadores e aplicativos de celular substituem um par de caracteres nesta faixa pela bandeira do país ou região correspondende. Por exemplo, este par...

- U+1F1E7: REGIONAL INDICATOR SYMBOL LETTER B
- U+1F1F7: REGIONAL INDICATOR SYMBOL LETTER R

...produz a bandeira do Brasil.

Isso funciona porque BR é o código do Brasil na norma [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).

Pesquisando em setembro de 2018, não encontrei nenhum terminal no GNU/Linux ou no Windows que exiba as bandeiras. Mas no terminal do Mac OS X 10.13, funciona assim:

![programa bandeira no MacOS 10.13](bandeira.png "Saída do programa bandeira.go no terminal do MacOS 10.13")

Se a combinação não corresponde a um código [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
, o terminal exibe um símbolo para cada letra.