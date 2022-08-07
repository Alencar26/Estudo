
-- Estudo lua
--
-- Nesse arquivo srá visto:
--
-- 1. Números negativos
-- 2. Decimais
-- 3. Valores Nuloos 
-- 4. Escape de texto

resultado = 4 - 10
print("Numero negativo: " .. resultado)

resultado = - 34
print(resultado)

numeroReal = 45.5
print("decimal: " .. numeroReal)

pi = 3.14588927289209
print("PI: " .. pi)


-- Null em lua é nil
umaVariavelNula = nil
print(umaVariavelNula)

-- ESCAPE DE STRING - usa \ para informar onde a string vai escapar
nome = "caixa d'agua"
print(nome)

nome = "Caixa D\'água"
print(nome)

-- MELHOR EXEMPLO DE ESCAPE
texto = "Você é \"rico\""
print(texto)

-- sequisar colocar a \ no texto, precisar fazer o escape da barra;
print("texto com \\")
-- quebra de linha - \n
variasLinhas = " linha1 \n linha2 \n linha3"
print(variasLinhas)
-- tabulacao - \t
tabulacao = "\tTabulacao antes de mim."
print(tabulacao)

-- retornar na posição inicial - \r
retornei = "\t\rRetornei"
print(retornei)
