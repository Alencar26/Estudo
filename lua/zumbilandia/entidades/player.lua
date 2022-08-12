local player = {}

function player.novo(nome)
  return {
    vida = 200,
    pocoes = {},
    nome = nome
  }
end

return player
