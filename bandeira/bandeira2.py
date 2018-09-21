#!/usr/bin/env python3

# https://tapoueh.org/blog/2017/07/playing-with-unicode/

import sys

INDICADOR_A = 0x1F1E6  # REGIONAL INDICATOR SYMBOL LETTER A

def main():
    if len(sys.argv) != 2 or len(sys.argv[1]) != 2:
        print('Informe o código do país. Ex: BR ou br')
        sys.exit()

    codigos = []
    for c in sys.argv[1].upper():
        indicador = ord(c) - ord('A') + INDICADOR_A
        codigos.append(chr(indicador))

    # terminal do Ubuntu 18.04 ainda não reconhece REGIONAL INDICATORs...
    print(''.join(codigos))

if __name__ == '__main__':
    main()
