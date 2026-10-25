
#função de filtro
def bigger_that_3(n):
    return n > 3

#implementação nativa do python
filtered_unmbers = list(filter(bigger_that_3, [1, 4, 6, 2, 5]))

print(filtered_unmbers)


#implementação customizada da função acima
numbers = [1, 4, 6, 2, 5]
filtered_unmbers = []

for num in numbers:
    if bigger_that_3(num):
        filtered_unmbers.append(num)

print(filtered_unmbers)
