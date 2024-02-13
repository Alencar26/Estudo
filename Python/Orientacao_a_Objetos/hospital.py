from datetime import date

class Paciente:
    
    def __init__(self, nome, idade, cpf, email) -> None:
        self.nome = nome
        self.idade = idade
        self.cpf = cpf
        self.email = email
        
    @classmethod
    def criar_paciente(cls, nome, ano_nascimento, cpf, email):
        idade = date.today().year - ano_nascimento
        return cls(nome, idade, cpf, email)
    
class Medico:
    pass

paciente_01 = Paciente.criar_paciente('João', 1990, '123.456.789-10', 'XXXXXXXXXXXXXX')
print(paciente_01.__dict__)
print(paciente_01.idade)