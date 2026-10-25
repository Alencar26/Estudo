import random, string

def gerar_senha(palavra):
    for letra in palavra:
        if letra.lower() == "a":
            palavra = palavra.replace("a", '4')
        elif letra.lower() == "e": 
            palavra = palavra.replace("e", '3')
        elif letra.lower() == "i":
            palavra = palavra.replace("i", '1')
        elif letra.lower() == "o":
            palavra = palavra.replace("o", '0')
        elif letra.lower() == "u":
            palavra = palavra.replace("u", '5')
        else:
            alfabeto = ["b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"]
            m = []
            for i in alfabeto:
                m.append(i.upper())
            alfabeto.extend(m)
            palavra = palavra.replace(letra, alfabeto[random.randint(0, len(alfabeto)-1)])
    return palavra
            

print("GERADOR DE SENHAS SIMPLES")
senha = gerar_senha(input("Informe a palavra base para sua senha: "))
print(f"Sua senha é: {senha}")
