numero_binario = input("Digite um número em binário: ")

tamanho = len(numero_binario)
potencia = tamanho - 1
numero_decimal = 0

for i in range(tamanho):
    numero_decimal += int(numero_binario[i]) * 2**potencia
    potencia -= 1

print("O número em decimal é:", numero_decimal)


def convert_hexa_to_deciaml(hexa):
    decimal = 0
    for i in range(len(hexa)):
        if hexa[i] == 'A':
            decimal += 10
        elif hexa[i] == 'B':
            decimal += 11
        elif hexa[i] == 'C':
            decimal += 12
        elif hexa[i] == 'D':
            decimal += 13
        elif hexa[i] == 'E':
            decimal += 14
        elif hexa[i] == 'F':
            decimal += 15
        else:
            decimal += int(hexa[i])
    return decimal

def convert_decimal_to_hexa(decimal):
    hexa = ''
    while decimal != 0:
        if decimal % 16 == 10:
            hexa += 'A'
        elif decimal % 16 == 11:
            hexa += 'B'
        elif decimal % 16 == 12:
            hexa += 'C'
        elif decimal % 16 == 13:
            hexa += 'D'
        elif decimal % 16 == 14:
            hexa += 'E'
        elif decimal % 16 == 15:
            hexa += 'F'
        else:
            hexa += str(decimal % 16)
        decimal = decimal // 16
    return hexa