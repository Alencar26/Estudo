local player = {}

function player.novo(nome)
  return {
    vida = 200,
    pocoes = {},
    nome = nome
  }
end

function player.obter_pocao(player, pocao)
  table.insert(player.pocoes, pocao)
end

return player
