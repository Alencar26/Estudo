expectJ1 = input("Jogador1: Par ou impar?").lower()
expectJ2 = input("Jogador2: Par ou impar?").lower()

j1 = int(input("Jogador1 digire um número:"))
j2 = int(input("Jogador2 digire um número:"))

if (j1 + j2) % 2 == 0:
    if expectJ1 == "par":
        print("PAR - Jogador 1 ganhou!")
    else:
        print("PAR - Jogador 2 ganhou!")
else:
    if expectJ1 == "impar":
        print("IMPAR - Jogador 1 ganhou!")
    else:
        print("IMPAR - Jogador 2 ganhou!")