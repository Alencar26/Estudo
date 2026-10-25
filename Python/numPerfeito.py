num = int(input("numero"))
somas = 0

for i in range(1, num):
    if num % i == 0:
        somas += i

if somas == num:
    print(num, "é um numero perfeito")
