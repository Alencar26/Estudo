
--ESTUDO DE LUA
--
-- ARRAY ou TABELA

vidaInimigos = {
  10, 13, 8, 2, 45, 50, 50, 50, 16, 27
}

print(vidaInimigos[3])


-- # mostra quantos elementos há em uma tabela
print("Quantos inimigos nós temos?")
print(#vidaInimigos)


-- Adicionar um elemento na tabela/array em tempo real
qntInimigos = #vidaInimigos
vidaInimigos[qntInimigos + 1] = 33

print(vidaInimigos[11])



-- loop para printar os dados da tabela
for i=1, #vidaInimigos,1 do
  print("Inimigo ".. i .. " tem " .. vidaInimigos[i] .. " de vida.")
end


-- DICIONÁRIO - chave/valor

itens = {
  machado = 100, -- essa forma tbm está correta
  ["espada"] = 450, -- essa forma tbm está correta
  escudo = 650,
  elmo = 150,
  chave = 45
}

print(itens["machado"])
-- ou
print(itens.machado)
print(itens.espada)
print(itens.elmo)

-- adicionando novo valor no dicionario
itens.armadura = 1036
print(itens.armadura)

-- printar todos os valores do dicionario com loop
for chave, valor in pairs(itens) do
   print(chave .. ": " .. valor)
end

-- no dicionário não é possível usar o # para ver o total de itens 
-- no dicionário. Para saber o tota, é preciso somar de forma manual com loop

-- mais exemplos

inimigo = {
  arqueiro = {
    posicoes = {x = 10, y = 20},
    forca = 110,
    nome = "John Willians"
  }, 
  canhao ={120, 40}
}

print("Informação do iniigo: " .. inimigo.arqueiro. posicoes.y)


