

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

from datetime import date
class Conta:

    taxa = 0.5
    
    #método de classe
    @classmethod
    def retornarCodigo(cls):
        print('Código: 555')
    
    #método estático    
    @staticmethod
    def retornarCodigoBanco():
        return '5598'

    #métoodos de instância
    def __init__(self, agencia, numero, saldo):
        self.agencia = agencia
        self.numero = numero
        self.saldo = saldo - Conta.taxa
        self.__historico = [{
                            'Operacao': '+ Credito',
                            'data': str(date.today()),
                            'saldo': saldo
                            }]
    
    def trasacao(self, saldo):
        operacao = self.__consultaOperacao(saldo)
        
        registro = {
            'operacao': operacao,
            'data': str(date.today()),
            'saldo': saldo
        }
        self.__historico.append(registro)
        
    def __consultaOperacao(self, saldo) -> str:
        if self.__historico[-1].get('saldo') > saldo:
            return '- Debito'
        elif self.__historico[-1].get('saldo') < saldo:
            return '+ Credito'
        else:
            return '= Manteve'
    
    def trasacao(self, saldo):
        operacao = self.__consultaOperacao(saldo)
        registro = {
            'operacao': operacao,
            'data': str(date.today()),
            'saldo': saldo
        }
        self.__historico.append(registro)
        
        
    def depositar(self, valor):
        self.saldo += valor
        self.trasacao(self.saldo)
    
    def sacar(self, valor):
        self.saldo -= valor
        self.trasacao(self.saldo)
        return self.saldo
    
    def trasferir(self, valor, destino):
        self.sacar(valor)
        destino.depositar(valor)
        
    def extrato(self):
        for saldo in self.__historico:
            print(saldo)
    
    @classmethod
    def transferirValores(self, valor, origem, destino):
        origem.sacar(valor)
        destino.depositar(valor)
    
        
    
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

#métodos estático e metodo de classe
Conta.retornarCodigo() #chamando método de classe
print(Conta.retornarCodigoBanco()) #chamando método estático


#--- Transferência
print(15 * '-', 'Transferência', 15 * '-')

conta1 = Conta(123, 1234, 3000)
conta2 = Conta(123, 1235, 5000)

print(f"Saldo conta 1: {conta1.saldo}")
print(f"Saldo conta 2: {conta2.saldo}")

print(5 * '-', 'Operação', 5 * '-')
conta1.trasferir(500, conta2)
print(f"Saldo conta 1: {conta1.saldo}")
print(f"Saldo conta 2: {conta2.saldo}")

print(5 * '-', 'Operação', 5 * '-')
Conta.transferirValores(1000, conta2, conta1)
print(f"Saldo conta 1: {conta1.saldo}")
print(f"Saldo conta 2: {conta2.saldo}")

print(5 * '-', 'Operação', 5 * '-')
conta1.trasferir(234, conta2)
print(f"Saldo conta 1: {conta1.saldo}")
print(f"Saldo conta 2: {conta2.saldo}")

print(5 * '-', 'Operação', 5 * '-')
Conta.transferirValores(450, conta2, conta1)
print(f"Saldo conta 1: {conta1.saldo}")
print(f"Saldo conta 2: {conta2.saldo}")

print(5 * '-', 'Operação', 5 * '-')
Conta.transferirValores(450, conta2, conta1)
print(f"Saldo conta 1: {conta1.saldo}")
print(f"Saldo conta 2: {conta2.saldo}")


print(5 * '-', 'Extrado / Historico', 5 * '-')
print(conta1.extrato())
print(conta2.extrato())
