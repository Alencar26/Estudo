"""
USANDO PACOTE RECEM GERADO PARA ENTENDER O FUNCIONAENTO DE UMA 
PACOTE EM PYTHON.
"""

from pacote import (
    principal, 
    secundario
)

print(principal.soma(5,7))
print(secundario.quadratica(10))

#EXEMPLO COM SUB DIRETÓRIO-------------------

from pacote.sub_dir import terciario as mod3

print(mod3.cubica(5))


#CHAMANDO APENAS O MÓDULO DESEJADO----------------

from pacote.principal import soma

print(soma(3, 2))

#--- 

from pacote.sub_dir.terciario import cubica

print(cubica(2))