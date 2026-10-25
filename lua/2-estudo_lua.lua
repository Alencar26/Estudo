--[[
--
-- Estudo lua
--
-- Nessa arquivo será abordado:
--
-- 1.Booleano
-- 2.Igualdade
-- 3.Conversão para texto
--
--]]


estaAtivo = false
naoEstaAtivo = true

print(estaAtivo == false)

-- Não pode concatenar booleano que não esteja convertido para string
print("Valor está inativo? R: " .. tostring(naoEstaAtivo))

print("Está inativo? R: " .. tostring(estaAtivo == false))


