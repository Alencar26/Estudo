


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
