pontos = int(input("""Bora jogar Jokenpô?
Quantos pontos para ganhar? resp: """))

j1_pontos = 0
j2_pontos = 0

j1_jogada = ""
j2_jogada = ""

while True:
    print("""Pedra...Papel...Tesoura...?
    [1] - Pedra
    [2] - Papel
    [3] - Tesoura""")

    j1_jogada = input("Jogador 1: ")
    j2_jogada = input("Jogador 2: ") 

    if j1_jogada == j2_jogada:
        print("empate")
    elif j1_jogada == "1" and j2_jogada == "3":
        j1_pontos += 1
        print("Jogador 1:", j1_pontos, "pontos.")
    elif j2_jogada == "1" and j1_jogada == "3":
        j2_pontos += 1
        print("Jogador 2:", j2_pontos, "pontos.")
    elif j1_jogada == "1" and j2_jogada == "2":
        j2_pontos += 1
        print("Jogador 2:", j2_pontos, "pontos.")
    elif j2_jogada == "1" and j1_jogada == "2":
        j1_pontos += 1
        print("Jogador 1:", j1_pontos, "pontos.")
    elif j1_jogada == "2" and j2_jogada == "3":
        j2_pontos += 1
        print("Jogador 2:", j2_pontos, "pontos.")
    elif j2_jogada == "2" and j1_jogada == "3":
        j1_pontos += 1
        print("Jogador 1:", j1_pontos, "pontos.")
    
    if j1_pontos == pontos:
        print("JOGADOR 1 GANHOU!")
        break

    if j2_pontos == pontos:
        print("JOGADOR 2 GANHOU!")
        break

