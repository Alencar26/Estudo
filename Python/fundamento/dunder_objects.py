"""
ENTENDENDO OBJETOS DUNDER.
UTILIZADOS PRINCIPALMENTE PARA USO INTERNO DO PYTHON.

SÃO FÁCEIS DE RECONHECER, UTILIZAM O PADRÃO __NOME__.

#PRINCIPAIS DE USO MAIS COMUM.

__init__ ---> utilizado no python 2.x identificar que era um pacote python.
         --> cria-se um arquivo __init__.py no diretório do pacote.
         ---> para que seja compatível com outras versões do python.
         ---> arquivo fica em branco.
__name__ ---> informar o nome do arquivo atual.
__file__ ---> informar o path do arquivo atual.
__main__ ---> informar se o arquivo é o principal.
__doc__  ---> informar a documentação do arquivo.

"""

print("My python file is path... --->", __file__)

"""
SUPONDO QUE ESSE ARQUIVO É UM MÓDUTO E TEM ALGUM CÓDIGO NELE QUE EU NÃO
QUEO QUE EU QUERO QUE SEJA EXECUTADO QUANDO O MÓDULO FOR IMPORTADO.

QUE EU SÓ QUERO EXECUTAR QUANDO ESTIVER RODANDO O MÓDULO DIRETAMENTE COM MAIN.

ENTÃO USA-SE:
"""

#esse código só roda quando o arquivo é executado diretamente.
if __name__ == "__main__":
    print("My module is running...")
    

#EXEMPLO DO DOCSTRIG COM MODULO RANDOM
import random
print(random.__doc__) #documentação do módulo random.

__doc__ = "My documentation is done. Show me!"
print(__doc__)