local player = require("entidades/player")
local pocao = require("entidades/pocao")
local zumbi = require("entidades/zumbi")

-- Criação de Instâncias das Classes Criadas
-- Essas variáveis são objetos!
player1 = player.novo("Paladino")
player2 = player.novo("Arqueiro")

zumbi1 = zumbi.novo()
zumbi2 = zumbi.novo()
zumbi_bomba = zumbi.novo_bomber()


-- Adicionar inventário
player.obter_pocao(player1, pocao.novo())
player.obter_pocao(player2, pocao.novo())

-- testando se os players tem 1 item.
assert(#player1.pocoes == 1)
assert(#player2.pocoes == 1)

-- Zumbis atacam os players
zumbi.atacar(zumbi1, player2)
zumbi.atacar(zumbi_bomba, player1)
zumbi.atacar(zumbi2, player1)
zumbi.atacar(zumbi_bomba, player1)
zumbi.atacar(zumbi2, player1)

-- verificando a força do zumbi bomba
assert(zumbi_bomba.forca == 100)

zumbi.atacar(zumbi2,player2)
player.usar_pocao(player2)
zumbi.atacar(zumbi_bomba, player2)
player.usar_pocao(player2)
