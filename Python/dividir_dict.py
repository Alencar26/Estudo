


def divide_dict(dicionario):
    chaves = tuple(dicionario.keys())
    valores = tuple(dicionario.values())
    return chaves, valores

meu_dicionario = {"a": 1, "b": 2, "c": 3}
tupla_chaves, tupla_valores = divide_dict(meu_dicionario)

print(tupla_chaves)
print(tupla_valores)