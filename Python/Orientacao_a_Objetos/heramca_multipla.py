class Funcionario:
    
    def trabalhar(self):
        print("Trabalhando")
        
class FrontEnd(Funcionario):
    def programar(self):
        print("Frontend")

class BackEnd(Funcionario):
    def programar(self):
        print("Backend")
        
class FullStack(FrontEnd, BackEnd):
    def programar(self):
        print("FullStack")
        
dev = FullStack().programar()