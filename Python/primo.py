
print(f"{15 * '-'} Primo {15 * '-'}")
number = int(input("Informe um númeor: "))

if number > 1:
    for n in range(2, number):
        if number % n == 0:
            print("Não Primo!")
            break
    else: print("Primo!")
else: print("Número <= 1")