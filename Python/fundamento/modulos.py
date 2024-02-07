#IMPORTANDO TODO O MÓDULO-----------------------------------------

import random


#IMPORTANDO FUNÇÃO ESPECÍFICA-------------------------------------

from random import choice

#IMPORTANDO FUNÇÕES ESPECÍFICAS----------------------------------

from random import randint, shuffle

#OU

from random import (
    randint,
    shuffle,
    choice
)

#USANDO ALIAS NO IMPORT----------------------------------------

import random as rd

#ou

from random import choice as sortear

lista = [1, 2, 3, 4, 5]
print(sortear(lista))

#ALIAS COM VARIAS FUNÇÕES---------------------------------------

from random import (
    randint as aleatorio,
    shuffle as embaralhar,
    choice as sortear
)

lista = [1, 2, 3, 4, 5]
print(sortear(lista))
print(aleatorio(1, 10))
embaralhar(lista)
print(lista)