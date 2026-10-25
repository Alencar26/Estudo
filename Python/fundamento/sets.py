"""
Sets: Coleção não ordenada, imutável e não permite valores duplicados
São conhecidos como conjuntos
"""

conjunto = {"item1", 2.4, True, 5}
print(conjunto)
print(type(conjunto))

#--- CONVERSÃO

tupla = (2,3,5,6)
conjunto_numeros = set(tupla)
print(conjunto_numeros)

#--- acessar itens

#não ordenado - na impessão não tem como garantir a ordem os elementos
conj = {"ITEM1", "ITEM2", "ITEM3"}
for item in conj:
    print(item)
    
#---

novo_conjunto = {1,2,3,4,5,6}
print(4 in novo_conjunto)

#--- ADD

print(conj)
conj.add("ITEM4")
conj.add(False)
conj.add(5)
print(conj)

#--- UPDATE

set1 = {1,2,3,4,5}
set2 = {"a", "b", "c", "d"}
set1.update(set2)
print(set1)

#--- REMOVE

set1.pop() #remove itens aleatórios
print(set1)

set1.remove("c") #remove item específico
print(set1)

set1.discard(True) #Não dá erro ao informar valor que não existe no set
set1.discard(3)
print(set1)

#--- USANDO FUNÇÕES IGUAL NA MATEMÁTICA

print("UNIÃO DOS CONJUNTOS")
c1 = {1,5,8,9}
c2 = {3,6,2,1}

c3 = c1.union(c2)
print(c3)

print("INTERCESSÃO DOS CONJUNTOS")
c1 = {1,5,3,9}
c2 = {3,6,2,1}

c3 = c1.intersection(c2) #exibe um conjunto com os valores repedidos
print(c3)

print("DIFERENÇA SIMÉTRICA DOS CONJUNTOS")
c3 = c1.symmetric_difference(c2) #conjunto com valores que não se repete
print(c3)

