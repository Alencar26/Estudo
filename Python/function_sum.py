
# implementação nativa do python
numbers = [1, 2, 3, 4, 5, 6, 7]
print(sum(numbers))


#implementação customizada da mesma função acima
def custom_sum(interable):
    result = 0
    for item in interable:
        result += item
    return result

print(custom_sum(numbers))
