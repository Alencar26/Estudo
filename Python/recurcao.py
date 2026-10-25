

def potencia_recursiva(base, expoente):
    if (expoente == 0):
        return 1
    else:
        return base * potencia_recursiva(base, expoente - 1)

base = 6
expoente = 9
resultado_recursivo = potencia_recursiva(base, expoente)
print(resultado_recursivo)


def potencia(base, expoente):
    resultado = base
    if expoente == 0:
        return 1
    else:
        for i in range(1, expoente):
            resultado = resultado * base
    return resultado

print(potencia(6,9))