nome_completo = input("nome: ")

primeiro_nome = ""
sobrenome = ""

#percorre a string e pega o primeiro nome
for caracter in nome_completo:
    if caracter == ' ':
        break
    primeiro_nome += caracter

#percorre a string do nome de trás para frente e add na variavel sobrenome
for i in range(len(nome_completo) - 1, -1, -1):
    if nome_completo[i] == " ":
        sobrenome = nome_completo[i+1:]
        break


print(primeiro_nome)
print(sobrenome)