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

"""
ARGUMENTO ARBITRÁRIO - ARGS
Quando não sabemos o número de argumentos que serão passados para a função
"""

def print_names(*args):
    print(args)
    
print_names('João', 'Maria', 'Pedro') # imprime um tupla com os argumentos

lista_nomes = ['João', 'Maria', 'Pedro']
print_names(lista_nomes) # imprime uma lista com os argumentos
print_names(*lista_nomes) # o * é usado para desempacotar a lista


"""
ARGUMENTO ARBITRÁRIO - KWARGS
Quando não sabemos o nome dos argumentos que serão passados para a função,
a principal diferença entre os dois exemplos acima é que o primeiro utiliza 
um dicionário para receber os argumentos, e o segundo utiliza um *args para receber
os argumentos.
"""

def print_names(**kwargs):
    print(kwargs)

print_names(name='João', last_name='Silva', age=30) # imprime um dicionário com os argumentos

def print_sala_de_aula(professora,**alunos):
    print(f'Professora: {professora}')
    for aluno, nota in alunos.items():
        print(f'Aluno: {aluno}, nota: {nota}')

print_sala_de_aula('Maria', João=10, Pedro=9, Ana=8)