import math

def area_base_circulo(raio):
    return math.pi * raio ** 2

def comprimento_base_circulo(raio):
    return 2 * math.pi * raio

def area_cilindro(altura, raio):
    return 2 * area_base_circulo(raio) + altura * comprimento_base_circulo(raio)

altura = float(input("Digite a altura do cilindro: "))
raio = float(input("Digite o raio do cilindro: "))

print(area_cilindro(altura, raio))  