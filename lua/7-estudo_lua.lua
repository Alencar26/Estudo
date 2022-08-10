

-- ESTUDO DE LUA
--
-- Nesse arquivo será abordado o assunto de loop com WHILE
--
condicao = true

while condicao do
 print("Esse é um loop infinito.\n Digite: [*] - continuar [n] - não continuar")
 entrada = io.read()
 if entrada == "n" or entrada == "N" then
   condicao = false
   print("Ufa!!! Você saiu do loop.\n")
 end
end

-- outro exemplo

indice = 1

while indice <= 10 do
  print(indice)
  indice = indice + 1
end

-- outro exemplo

while io.read() == "F" do
 print("ATIRAR") 
end
