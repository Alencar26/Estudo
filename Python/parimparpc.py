import random

print("Olá humano, bora jogar par ou impar comigo?")

aposta = input("Par ou Impar?").lower()

if aposta == "par":
    print("Então eu sou Impar!")
else:
    print("Então eu sou Par!")


jogadaHumano = int(input("Digite um valor de 1 a 10: "))
jogadaPc =  random.randint(1,10)
print("Ok, meu lance é", jogadaPc)
print()

if (jogadaHumano + jogadaPc) % 2 == 0:
    if aposta == "par":
        print("Parabéns humano, você ganhou!")
    else:
        print("AHHHH! HA-HA-HA Eu ganhei, vou dominar o mundo!")
else:
    if aposta == "impar":
        print("Parabéns humano, você ganhou!")
    else:
        print("AHHHH! HA-HA-HA Eu ganhei, vou dominar o mundo!")