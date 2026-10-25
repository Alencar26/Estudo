local fantasma = require("fantasma")
local player = require("player").novo()

local fantasmas = {}

for nome = 1, 2 do
  local gasparzinho = fantasma.novo(nome)
  player:adicionarOuvinte("onPoweredUp", gasparzinho, function()
    gasparzinho:fugir()
  end)
  player:adicionarOuvinte("onPoweredDown", gasparzinho, function()
    gasparzinho:atacar()
  end)
  player:adicionarOuvinte("onStarted", gasparzinho, function()
    gasparzinho:atacar()
  end)

  gasparzinho:adicionarOuvinte("onDied", player, function()
    player:removerOuvinte(gasparzinho)
  end)

  table.insert(fantasmas, gasparzinho)
end

player:iniciar()
player:poderoso()
player:fraco()
fantasmas[2]:morrer()
player:poderoso()
player:fraco()
fantasmas[1]:morrer()
player:poderoso()
