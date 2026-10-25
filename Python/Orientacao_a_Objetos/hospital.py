from datetime import date

class Paciente:
    
    def __init__(self, nome, idade, cpf, email) -> None:
        self.__nome = nome  #private
        self.idade = idade  #public
        self.cpf = cpf      #public
        self._email = email #protected
        
    @classmethod
    def criar_paciente(cls, nome, ano_nascimento, cpf, email):
        idade = date.today().year - ano_nascimento
        return cls(nome, idade, cpf, email)
    
    def retornaNome(self):
        self.__metodo_privado() #chamando o método privado dentro do método publico
        return self.__nome
    
    def __metodo_privado(self):
        print('Este método é privado')
    
class Medico:
    pass

paciente_01 = Paciente.criar_paciente('João', 1990, '123.456.789-10', 'XXXXXXXXXXXXXX')
print(paciente_01.__dict__)
print(paciente_01.idade)

print(paciente_01.retornaNome())