
-- ESTUDO DE LUA
--
-- Nesse arquivo será visto a implementação de funções

-- função sem parâmetro
function somarUm()
  print(1 + 1)
end

somarUm()

-- função com parâmetro
function somar(a,b)
  print(a + b)
end

somar(5,16)

-- com string
function pegarNome(nome)
  print("Seu nome é: " .. nome)
end

print("Informe seu nome: ")
pegarNome(io.read())

