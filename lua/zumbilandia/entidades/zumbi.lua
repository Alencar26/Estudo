local inimigo = require("entidades/inimigo")
local player = require("entidades/player")
local zumbi = {}

function zumbi.novo()
  local zumbi = inimigo.novo(50, "zumbis")
  zumbi.come_cerebro = true
  zumbi.explode = false
  return zumbi
end

function zumbi.novo_bomber()
  local zumbi = zumbi.novo()
  zumbi.come_cerebro = false
  zumbi.explode = true
  zumbi.forca = 100
  return zumbi
end

function zumbi.atacar(zumbi, oponente)
  if oponente.vida > 0 then
    print("Zumbi atacou: " .. oponente.nome .. "!")
    player.atacado(oponente, zumbi.forca)

    if oponente.vida <= 0 then
      print(oponente.nome .. " - Morreu!")
    end
  else
    print("Zumbi está devorando o " .. oponente.nome .. " que está morto.")
  end
end

return zumbi
