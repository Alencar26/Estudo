class Pessoa:
    
    def __init__(self, nome, idade):
        self.nome = nome
        self.idade = idade
        self.nomeclasse = self.__class__.__name__

    def falar(self):
        print(f'{self.nomeclasse} falando...')

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
    
    def dizer_aula(self):
        print(f'{self.nomeclasse} dando aula...')
        
aluno1 = Aluno('João', 20)
aluno1.falar()

professor1 = Professor('Maria', 30)
professor1.falar()

aluno1.estudar()
professor1.dizer_aula()

print(aluno1)