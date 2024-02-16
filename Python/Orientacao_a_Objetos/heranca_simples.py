class Pessoa:
    
    def __init__(self, nome, idade):
        self.nome = nome
        self.idade = idade
        self.nomeclasse = self.__class__.__name__

    def falar(self):
        print(f'{self.nomeclasse} falando...')
        
    def trabalha(self):
        print(f'{self.nomeclasse} trabalhando...')

class Aluno(Pessoa):
    
    def __init__(self, nome, idade):
        super().__init__(nome, idade)
        self.matricula = 123456
        self.curso = 'Informática'
        self.turma = '1ºA'
        self.ano = 2021
        self.semestre = 1
        self.periodo = 'Manhã'
        self.turno = 'Matutino'
        self.horario = '08:00'
        self.sala = 'A1'
        self.disciplina = 'Python'
        self.professor = 'Maria'
        self.aluno = 'João'
        self.nota = 7.5
        self.faltas = 0
        self.presenca = 100
        self.status = 'Aprovado'
        self.observacoes = 'Sem observações'
        self.data = '01/01/2021'
        self.hora = '08:00'
    
    def __str__(self) -> str:
        return f"Nome: {self.nome} - Matricula {self.matricula} - Curso {self.curso}"
    
    def estudar(self):
        print(f'{self.nomeclasse} estudando...')
        
class Professor(Pessoa):
    
    def trabalha(self):
        print(f'{self.nomeclasse} Ensinando...')

class Administrativo(Pessoa):
    
    def trabalha(self):
        print(f'{self.nomeclasse} Organizando planilhas...')

class Zelado(Pessoa):

    def varrer(self):
        print(f'{self.nomeclasse} Varrendo...')
        
        
aluno1 = Aluno('João', 20)
aluno1.falar()
aluno1.estudar()
print(aluno1)


professor1 = Professor('Maria', 30)
professor1.falar()

admin = Administrativo('Pedro', 40)
zelador = Zelado('Joana', 50)

professor1.trabalha()
admin.trabalha()
zelador.trabalha()