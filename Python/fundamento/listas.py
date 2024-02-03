#brincando com listas

lista = [True, "chicago", 2.5, False, 4, 8, "carro", 'e']

print(lista)
print(lista[::]) #usando lib numpy
print(lista[2:]) #imprimeindo do item 2 até o final
print(lista[:3]) #imprimindo do primeiro até o item 3
print(lista[1:4]) #do item 1 ao 4

print(lista[::2]) #exibindo de 2 em dois
print(lista[1:6:3]) #exibindo do 1 ao 6 e de 3 em três

palavra = "hidrometro"
print(palavra[:5]) # imprimindo somente o "hidro"
print(palavra[5:]) # imprimindo somente o "metro"
print(palavra[::2]) # imprimindo somente letras em posição par

#---

lista1 = [1,2,3,4,5]
lista2 = ['a',  'b', 'c', 'd', 'e']
lista3 = [True, False, False, True]

lista1 = lista2 + lista3
print(lista1)

print(len(lista1))
print(len(lista3))

print(max(lista2)) 
print(min(lista2))

#---

frase = "Corrida de Cavalos"
valor = range(10)

print(frase)
print(valor) #print -> range(0, 10)

lista4 = list(range(10))
print(lista4) # print -> [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

lista5 = list(frase)
print(lista5) #print -> ['C', 'o', 'r', 'r', 'i', 'd', 'a', ' ', 'd', 'e', ' ', 'C', 'a', 'v', 'a', 'l', 'o', 's']

#---

print(type(lista2)) #print -> <class 'list'>

print(lista2) # ['a', 'b', 'c', 'd', 'e']
lista2[2] = "carro"
print(lista2) # ['a', 'b', 'carro', 'd', 'e']
lista2[1:4] = ["Arroz", "Beterraba", "Cebola"]
print(lista2) # ['a', 'Arroz', 'Beterraba', 'Cebola', 'e']
lista2[1:2] = ["Bicicleta", "Bola", "Roupa", "Brinco"]
print(lista2) # ['a', 'Bicicleta', 'Bola', 'Roupa', 'Brinco', 'Beterraba', 'Cebola', 'e']
lista2[1:7] = ["Canoa"]
print(lista2) # ['a', 'Canoa', 'e']

#--- ADD

for x in lista1:
    print(x) #print -> item

for i in range(len(lista1)):
    print(i) #print -> index
    
for i in range(len(lista1)):
    print(i, lista1[i]) #print -> index + item
    
email = "meunome@email.com"
email = list(email.split('@'))

email.append(["provedor", "contato"])

email.insert(0, "Meu email")
email.insert(len(email) + 1, ["Nova lista", "Dentro da lista"])

email.extend(["Não gera", "outra lista", "inclui na mesma"])

for item in email:
    print(item)
    
#--- REMOVE

email.pop() # remova ultimo elemento
email.pop(3) # remova pelo index
email.remove("Meu email") # remove item específico
#del email # remove a lista
del email[2] # remove por index


for item in email:
    print(item)
    
email.clear() # remove all elements (sem gerar erro na lista)
print(email)

#--- ORDENAÇÃO

carrinho_compars = ["Mouse", "Teclado", "Placa de Vídeo", "Monitor", "Gabinete"]
carrinho_compars.sort() #ordenção 
print(carrinho_compars)

carrinho_compars.sort(reverse=True) #ordenção reverso
carrinho_compars.reverse() # mesma coisa da linha de cima
print(carrinho_compars)

nomes = ["Carlos", "João", "Ana", "Bruna", "marcos", "andré", "heloísa", "judas"]
nomes.sort() # ordena os nomes Maiusculos e Minusculos separadamente
print(nomes)

nomes.sort(key= str.lower) #ignorando letras maiusculas e minusculas
print(nomes)

#---

original = ['a', 'b', 'c']
copia = original

print(original)
print(copia)

original.append('e')

print(original)
print(copia) # a cópia repete as alterações do original

#--> SOLUÇÃO para cópia

original_new = ['z', 'w', 'x']
copia_new = original_new.copy() # add copy

print(original_new)
print(copia_new)

original_new.append('y')

print(original_new)
print(copia_new) # a cópia repete as alterações do original