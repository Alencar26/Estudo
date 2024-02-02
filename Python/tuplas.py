
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

