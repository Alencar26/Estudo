--
--
-- ESTUDO LUA
--
-- CRIÇÃO DE PACOTE

local calculadora = {

  somar = function(x, y)
    return x + y
  end
}

function calculadora.multiplicar(x, y)
  return x * y
end


return calculadora
