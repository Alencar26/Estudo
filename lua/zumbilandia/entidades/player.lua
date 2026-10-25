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

function player.atacado(player, forca)
  if (player.vida > 0) then
    player.vida = player.vida - forca
  end 
end

function player.usar_pocao(player)
  if #player.pocoes > 0 then
    player.vida = player.vida + player.pocoes[1].vida
    table.remove(player.pocoes, 1)
    print(player.nome .. " usou uma poção e agora tem " .. player.vida .. " de vida!")
  else
    print(player.nome .. " - Não tem poções.")
  end
end

return player
