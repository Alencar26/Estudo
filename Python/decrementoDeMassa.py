massaInicial = float(input("Olá, informe a massa inicial do seu isótopo radioativo: "))

tempo = 0

while massaInicial >= 0.05:
    print("Sua massa atualmente tem a massa:", massaInicial, "gramas")
    print("Tempo atual é de:", tempo,"segundos")
    massaInicial = massaInicial / 2
    tempo = tempo + 50

print("Sua massa está em:", massaInicial)
print("Duração de:", tempo,"segundos")   