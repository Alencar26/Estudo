

-- ESTUDO DE LUA
--
-- Nessa arquivo será abordado alguns exemplos da biblioteca de matemática.


-- arredondar um número
--
-- Arredonda para baixo
print(math.floor(0.5))
print(math.floor(1.6))
print("")
-- Arredonda para cima
print(math.ceil(0.5))
print(math.ceil(1.6))

-- retorna o maior valor
print(math.max(3, 6, 2, 7))


--substituir string por outra
print(string.gsub("olá mundo", "mundo", "game"))

--substituir string e n ocorrencias
-- código abaixo só vai realizar a troca nas duas primeiras ocorrencias
print(string.gsub("Olá Mundo, Olá Mundo, Olá Mundo, Olá Mundo", "Mundo", "Cosmos", 2))


-- manipulação de arquivos
-- exemplo mapa de um jogo
--      X = enimigo
--      # = arbusto
--      $ = cofre 
--      = = caminho

mapa = io.lines("mapa_exemplo.txt")
for linha in mapa do
   print(linha)
end
-- ou direto no for
for linha in io.lines("mapa_exemplo.txt") do
  print(linha)
end




----- USANDO MINHA PRÓPRIA BIBLIOTECA

local calculadora = require("calculadora")

print(calculadora.somar(2,2))
print(calculadora.multiplicar(3,3))
