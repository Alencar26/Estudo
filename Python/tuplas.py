
# Tuplas

#--- Diferença entre lista e tupla
lista = [1, 3, 4, 6]
print(lista)
print(len(lista))
print(type(lista))

print(30*'-')

tupla = (1, 3, 4, 6)
print(tupla)
print(len(tupla))
print(type(tupla))

print(30*'-')

print(tupla[2])

tupla2 = ("item1", "item2", "item3", "item1", "item1")
print(tupla2.count("item1")) # quantas vezes o item paarece na tupla

""" -- tupllas são imutáveis
tupla2[0] = "novo valor" #tupla não aceita isso!
"""

#tuplas são utilizadas em situações com "listas" fechadas
#exemplos UF não precisa ser modificada, só escolher os valores desejados
UF = ("SP", "PR", "GO", "DF")

#--- Tupla com um valor

exemplo = ("exemplo") #errado
print(type(exemplo)) # <class 'str'>

um_valor = ("valor_unico",) #certo
print(type(um_valor)) # <class 'tuple'>

um_valor_new = "valor_unico", #certo <-- tbm é uma tupla
print(type(um_valor_new)) # <class 'tuple'>

tupla_nome = "nome1", "nome2", "nome3", "nome4" #certo <-- tbm é uma tupla
print(type(tupla_nome)) # <class 'tuple'>

#---- Tranformando tupla em lista - Add valor _ tranforma em tupla novamente

tupla_itens = ("item1", "item2", "item3", "item4")
print(tupla_itens)

lista_itens = list(tupla_itens)
print(lista_itens)

lista_itens.append("item5")
print(lista_itens)

tupla_itens = tuple(lista_itens)
print(tupla_itens)

#--- REMOVE

del tupla_itens

#--- CONCATENAÇÃO

tupla_a = ("itemA",)
tupla_b = ("itemB", "itemC")

print(tupla_a + tupla_b)
print(tupla_a * 3)

tupla_3x = tupla_a * 3

for i in range(len(tupla_3x)):
    print('-' * 7)
    print(i, tupla_3x[i])
    
# --- desempacotar as variáveis - o mesmo vale para listas

pacote = ("pão", "leite", "arroz", "pizza")

(x, y, z, a) = pacote
print(x)
print(y)
print(z)
print(a)

pao = pacote[0]
print(pao)

(x, *y) = pacote # x -> 1 item | y -> lista dos demais itens
print("x:", x)
print("y:", y)

(x, *y, z) = pacote # x -> 1º item | y -> itens entre 1º e ultimo | z -> ultimo item
print("x:", x)
print("y:", y)
print("z:", z)