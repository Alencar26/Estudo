

print('-'*30)
n = int(input('Digite um número: '))
print('-'*30)
soma = 0
soma_primos = 0


for i in range(1, n + 1):
    for j in range(1, i + 1):
        if i % j == 0:
            soma += 1
    if soma == 2:
        soma_primos += i
    soma = 0

print('~'*30)
print(f"soma: {soma_primos}")
print('~'*30)


