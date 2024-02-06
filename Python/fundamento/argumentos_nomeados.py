"""
ARGUMENTOS NOMEADOS NA FUNÇÃO
"""

def print_name(name, last_name, age):
    print(f'Nome: {name}')
    print(f'Sobrenome: {last_name}')
    print(f'Idade: {age}')
    
print_name(name='João', last_name='Silva', age=30)

"""
PARÂMETRO PADRÃO
"""

def print_name(name=None, last_name="Inexistente", age="Inexistente"):
    print(f'Nome: {name}')
    print(f'Sobrenome: {last_name}')
    print(f'Idade: {age}')
    
print_name(last_name='Silva', age=30)
print_name(name='João')
print_name()

#OUTRO EXEMPLO

# No exemplo abaixo temos 2 parâmetros obrigatórios e um parâmetro opcional.
def print_imovel(descricao_imovel, n_quartos, vagasGaragem=None):
    print(f'Descrição do imóvel: {descricao_imovel}')
    print(f'Número de quartos: {n_quartos}')
    if vagasGaragem:
        print(f'Número de vagas na garagem: {vagasGaragem}')
        
print_imovel('Casa com 2 quartos', 2, 3)
print_imovel('Casa com 1 quarto', 1)