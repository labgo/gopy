# Emojis de bandeiras

Este diretório contém programas em Go e Python para exibir bandeiras usando os caracteres Unicode [REGIONAL INDICATOR SYMBOL](https://runefinder2018.appspot.com/?q=regional+indicator+symbol).

![programa bandeira no MacOS 10.13](bandeira.png "Saída do programa bandeira.go no terminal do MacOS 10.13")

## Como funciona

A norma Unicode inclui [algumas bandeiras genéricas](https://runefinder2018.appspot.com/?q=flag), mas não os emojis de bandeiras nacionais.

Porém, os caracteres especiais REGIONAL INDICATOR SYMBOL de A até Z (U+1F1E6...U+1F1FF) podem ser combinados em pares para formar códigos [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) de países ou regiões, produzindo emojis de bandeiras no software cliente.

Por exemplo, este par...

- U+1F1E7 — REGIONAL INDICATOR SYMBOL LETTER B
- U+1F1F7 — REGIONAL INDICATOR SYMBOL LETTER R

...produz o emoji da bandeira do Brasil 🇧🇷.

Pesquisando em setembro de 2018, não encontrei nenhum terminal no GNU/Linux ou no Windows que exiba as bandeiras. Mas no terminal do Mac OS X 10.13, funciona como na figura acima.

Se a combinação não corresponde a um código [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
, o terminal exibe um símbolo para cada letra.
