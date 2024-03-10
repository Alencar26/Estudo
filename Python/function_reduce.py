from functools import reduce
#função de transformaçaõ
def add(x,y):
    return x + y

#objetivo do reduce é reduzir sua lista a um único valor, não necessáriamente um soma

#função nativa
print(reduce(add, [1,2,3,4,5]))

#implementção custom

def custom_reduce(func, interable, initial=None):
    if initial is None:
        result = interable[0]
        interable = interable[1:]
    else:
        result = initial

    for item in interable:
        result = func(result, item)
    return result

numbers = [1,2,3,4,5]
total_sum = custom_reduce(add, numbers)

print(total_sum)
