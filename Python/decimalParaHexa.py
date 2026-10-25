# Recebe um número decimal do usuário
decimal = int(input("Digite um número decimal: "))

# Lista para armazenar os dígitos do número hexadecimal
hexadecimal = []

# Converte o número decimal em hexadecimal
while decimal > 0:
    resto = decimal % 16
    if resto < 10:
        hexadecimal.append(str(resto))
    else:
        hexadecimal.append(chr(resto - 10 + ord('A')))
    decimal //= 16

# Imprime o número hexadecimal
print("O número em hexadecimal é: 0x" + "".join(reversed(hexadecimal)))




# Na linha 2, o programa solicita ao usuário que digite um número decimal e o converte para inteiro
#  utilizando a função int().

# Na linha 5, é criada uma lista vazia hexadecimal que irá armazenar os dígitos do número hexadecimal.

# A partir da linha 8, o programa inicia um loop while que executa
#  enquanto o valor de decimal for maior do que zero. O objetivo é converter
#  o número decimal em hexadecimal utilizando o método de divisões sucessivas.

# Na linha 9, é calculado o resto da divisão de decimal por 16 e armazenado na variável resto.

# Na linha 10, o programa verifica se o valor de resto é menor do que 10.
#  Se for, adiciona o valor em string na lista hexadecimal. Caso contrário,
#  adiciona a letra correspondente na tabela ASCII (A, B, C, D, E ou F) na lista hexadecimal.

# Na linha 14, o programa divide o valor de decimal por 16 e utiliza o operador // 
# para realizar uma divisão inteira. Isso é necessário para que o loop while continue 
# até que não haja mais divisões a serem realizadas.

# Na linha 17, o programa imprime o resultado da conversão em hexadecimal. 
# O método reversed() é utilizado para inverter a ordem dos dígitos na lista
#  hexadecimal e join() é utilizado para concatenar os elementos da lista em uma
#  única string. O resultado final é concatenado com a string "0x" para indicar que
#  se trata de um número hexadecimal.