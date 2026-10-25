
"""
Operadores que eu não conheço muito bem:

Exponencial **
Divisão por inteiro //

"""

print("Operadores:")
print(2**6)
print(5//3) #arredonda o número, tirando as casa decimais

nome = "André"
funcao = "QA"

print(f"Olá, meu nome é {nome} e minha função é de {funcao}")
print("Olá, meu nome é {0} e minha função é de {1}".format(nome, funcao))

"""
Operadores de associação 

in - x in y
not in - x not in y

"""

x = "paralelepipido"
y = "para"
w = "pipi"
z = "lolo"

print(y in x)
print(w in x)
print(z in x)
print(z not in x)

####################################################

# Exercícos

#01 - Calcular a área de um retângulo

# área = base * altura

print("Calcular a área de um retângulo")
base = float(input("Informe a medida da Base do retângulo: "))
altura = float(input("Informe a medida da Altura do retângulo: "))
print(f"A área do seu retângulo é de: {base * altura}")


print("\n calcular produto com desconto")
valor_intergal = input("Informe o valor do seu produto: ")
desconto = input("Informe a porcentagem do desconto a ser aplicado: ")

produto_com_desconto = float(valor_intergal) - (float(valor_intergal) * (float(desconto) / 100))
print(f"Produto com desconto: R${produto_com_desconto:.2f}")


################################################################

f = 1

if f != 1: print("diferente")

print("Igual") if f == 3 else print("diferente")

#################################################################3

"""
cond = None
while cond != "Sair".lower():
    cond = input("Informe o valor: ").lower()
    print("Igual a none") if cond is None else print(f"Seu valor é {cond}")
else: 
    print("Voce saiu do while")
"""
########################################################

for x in range(6):
    print(x)

for x in range(3,10):
    print(x)

for x in range(2, 20, 5):
    print(x)

for x in "abacate":
    print(x)
else: 
    print("fim do loop")


##################################################################
    
