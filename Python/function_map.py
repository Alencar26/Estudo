
#função de transformação (duplica valores)
def double(x):
    return x * 2

#função nativa
mapped_itens = list(map(double, [1,2,3]))
print(mapped_itens)


#implementação custom do map
mapped_itens = []
for item in [1,2,3]:
    mapped_itens.append(double(item))

print(mapped_itens)
