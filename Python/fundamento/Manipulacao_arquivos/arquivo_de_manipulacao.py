import os

__path__ = "./fundamento/Manipulacao_arquivos"

if not os.path.exists(__path__):
    os.mkdir(__path__)
else:
    if not os.path.exists(f"{__path__}/receita.txt"):
        os.mknod(f"{__path__}/receita.txt")
    else:
        pass
        
arquivo = open(f"{__path__}/receita.txt")
print(f"{30*'-'} - Aberto\n",arquivo.read())
arquivo.close()
print(f"{30*'-'} - fechado\n") if arquivo.closed else print(f"{30*'-'} - aberto\n")

brigadeiro = open(f"{__path__}/receitas/brigadeiro.txt", "r")
print(f"{30*'-'} - Aberto\n",brigadeiro.read())
brigadeiro.close()
print(f"{30*'-'} - fechado\n") if brigadeiro.closed else print(f"{30*'-'} - aberto\n")

#--- Ler só primeira linha
arquivo = open(f"{__path__}/receita.txt", "r")
brigadeiro = open(f"{__path__}/receitas/brigadeiro.txt", "r")

print(arquivo.readline())
print(brigadeiro.readline())

arquivo.close()
brigadeiro.close()

#--- Forma mais prática para abrir e fechar aquivo

with open(f"{__path__}/receita.txt", "r") as arquivo_nova_forma:
    print(f"{30*'-'} - Aberto\n",arquivo_nova_forma.read())
print(f"{30*'-'} - Fechado\n") if arquivo_nova_forma.closed else print(f"{30*'-'} - aberto\n")
  
#--- Escrevendo no arquivo
texto = """
Passando uma variável que recebeu aspas triplas é possível 
haver quebra de linha no texto.

Sem se preaocupar os \\n para que o texto fiue mais orgânico.
Massa né?! 
"""
with open(f"{__path__}/receita.txt", "a") as arquivo_nova_forma:
    arquivo_nova_forma.write("\n15min ao forno em 220 ºc - pré aquecido")
    arquivo_nova_forma.write(texto)