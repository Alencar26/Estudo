

termos = int(input('Quantos termos de fibonacci você quer mostrar? '))

t1 = 0
t2 = 1

contador = 0

while contador < termos:
    print(t1)
    t3 = t1 + t2
    t1 = t2
    t2 = t3
    contador += 1