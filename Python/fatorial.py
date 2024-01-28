
print(f"{15 * '-'} Fatorial {15 * '-'}")
number = int(input("Informe um número: "))
f = number
if number == 0: print("Não existe fatorial <= 0")
elif number == 1: print(f"Fatorial de {number} é: 1")
else:
    for n in range(number, 1, -1):
        f = f * (n -1)
    print(f"Fatorial de {number} é: {f}")
    