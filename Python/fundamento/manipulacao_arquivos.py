import os

print("Sistema:",os.name)
print("Núcleos de CPU:",os.cpu_count())

#validar se um arquivo ou pasta existe
print(os.path.exists("fundamento/pacote"))
print(os.path.exists("teste.txt"))

#criar um arquivo
os.mknod("fundamento/OlaMundo.py")
os.mkdir('fundamento/novaPasta')

#diretório atual do arquivo
print(os.path.dirname(__file__))
print(os.getcwd())

__path__ = os.getcwd()
print(os.listdir(__path__))

#remover arquivos e diretórios
os.remove("fundamento/OlaMundo.py")
os.rmdir('fundamento/novaPasta')

#renomear arquivo e diretórios
os.mknod("teste.txt")
os.rename("teste.txt", "testeRenomeado.txt")
os.rename("fundamento/pacote", "fundamento/Pacote")

#MOVER DIRETÓRIO
os.rename("testeRenomeado.txt", "fundamento/Pacote/testeMovido.txt")