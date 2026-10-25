import random

lista: list[int] = [1,2,3,4,5,6,7,8,9]

aleatorio = random.choice(lista)
print(aleatorio)

#embaralhar a lista
random.shuffle(lista)
print(lista)

palavra: str = "Aleatorio"
letra: str = random.choice(palavra)
print(letra)

#num aleatório de 1 a 10
numero_aleatorio = random.randint(1,10)
print(numero_aleatorio)

#num aleatório de 1 a 10, pulando de 2 em 2
numero_aleatorio = random.randrange(1, 10, 2)
print(numero_aleatorio)

#num aleatório de 1 a 10, com decimais
numero_decimal_aleatorio: float = random.uniform(1, 10)
print(numero_decimal_aleatorio)

#num aleatório de 0 a 1, com decimais
num_0_1: float = random.random()
print(num_0_1)
