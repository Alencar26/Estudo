
lista = ['a', 'b', 'c', 'd', 'e']

def lista_para_dicionario(lista):
    dicionario = {}
    for i in range(len(lista)):
        dicionario[i] = lista[i]
    return dicionario

print(lista_para_dicionario(lista))