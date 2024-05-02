"""
Você é um ladrão profissional que planeja roubar casas ao longo
de uma rua. Cada casa tem uma certa quatia de dinheiro escondida,
a única restrição que impede de roubar cada uma dekas é que as 
casas adjacentes têm siistemas de segurança conectadas e entrará 
em contato automaticamente com a polícia se duas casas adjacentes
forem roubadas na mesma noite.


Dado um array de nums representando a quantidade de dinheiro de cada
casa, retorne a quatidade máxima de dinheiro que você consegue roubar
em uma única noite sem chamar a polícia.
"""

casa_na_rua: list[int] = [2,7,9,3,1]

def roubar(casas: list[int]) -> int:
    acomulado_duas_anterior, acomulado_anterior = 0,0
    for atual in casas:
        max_acomulado = max(atual + acomulado_duas_anterior, acomulado_anterior)
        acomulado_duas_anterior = acomulado_anterior
        acomulado_anterior = max_acomulado
    return acomulado_anterior

print(roubar(casa_na_rua))

"""
Comportamento do algorítimo:

0,0,2
0,2,7
2,7,9
7,11,3
11,11,1
11,12,none
"""

