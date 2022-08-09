
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


-- MELHORANDO AS FUNÇÕES - FUNÇÕES COM RETORNE DE VALORES:

function pular(intensidade)
  print("Vou pular com a intensidade: " .. intensidade)
end

function calcularFisica(forca)
  return forca * 0.5 + 32.98 / 4
end

function calcularFormulaSecreta(senha)
  return senha + 1
end

pular(calcularFisica(13.5) + calcularFormulaSecreta(987))

-- OUTRA FORMA DE EXECUTAR      

fisica = calcularFisica(1)
secreta = calcularFormulaSecreta(1)

pular(fisica + secreta)
--OU
intensidade = fisica + secreta
pular(intensidade)
