"""
Mini game improvisando Do While

'Advinhar um número'    
"""
from random import randint

palpite = 7
numero = randint(1, 10)

while bool(palpite) is True:
    print("Qual o número correto?")
    palpite = int(input("Seu palpite: "))
    if palpite == numero:
        print("Parabéns você acertou!!!")
        break
    else:
        print("Que pena, você errou.")
#--------------------------------------------------------

while True:
    print(30 * "-")
    print("Qual o número correto?")
    palpite = int(input("Seu palpite: "))
    if palpite == numero:
        print("Parabéns você acertou!!!")
        break
    else:
        print("Que pena, você errou.")

#---------------------------------------------------------

def recuusivo_game():
    print("Qual o número correto?")
    palpite = int(input("Seu palpite: "))
    if palpite == numero:
        print("Parabéns você acertou!!!")
    else:
        print("Que pena, você errou.")
        recuusivo_game()
        
recuusivo_game()
