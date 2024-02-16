from abc import ABC, abstractmethod

# ABC significa Abstract Base Class
#precisamos herdar de ABC para que a classe seja abstrata
class Pessoa(ABC):
    def __init__(self, nome, idade):
        self.nome = nome
        self.idade = idade
        self.nomeclasse = self.__class__.__name__

    @abstractmethod
    def falar(self):
        pass
                
class Cliente(Pessoa):
    
    def falar(self) -> str:
        print(f'{self.nomeclasse} falando...')
        return f'{self.nomeclasse} falando...'

class Aluno(Pessoa):
    
    def falar(self) -> int:
        print(f'{self.nomeclasse} estudando...')
        return 0


cliente = Cliente('João', 25)
cliente.falar()

aluno = Aluno('Maria', 18)
aluno.falar()