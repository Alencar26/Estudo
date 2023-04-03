print("===========VALIDAÇÃO DO CPF=============")

digito1 = False
digito2 = False

print("Informe seu cpf com os dígitos separadamente:")

d1 = int(input("Digite o dígito 1: "))
d2 = int(input("Digite o dígito 2: "))
d3 = int(input("Digite o dígito 3: "))
d4 = int(input("Digite o dígito 4: "))
d5 = int(input("Digite o dígito 5: "))
d6 = int(input("Digite o dígito 6: "))
d7 = int(input("Digite o dígito 7: "))
d8 = int(input("Digite o dígito 8: "))
d9 = int(input("Digite o dígito 9: "))
d10 = int(input("Digite o dígito 10: "))
d11 = int(input("Digite o dígito 11 "))

#cpf com máscara aplicada
#cpf = ""+d1+d2+d3+"."+d4+d5+d6+"."+d7+d8+d9+"-"+d10+d11
#print("O CPF informado foi:",cpf)

#validação do primeiro dígito
x = d1 * 10 + d2 * 9 + d3 * 8 + d4 * 7 + d5 * 6 + d6 * 5 + d7 * 4 + d8 * 3 + d9 * 2
print(x)
#validação do resto da divisão
x2 = (x * 10) % 11

if x2 == 10:
    if d10 == 0:
        digito1 = True
else: 
    if x2 == d10:
        digito1 = True

#validação segundo dígito
y = d1 * 11 + d2 * 10 + d3 * 9 + d4 * 8 + d5 * 7 + d6 * 6 + d7 * 5 + d8 * 4 + d9 * 3 + d10 * 2

#validação do resto da divisão
y2 = (y * 10) % 11

if y2 == 10:
    if d11 == 0:
        digito2 = True
else: 
    if y2 == d11:
        digito2 = True

#resultado
if digito1 and digito2:
    print("Seu CPF é válido!")
else:
    print("Seu CPF é inválido!")