

# Paradígma Imperativo - Procedural 

import email
def Registrar(nome, idade, cpf, email):
    paciente = {'nome': nome, 'idade': idade, 'cpf': cpf, 'email': email}
    return paciente

# Paradígma Orientado a Objetos - OO

class Paciente:
    
    def __init__(self, nome, idade, cpf, email):
        self.nome = nome
        self.idade = idade
        self.cpf = cpf
        self.email = email
        
paciente1 = Paciente('João', 20, '123.456.789-10', 'XXXXXXXXXXXXXX')
paciente2 = Paciente('Maria', 30, '123.456.789-11', 'XXXXXXXXXXXXXX')
print(paciente1.nome)
print(paciente2.idade)

print(paciente1.__dict__) #para imprimir minha instância
print(paciente2.__dict__)

#------------------------------------------------------------------------

class Conta:

    taxa = 0.5

    def __init__(self, agencia, numero, saldo):
        self.agencia = agencia
        self.numero = numero
        self.saldo = saldo - Conta.taxa
        
    def depositar(self, valor):
        self.saldo += valor
    
    def sacar(self, valor):
        self.saldo -= valor
        return self.saldo
    
cliente = Conta(123, 1234, 1500)
print(f"Conta aberta: {cliente.saldo}")
cliente.depositar(100)
print(f"Saldo adicionado: {cliente.saldo}")
cliente.sacar(50)
print(f"Saldo sacado: {cliente.saldo}")

print(cliente.__dict__)
cliente.bonus = 350 #adicionando atributo dinâmico (criado em tempo de execução)
print(cliente.__dict__)

del cliente.bonus #deletando atributo dinâmico (deletado em tempo de execução)
print(cliente.__dict__)