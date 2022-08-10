


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
