liquidos_ingeridos = {}

def adicionar_liquido(liquido, quantidade):
    if liquido in liquidos_ingeridos:
        liquidos_ingeridos[liquido] += quantidade
    else:
        liquidos_ingeridos[liquido] = quantidade

def obet_liquido(liquido):
    if liquido in liquidos_ingeridos:
        return liquidos_ingeridos[liquido]

while True:
    opcao = int(input("1 - Adicionar Liquido\n2 - Obter Liquido\n3 - Sair\nR: "))
    if (opcao == 1):
        liquido = input("Digite o qual é o líquido ingerido: ").lower()
        quantidade = float(input("Digite a quantidade ingerida em L: "))
        adicionar_liquido(liquido, quantidade)
    elif (opcao == 2):
        liquido = input("Digite o qual é o líquido ingerido: ").lower()
        print(f"{liquido}: {obet_liquido(liquido)} L")
    else:
        print("\nRegistro de líquidos ingeridos no dia:")
        for liquido, quantidade in liquidos_ingeridos.items():
            print(f"{liquido}: {quantidade} L")
        break