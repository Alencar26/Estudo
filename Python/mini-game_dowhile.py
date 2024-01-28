"""
Mini game improvisando Do While

'Advinhar um número'    
"""

palpite = 7
numero = 9

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