class Aluno:
    def __init__(self, name, age, sex):
        self.__name = name
        self.__age = age
        self.__sex = sex
        self.__codigo_aluno = 32562

    def __str__(self):
        return f"Name: {self.name}, Age: {self.age}, Sex: {self.sex}"
    
    def get_name(self):
        return self.__name
    
    def set_name(self, name):
        self.__name = name
        
    def get_age(self):
        return self.__age
    
    def set_age(self, age):
        self.__age = age
        
    def get_sex(self):
        return self.__sex
    
    def set_sex(self, sex):
        self.__sex = sex
    
    def delete_sex(self, sex):
        del self.__sex
    
    @property   
    def codigo(self):
        return self.__codigo_aluno
    
    @codigo.setter
    def codigo(self, cod_aluno):
        self.__codigo_aluno = cod_aluno
        return self.__codigo_aluno
        
    name = property(get_name, set_name)
    age = property(get_age, set_age)
    sex = property(get_sex, set_sex, delete_sex, doc =  'sexo do aluno, getters e settes (delete tmb)')
    


aluno = Aluno('Joao', 20, 'M')
aluno.name = 'Maria'
print(aluno.name)
print(aluno)

print(aluno.codigo)
aluno.codigo = 00000
print(aluno.codigo)

print(Aluno.sex.__doc__)