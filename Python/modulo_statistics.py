import statistics

lista: list[int] = [13,42,42,34,75,42,96,27,18,9,10]

#sem usar a lib.
media: float = sum(lista) / len(lista)
print(media)

#usando a lib.
media2: float = statistics.mean(lista)
print(media2)

mediana: float = statistics.median(lista)
print(mediana)

moda: float = statistics.mode(lista)
print(moda)

#retorna uma lista com as modas (itens que mais aparecem)
letras: str = 'aaasfjfjjjjffsoootottrrrrrrppppppaasasss'
list_letras: list[str] = statistics.multimode(letras)
print(list_letras)