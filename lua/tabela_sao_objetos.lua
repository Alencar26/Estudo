-- ==============================================
-- Razões pelas quais tabelas Lua são Objetos
-- ==============================================

-- Tem estado
zumbi = {vida = 10, x = 20, y = 30}

-- Tem identidade, independentemente dos valores
print(zumbi)
zumbi.vida = 50
print(zumbi)

-- 2 tabelas com os mesmos valores são objetos diferente
novoZumbi = {vida = 50, x = 20, y = 30}
assert(zumbi ~= novoZumbi)
print(novoZumbi)

-- uma tabela pode ter valores diferente em momentos diferentes-- mas é sempre o mesmo objeto
zumbi.vida = 30
print(zumbi)

-- Tabelas tem cilco de vida independenre de lugar e de quem as criou

function novoZumbi(vida)
  return {vida = vida}
end

zumbiX = novoZumbi(100)
print(zumbiX)

function alterar_zumbi(zumbi)
  zumbi.vida = zumbi.vida + 1
  print("Dentro: ")
  print(zumbi)
end

alterar_zumbi(zumbiX)
print(zumbiX.vida)
print(zumbiX)
