


-- ESTUDO DE LUA
--
-- Nessa arquivo será visto condicionais IF, ELSE e THEN 


function condicao(a,b)
  result = a + b
  if result < 10 then
    print("Valor é menor que 10")
  elseif result == 11 then
    print("Valor 11")
  else 
    print("valor maior que 11")
  end 
end

condicao(2,6)

-- operador E (and)


function fizzBuzz(a)

  for i=1,a do
    if i % 3 == 0 and i % 5 == 0 then
      print("FizzBuzz")
    elseif result % 3 ~= 0 and i % 5 == 0 then
      print("Buzz")
    else
      print("Fizz")
    end
  end
end

fizzBuzz(30)

-- operador OR (ou)

function contarPontos(nome)

  if(nome == "Tartaruga" or nome == "planta") then
    return 100
  else
    return 50
  end
end

ponto = contarPontos("Arroz")
print("[Arroz] - Seu ponto foi: " .. ponto)

ponto = contarPontos("planta")
print("[planta] - Seu ponto foi: " .. ponto)

ponto = contarPontos("Tartaruga")
print("[Tartaruga] - Seu ponto foi: " .. ponto)


-- Operador NOT - inverte o valor seguinte
-- exemplo: not true > isso é igual a false
-- not true = false


if not false then
 print("Não é false, logo é verdadeiro.")
end
