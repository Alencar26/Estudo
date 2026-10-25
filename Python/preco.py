produto = 364.99

print("Olá cidadão de bem, você acaba de adiquirir um microondas no valor de R$",produto)
print()

def func(produto):
    print("Formas de pagamento:")
    print("1 - A vista no drinheiro (10% desconto)")
    print("2 - A vista no cartão (5% desconto)")
    print("3 - 2x sem juros")
    print("4 - 2x com juros de 10%")

    escolha = input("Escolha uma opção: ")

    if escolha == "1":
        produto -= produto * 0.10
        print("Você vai pagar: R$" ,produto)
    elif escolha == "2":
        produto -= produto * 0.05
        print("Você vai pagar: R$" ,produto)
    elif escolha == "3":
        duasX = format(produto / 2, '.2f');
        print("Você vai pagar 2x de: R$" ,duasX)
    elif escolha == "4":
        duasXJuros = format((produto + (produto * 0.10)) / 2, '.2f')
        print("Você vai pagar 2x de: R$" ,duasXJuros, "com juros de 10%")
    else:
        print("Opção inválida!\n\n")
        func(produto)

func(produto)